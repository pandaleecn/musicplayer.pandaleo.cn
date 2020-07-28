package service

import (
	"musicplayer.pandaleo.cn/datamodels"
	"musicplayer.pandaleo.cn/repositories"
)

type SongService interface {
	GetAll() []datamodels.Song
	GetByID(id int64) (datamodels.Song, bool)
	DeleteByID(id int64) bool
	UpdatePosterAndLinkByID(id int64, poster string, link string) (datamodels.Song, error)
}

func NewSongService(repo repositories.SongRepository) SongService {
	return &songService{
		repo: repo,
	}
}

type songService struct {
	repo repositories.SongRepository
}

func (s *songService) GetAll() []datamodels.Song  {
	return s.repo.SelectMany(func(_ datamodels.Song) bool {
		return true
	}, -1)
}

func (s *songService) GetByID(id int64) (datamodels.Song, bool) {
	return s.repo.Select(func(song datamodels.Song) bool {
		return song.ID == id
	})
}

func (s *songService) UpdatePosterAndLinkByID(id int64, poster string, link string) (datamodels.Song, error) {
	return s.repo.InsertOrUpdate(datamodels.Song{
		ID: id,
		Poster: poster,
		Link: link,
	})
}

func (s *songService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(song datamodels.Song) bool {
		return song.ID == id
	}, 1)
}