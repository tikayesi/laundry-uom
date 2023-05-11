package usecase

import (
	"github.com/eulbyvan/go-enigma-laundry/model"
	"github.com/eulbyvan/go-enigma-laundry/repository"
)

type UomUseCase interface {
	GetAllUom() ([]model.Uom, error)
	GetUomById(storeId string) ([]model.Uom, error)
	RegisterUom(newUom *model.Uom) error
}

type uomUseCase struct {
	uomRepo repository.UomRepository
}

func (p *uomUseCase) GetAllUom() ([]model.Uom, error) {
	return p.uomRepo.GetAll()
}

func (p *uomUseCase) GetUomById(id string) ([]model.Uom, error) {
	return p.uomRepo.GetUomById(id)
}

func (p *uomUseCase) RegisterUom(newUom *model.Uom) error {
	return p.uomRepo.InsertUom(newUom)
}

func NewUomUseCase(uomRepository repository.UomRepository) UomUseCase {
	return &uomUseCase{
		uomRepo: uomRepository,
	}
}
