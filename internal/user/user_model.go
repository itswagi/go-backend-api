package user

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateUserDTO struct {
	Name string `json:"name"`
}
