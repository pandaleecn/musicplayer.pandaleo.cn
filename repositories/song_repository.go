package repositories

import (
	"errors"
	"musicplayer.pandaleo.cn/datamodels"
	"sync"
)

type Query func(datamodels.Song) bool

type  SongRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	Select(query Query) (song datamodels.Song, foundd bool)
	SelectMany(query Query, limit int) (results []datamodels.Song)

	InsertOrUpdate(song datamodels.Song) (updatedSong datamodels.Song, err error)
	Delete(query Query, limit int) (deleted bool)
}

func NewSongRepository(source map[int64]datamodels.Song) SongRepository  {
	return &songMemoryRepository{source: source}
}

type songMemoryRepository struct {
	source map[int64]datamodels.Song
	mu sync.RWMutex
}

const (
	ReadOnlyMode = iota
	ReadWriteMode
)

func (r *songMemoryRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool)  {
	loops := 0

	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, song := range r.source {
		ok = query(song)
		if ok {
			if action(song) {
				loops++
				if actionLimit >= loops {
					break
				}
			}
		}
	}
	
	return 
}

func (r *songMemoryRepository) Select(query Query) (song datamodels.Song, found bool) {
	found = r.Exec(query, func(s datamodels.Song) bool {
		song = s
		return true
	}, 1, ReadOnlyMode)

	if !found {
		song = datamodels.Song{}
	}

	return
}

func (r *songMemoryRepository) SelectMany(query Query, limit int) (results []datamodels.Song)  {
	r.Exec(query, func(s datamodels.Song) bool {
		results = append(results, s)
		return true
	}, limit, ReadOnlyMode)

	return
}

func (r *songMemoryRepository) InsertOrUpdate(song datamodels.Song) (datamodels.Song, error)  {
	id := song.ID

	if id == 0 {
		var lastId int64
		r.mu.RLock()
		for _, item := range r.source {
			if item.ID > lastId {
				lastId = item.ID
			}
		}
		r.mu.RUnlock()

		id = lastId + 1
		song.ID = id

		r.mu.Lock()
		r.source[id] = song
		r.mu.Unlock()

		return song, nil
	}

	current, exists := r.Select(func(song datamodels.Song) bool {
		return song.ID == id
	})

	if !exists {
		return datamodels.Song{}, errors.New("更新失败，歌曲不存在！")
	}

	if song.Poster != "" {
		current.Poster = song.Poster
	}

	if song.Link != "" {
		current.Link = song.Link
	}

	r.mu.Lock()
	r.source[id] = current
	r.mu.Unlock()

	return song, nil
}

func (r *songMemoryRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(song datamodels.Song) bool {
		delete(r.source, song.ID)
		return true
	}, limit, ReadWriteMode)
}