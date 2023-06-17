package service

import (
	"github.com/skyezon/parking-lot/common/errors"
	"github.com/skyezon/parking-lot/db/model"
)

func CreateParkingLot(totalLot int)error{
    err := model.NewParkingLot(totalLot)
    if err != nil {
        return errors.LogErr(err)
    }
    return nil
}
