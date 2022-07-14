package model

type MeetUp struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        *User  `json:"user"`
}

type NewMeetUp struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
