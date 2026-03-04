package rooms

type service struct {
	repository *repository
}

func newService(r *repository) *service {
	return &service{
		repository: r,
	}
}
