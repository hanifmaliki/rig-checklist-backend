package repository

type Repository interface {
	UserRepository
	RigRepository
	LocationRepository
	WellRepository
	FieldRepository
	ActivityRepository
	FormRepository
	QuestionRepository
	AnswerRepository
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}
