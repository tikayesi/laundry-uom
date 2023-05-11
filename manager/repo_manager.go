package manager

import "github.com/eulbyvan/go-enigma-laundry/repository"

type RepositoryManager interface {
	UomRepository() repository.UomRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (r *repositoryManager) UomRepository() repository.UomRepository {
	return repository.NewUomRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infraManager,
	}
}
