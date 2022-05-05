package usecase

import (
	"sake_io_api/domain/model"
	"sake_io_api/domain/repository"
)

type UserUsecase interface {
	GetUserAll() []model.User
	GetUser(user_id int) []model.User
	RegisterUser(u *model.RegisterUserType)
	UpdateUser(u *model.UpdateUserType)
	DeleteUser(user_id int)
}

//今はリポジトリのラッパのようになってるけど、よりビジネスロジックが複雑になった時に他の処理を組み合わせたりできるようにしていると解釈
type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu userUsecase) GetUserAll() []model.User {
	return uu.userRepository.GetUserAll()
}

func (uu userUsecase) GetUser(user_id int) []model.User {
	return uu.userRepository.GetUser(user_id)
}

func (uu userUsecase) RegisterUser(u *model.RegisterUserType) {
	uu.userRepository.RegisterUser(u)
}

func (uu userUsecase) UpdateUser(u *model.UpdateUserType) {
	uu.userRepository.UpdateUser(u)
}

func (uu userUsecase) DeleteUser(user_id int) {
	uu.userRepository.DeleteUser(user_id)
}
