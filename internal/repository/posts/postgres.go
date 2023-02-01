package posts

import (
	"advocate-back/pkg/db"
	"log"
)

var pgSchema = ` CREATE TABLE IF NOT EXISTS posts (
    postId uuid  NOT NULL primary key,
    createdAt timestamp  NOT NULL default NOW(),
    postData jsonb  NOT NULL
)`

type postsRepository struct {
	db db.DB
}

func NewPostsRepository(db db.DB) (*postsRepository, error) {
	_, err := db.DB().Exec(pgSchema)
	if err != nil {
		log.Println("Schema didn't created")
		return nil, err
	}
	return &postsRepository{db: db}, nil
}
