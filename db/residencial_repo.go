package db

import (
	"gorm.io/gorm"
)

type PeopleRepo struct {
	db *gorm.DB
}

func NewPeopleRepo(db *gorm.DB) *PeopleRepo {
	db.AutoMigrate(&People{})
	return &PeopleRepo{
		db: db,
	}
}

func (p PeopleRepo) CreatePeople(model People) error {
	err := p.db.Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func (p PeopleRepo) DeletePeople(identity int64) error {
	err := p.db.Where("IdentityNumber=?", identity).Delete(&People{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (p PeopleRepo) UpdatePeople(model People) error {
	err := p.db.Model(&model).Omit("IdentityNumber").Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (p PeopleRepo) GetListPeople() ([]People, error) {
	var models []People
	err := p.db.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (p PeopleRepo) GetPeople(identity int64) (*People, error) {
	var model People
	err := p.db.Where("IdentityNumber = ?", identity).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}
