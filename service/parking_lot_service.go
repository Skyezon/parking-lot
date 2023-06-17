package service

import (
	"fmt"

	"github.com/skyezon/parking-lot/common/errors"
	"github.com/skyezon/parking-lot/db/model"
)

func CreateParkingLot(totalLot int) error {
	err := model.NewParkingLot(totalLot)
	if err != nil {
		return errors.LogErr(err)
	}
	return nil
}

func ParkParkingLot(regisNumber string, color string) (int, error) {
	currLot, err := model.GetLotInstance()
	if err != nil {
		return 0, errors.LogErr(err)
	}

	car, err := model.NewCar(color, regisNumber)
	if err != nil {
		return 0, errors.LogErr(err)
	}

	absoluteSlot, err := currLot.Park(car)
	if err != nil {
		return 0, errors.LogErr(err)
	}
	return absoluteSlot + 1, nil
}

func LeaveParkingLot(absoluteSlotNumber int) error {
	currLot, err := model.GetLotInstance()
	if err != nil {
		return errors.LogErr(err)
	}

	err = currLot.Leave(absoluteSlotNumber)
	if err != nil {
		return errors.LogErr(err)
	}

	return nil
}

func StatusParkingLot() (string, error) {
	currLot, err := model.GetLotInstance()
	if err != nil {
		return "",errors.LogErr(err)
	}
    res := "Slot No. Registration No Colour\n"
    for idx,car := range currLot.Lots{
        if car == (model.Car{}){
            continue
        }
        res += fmt.Sprintf("%d %s %s\n",idx+1,car.RegisNumber,car.Color)
    }
	return res,nil
}
