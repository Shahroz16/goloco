package goloco_repo

import (
	golocopb "github.com/goloco/src/locationservice/genproto"
)

type LocationModel struct {
	Id        uint `gorm:"AUTO_INCREMENT"`
	UserId    string
	Longitude float64
	Latitude  float64
	Timestamp string
}

func (LocationModel) TableName() string {
	return "location"
}

func (model LocationModel) GetLocationResponse() *golocopb.LocationResponse {
	return &golocopb.LocationResponse{
		Location: &golocopb.Location{
			Id:        int32(model.Id),
			UserId:    model.UserId,
			Longitude: model.Longitude,
			Latitude:  model.Latitude,
			Timestamp: model.Timestamp,
		}}
}

func GetLocationModel(request *golocopb.LocationRequest, isUpdate bool) LocationModel {
	locationModel := LocationModel{
		UserId:    request.GetLocation().GetUserId(),
		Latitude:  request.GetLocation().GetLatitude(),
		Longitude: request.GetLocation().GetLongitude(),
		Timestamp: request.GetLocation().GetTimestamp(),
	}

	if isUpdate {
		locationModel.Id = uint(request.GetLocation().GetId())
	}
	return locationModel
}

func GetLocationModelFromLocation(request *golocopb.Location) LocationModel {
	locationModel := LocationModel{
		Id:        uint(request.GetId()),
		UserId:    request.GetUserId(),
		Latitude:  request.GetLatitude(),
		Longitude: request.GetLongitude(),
		Timestamp: request.GetTimestamp(),
	}
	return locationModel
}
