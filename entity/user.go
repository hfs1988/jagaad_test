package entity

type User struct {
	ID       string       `json:"_id"`
	Index    int          `json:"index"`
	GUID     string       `json:"guid"`
	IsActive bool         `json:"isActive"`
	Balance  string       `json:"balance"`
	Tags     []string     `json:"tags"`
	Friends  []UserFriend `json:"friends"`
}

type UserFriend struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
