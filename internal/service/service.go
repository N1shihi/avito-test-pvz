package service

type Usersrv struct {
	userRepo UserRepository
}

func New(userRepo UserRepository) *Usersrv {
	return &Usersrv{userRepo: userRepo}
}
