package model

type MeetUp struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      string `json:"userId"`
}

type NewMeetUp struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
