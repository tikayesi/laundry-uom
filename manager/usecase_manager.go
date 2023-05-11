package manager

import "github.com/eulbyvan/go-enigma-laundry/usecase"

type UseCaseManager interface {
	UomUseCase() usecase.UomUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) UomUseCase() usecase.UomUseCase {
	return usecase.NewUomUseCase(u.repoManager.UomRepository())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
