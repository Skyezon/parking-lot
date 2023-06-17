package service

import (
	"fmt"
	"strings"

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
		return "", errors.LogErr(err)
	}
	res := "Slot No. Registration No Colour\n"
	for idx, car := range currLot.Lots {
		if car == (model.Car{}) {
			continue
		}
		res += fmt.Sprintf("%d %s %s\n", idx+1, car.RegisNumber, car.Color)
	}
	return res, nil
}

func GetRegisNumberByColor(color string) (string, error) {
	currLot, err := model.GetLotInstance()
	if err != nil {
		return "", errors.LogErr(err)
	}
	res := ""
	for _, car := range currLot.Lots {
		if !strings.EqualFold(car.Color, color) {
			continue
		}
		res += fmt.Sprintf("%s, ", car.RegisNumber)

	}
	//clean up the ", "
	if res != "" {
		res = res[:len(res)-2]
	}
	return res, nil
}

func GetSlotByColor(color string) (string, error) {
	currLot, err := model.GetLotInstance()
	if err != nil {
		return "", errors.LogErr(err)
	}
	res := ""

	for idx, car := range currLot.Lots {
		if strings.EqualFold(car.Color, color) {
			res += fmt.Sprintf("%d, ", idx+1)
		}
	}

	if res != "" {
		res = res[:len(res)-2]
	}

	return res, nil
}

func GetSlotByRegisNum(regisNumber string) (string, error) {
	currLot, err := model.GetLotInstance()
	if err != nil {
		return "", errors.LogErr(err)
	}
	if len(currLot.Lots) == 0 || len(regisNumber) == 0 {
		return "", errors.LogErr(fmt.Errorf(errors.NOT_FOUND))
	}
	for idx, car := range currLot.Lots {
		if strings.EqualFold(regisNumber, car.RegisNumber) {
			return fmt.Sprintf("%d", idx+1), nil
		}
	}
	return "", errors.LogErr(fmt.Errorf(errors.NOT_FOUND))
}
