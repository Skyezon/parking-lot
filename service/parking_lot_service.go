package service

import (
	"fmt"
	"strconv"
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
    for i:= 0 ;i < currLot.TotalLot ; i++ {
        if currLot.Lots[i] == (model.Car{}){
            continue
        }
		res += fmt.Sprintf("%d %s %s", i+1, currLot.Lots[i].RegisNumber, currLot.Lots[i].Color)
        if i != currLot.TotalLot -1 {
            res += "\n"
        }

    }
	return res, nil
}

func GetRegisNumberByColor(color string) (string, error) {
    if color == ""{
        return "", errors.LogErr(fmt.Errorf("color cannot be empty"))
    }
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

func BulkCommander(payload string) (string) {
	res := ""
	commands := strings.Split(payload, "\n")
	for idx, command := range commands {
		tempRes, err := executeCommand(command)
		if err == nil {
			res += tempRes 
		} else {
			res += err.Error() 
		}
        if idx != len(commands) -1 {
            res += "\n"
        } 
	}
    return res
}

func executeCommand(onelinePayload string) (string, error) {
	splitted := strings.Split(onelinePayload, " ")
	if len(splitted) == 0 {
		return "", errors.LogErr(errors.INSUFFICIENT_PARAMETER)
	}
    for idx,split := range splitted{
        splitted[idx] = strings.TrimSpace(split)
    }

	switch splitted[0] {
	case "create_parking_lot":
		if len(splitted) < 2 {
			return "", errors.LogErr(errors.INSUFFICIENT_PARAMETER)
		}
		totalLot := splitted[1]
		totalLotInt, err := strconv.Atoi(totalLot)
		if err != nil {
			return "", errors.LogErr(err)
		}
		err = CreateParkingLot(totalLotInt)
		if err != nil {
			return "", errors.LogErr(err)
		}
		return fmt.Sprintf("Created a parking lot with %d slots", totalLotInt), nil
	case "park":
		if len(splitted) < 3 {
			return "", errors.LogErr(errors.INSUFFICIENT_PARAMETER)
		}
		regisNum := splitted[1]
		color := splitted[2]
		slot, err := ParkParkingLot(regisNum, color)
		if err != nil {
			return "", errors.LogErr(err)
		}
		return fmt.Sprintf("Allocated slot number: %d", slot), nil
	case "leave":
		if len(splitted) < 2 {
			return "", errors.LogErr(errors.INSUFFICIENT_PARAMETER)
		}
		slot := splitted[1]
		slotInt, err := strconv.Atoi(slot)
		if err != nil {
			return "", errors.LogErr(fmt.Errorf("invalid slot number"))
		}
		err = LeaveParkingLot(slotInt - 1)
		if err != nil {
			return "", errors.LogErr(err)
		}
		return fmt.Sprintf("Slot number %d is free", slotInt), nil
	case "status":
		res, err := StatusParkingLot()
		if err != nil {
			return "", errors.LogErr(err)
		}
		return res , nil
	case "registration_numbers_for_cars_with_colour":
		if len(splitted) < 2 {
			return "", errors.LogErr(errors.INSUFFICIENT_PARAMETER)
		}
		color := splitted[1]
		res, err := GetRegisNumberByColor(color)
		if err != nil {
			return "", errors.LogErr(err)
		}
		return res , nil
	case "slot_number_for_registration_number":
		if len(splitted) < 2 {
			return "", errors.LogErr(errors.INSUFFICIENT_PARAMETER)
		}
		regisNumber := splitted[1]
		res, err := GetSlotByRegisNum(regisNumber)
		if err != nil {
			return "", errors.LogErr(err)
		}
		return res , nil
	case "slot_numbers_for_cars_with_colour":
		if len(splitted) < 2 {
			return "", errors.LogErr(errors.INSUFFICIENT_PARAMETER)
		}
		color := splitted[1]
		res, err := GetSlotByColor(color)
		if err != nil {
			return "", errors.LogErr(errors.INSUFFICIENT_PARAMETER)
		}
		return res , nil
	default:
		return "", nil
	}
}
