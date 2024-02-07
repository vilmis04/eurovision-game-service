package group

type Group struct {
	Name        string   `json:"name"`
	Members     []string `json:"members"`
	Owner       string   `json:"owner"`
	DateCreated string   `json:"dateCreated"`
}
