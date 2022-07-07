CREATE DATABASE IF NOT EXISTS users;
use users;

-- CREATE TABLE users (
--     user_id INT(8) NOT NULL,
--     password VARCHAR(270) NOT  NULL,
--     email varchar(30) unique,
--     youtube_api_key VARCHAR(270) NOT NULL,
--     toeic_score INT(3) NULL,
--     last_watched_video_id VARCHAR(30) NULL
--     ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ALTER TABLE users
--   ADD PRIMARY KEY (user_id);

-- ALTER TABLE users
--   MODIFY user_id int(8) AUTO_INCREMENT,AUTO_INCREMENT=1;

-- INSERT INTO users (password, email, youtube_api_key, toeic_score, last_watched_video_id)
--         VALUES ("5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8",
--                 "example@gmail.com",
--                 "AIzaSyDKY9T6Z7XYPRlCx6fLdfeujLfmPnDkldk",
--                 0,
--                 "BN9yqF6Um98");