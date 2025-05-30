package user

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindAll() []User {
	return s.repo.GetAll()
}

func (s *UserService) Create(name string) User {
	u := User{Name: name}
	s.repo.Create(u)
	return u
}