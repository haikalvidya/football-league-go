package payload

type LeagueInfo struct {
	ClubName string `json:"clubname"`
	Points   int    `json:"points"`
	Played   int    `json:"played"`
}

type RecordGameRequest struct {
	ClubHomeName string `json:"clubhomename" validate:"required"`
	ClubAwayName string `json:"clubawayname" validate:"required"`
	Score        string `json:"score" validate:"required"`
}

type LeagueStandingResponse struct {
	ClubName string `json:"clubname"`
	Points   int    `json:"points"`
}

type ClubStandingResponse struct {
	ClubName string `json:"clubname"`
	Standing int    `json:"standing"`
}
