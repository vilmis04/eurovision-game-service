package group

import "time"

type Group struct {
	Name        string    `json:"name"`
	Members     []string  `json:"members"`
	Owner       string    `json:"owner"`
	DateCreated time.Time `json:"dateCreated"`
}

type CreateGroupRequestBody struct {
	Name string `json:"name"`
}

type UpdateGroupRequestBody struct {
	Name    *string   `json:"name"`
	Members *[]string `json:"members"`
}
