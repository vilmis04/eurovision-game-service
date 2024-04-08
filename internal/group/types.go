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
