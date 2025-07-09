package entity

type UserPost struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TestMap struct {
	Maps []UserPost `json:"maps"`
}
