package repo

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"

	repo_entity "gsam/domain/entity/repo"
)

type VideoRepo interface {
	SetupDatabase()
	GetVideo() ([]repo_entity.Video, error)
}

type VideoRepoImpl struct {
	db *sqlx.DB
}

func NewVideoRepo(db *sqlx.DB) VideoRepo {
	return &VideoRepoImpl{
		db: db,
	}
}

func (r *VideoRepoImpl) SetupDatabase() {
	schema := `
    CREATE TABLE IF NOT EXISTS videos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        description TEXT,
        url TEXT
    );`
	r.db.MustExec(schema)

	var count int
	_ = r.db.Get(&count, "SELECT COUNT(*) FROM videos")
	if count == 0 {
		r.db.MustExec(`
        INSERT INTO videos (title, description, url) VALUES
        ('Go Tutorial', 'Learn Go in 10 minutes', 'https://example.com/go'),
        ('Gin Framework', 'REST API with Gin', 'https://example.com/gin'),
        ('Golang Tips', 'Advanced tips and tricks', 'https://example.com/tips');`)
	}

}

func (r *VideoRepoImpl) GetVideo() ([]repo_entity.Video, error) {
	var videos []repo_entity.Video
	err := r.db.Select(&videos, "SELECT id, title, description, url FROM videos")
	if err != nil {
		return nil, err
	}

	return videos, nil
}
