package repository

import (
	"github.com/eulbyvan/go-enigma-laundry/model"
	"github.com/eulbyvan/go-enigma-laundry/utils"
	"github.com/jmoiron/sqlx"
)

type UomRepository interface {
	GetAll() ([]model.Uom, error)
	GetUomById(id string) ([]model.Uom, error)
	InsertUom(newUom *model.Uom) error
}

type uomRepository struct {
	db *sqlx.DB
}

func (p *uomRepository) GetAll() ([]model.Uom, error) {
	var uoms []model.Uom
	err := p.db.Select(&uoms, utils.SELECT_UOM_LIST)
	if err != nil {
		return nil, err
	}

	return uoms, nil
}

func (p *uomRepository) GetUomById(id string) ([]model.Uom, error) {
	var uoms []model.Uom
	err := p.db.Select(&uoms, utils.SELECT_UOM_ID)
	if err != nil {
		return nil, err
	}

	return uoms, nil
}

func (p *uomRepository) InsertUom(newUom *model.Uom) error {
	_, err := p.db.NamedExec(utils.INSERT_UOM, newUom)
	if err != nil {
		return err
	}
	return nil
}

func NewUomRepository(db *sqlx.DB) UomRepository {
	return &uomRepository{
		db: db,
	}
}
