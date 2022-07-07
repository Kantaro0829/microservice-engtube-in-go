package model

// CREATE TABLE users (
//     user_id INT(8) NOT NULL,
//     password VARCHAR(270) NOT  NULL,
//     email varchar(30) unique,
//     youtube_api_key VARCHAR(270) NOT NULL,
//     toeic_score INT(3) NULL,
//     last_watched_video_id VARCHAR(30) NULL
//     ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

type User struct {
	ID                 int `gorm:"primary_key AUTO_INCREMENT"`
	Name               string
	Password           string
	Email              string `gorm:"unique"`
	YoutubeApiKey      string
	LastWatchedVideoId string
}
