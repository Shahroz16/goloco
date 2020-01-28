package goloco_repo

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	CreateLocation(location LocationModel) (LocationModel, error)
	GetLocation(id int32) (LocationModel, error)
	UpdateLocation(model LocationModel) (LocationModel, error)
	DeleteLocation(id int32) (int32, error)
	GetLocations() ([]LocationModel, error)
}

type repo struct {
	db     *gorm.DB
	logger logrus.Logger
}

func (repo repo) GetLocations() ([]LocationModel, error) {
	var locations []LocationModel
	err := repo.db.Find(&locations).Error
	return locations, err
}

func (repo repo) DeleteLocation(id int32) (int32, error) {
	locationModel := LocationModel{Id: uint(id)}
	err := repo.db.Delete(locationModel).Error
	return int32(locationModel.Id), err
}

func (repo repo) UpdateLocation(location LocationModel) (LocationModel, error) {
	err := repo.db.Save(&location).Error
	return location, err
}

func (repo repo) CreateLocation(location LocationModel) (LocationModel, error) {
	err := repo.db.Create(&location).Error
	return location, err
}

func (repo repo) GetLocation(id int32) (LocationModel, error) {
	locationModel := LocationModel{Id: uint(id)}
	err := repo.db.Where(&locationModel).First(&locationModel).Error
	return locationModel, err
}

func NewRepo(db *gorm.DB, logger logrus.Logger) Repository {
	return repo{
		db:     db,
		logger: logger,
	}
}
