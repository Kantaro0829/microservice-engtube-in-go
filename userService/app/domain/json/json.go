package json

type CreateUserRequest struct {
	Name               string `json:"name"`
	Password           string `json:"password"`
	Email              string `json:"email"`
	YoutubeApiKey      string `json:"apikey"`
	LastWatchedVideoId string `json:"last_video"`
}
