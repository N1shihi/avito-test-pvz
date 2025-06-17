package handler

type Handler struct {
	userSrv UserSrv
}

func New(userSrv UserSrv) *Handler {
	return &Handler{
		userSrv: userSrv,
	}
}
