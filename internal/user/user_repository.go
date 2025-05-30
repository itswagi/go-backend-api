package user

type UserRepository interface {
	GetAll() []User
	Create(User) error
}

type InMemoryUserRepo struct {
	users []User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{}
}

func (r *InMemoryUserRepo) GetAll() []User {
	return r.users
}

func (r *InMemoryUserRepo) Create(u User) error {
	u.ID = len(r.users) + 1
	r.users = append(r.users, u)
	return nil
}
