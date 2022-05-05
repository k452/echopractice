package usecase

import (
	"sake_io_api/domain/model"
	"sake_io_api/domain/repository"
)

type ManagementUsecase interface {
	GetManagementAll() []model.Management
	GetManagement(user_id int) []model.Management
	RegisterManagement(u *model.RegisterManagementType)
	UpdateManagement(u *model.UpdateManagementType)
	DeleteManagement(management_id int)
}

//今はリポジトリのラッパのようになってるけど、よりビジネスロジックが複雑になった時に他の処理を組み合わせたりできるようにしていると解釈
type managementUsecase struct {
	managementRepository repository.ManagementRepository
}

func NewManagementUsecase(mr repository.ManagementRepository) ManagementUsecase {
	return &managementUsecase{
		managementRepository: mr,
	}
}

func (mu managementUsecase) GetManagementAll() []model.Management {
	return mu.managementRepository.GetManagementAll()
}

func (mu managementUsecase) GetManagement(user_id int) []model.Management {
	return mu.managementRepository.GetManagement(user_id)
}

func (mu managementUsecase) RegisterManagement(u *model.RegisterManagementType) {
	mu.managementRepository.RegisterManagement(u)
}

func (mu managementUsecase) UpdateManagement(u *model.UpdateManagementType) {
	mu.managementRepository.UpdateManagement(u)
}

func (mu managementUsecase) DeleteManagement(management_id int) {
	mu.managementRepository.DeleteManagement(management_id)
}
