package repository

type Repository interface {
	UserPetrosRepository
	HomeContentRepository
	ProductRepository
	ProductMenuRepository
	CaseStudyRepository
	CaseStudyMenuRepository
	FooterRepository
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}
