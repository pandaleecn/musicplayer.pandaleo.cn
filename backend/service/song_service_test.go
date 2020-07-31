package service

import (
	"context"
	"musicplayer/entity"
	"musicplayer/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)


func TestSongService_Insert(t *testing.T) {
	conn, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	db := &sql.MySQL{Conn: conn}
	service := NewSongService(db)
	newSong := entity.Song{
		SheetID: 1,
		Name:    "七里香",
		Singer: "周杰伦",
		ImageURL: "https://animage",
		Description: "111",
	}
	mock.ExpectExec("INSERT INTO songs (, position, image_url) VALUES (?,?,?);").
		WithArgs(newSong.SheetID, newSong.Name, newSong.Singer, newSong.ImageURL, newSong.Description).WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := service.Insert(context.TODO(), newSong)
	if err != nil {
		t.Fatal(err)
	}

	if id != 1 {
		t.Fatalf("expected ID to be 1 as this is the first entry")
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}