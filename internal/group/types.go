package group

import "time"

type Group struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Members     []string  `json:"members"`
	Owner       string    `json:"owner"`
	DateCreated time.Time `json:"dateCreated"`
}

type CreateGroupRequestBody struct {
	Name string `json:"name"`
}

type UpdateGroupRequestBody struct {
	Members []string `json:"members"`
}

type JoinGroupRequestBody struct {
	InviteCode string `json:"inviteCode"`
}

type Member struct {
	Name     string `json:"name"`
	Score    uint16 `json:"score"`
	Position int    `json:"position"`
}

type Leaderboard struct {
	Groups     map[int64]string `json:"groups"`
	PlayerList []Member         `json:"playerList"`
}
