package model

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/skyezon/parking-lot/common/errors"
)

type Car struct {
	Color       string
	RegisNumber string
}

func NewCar(color, regisNumber string)  (Car, error) {
	color = strings.ReplaceAll(color, " ", "")
	regisNumber = strings.ReplaceAll(regisNumber, " ", "")
	if color == "" {
		return Car{}, fmt.Errorf("color cannot be empty")
	}

	if err := validateRegisNumber(regisNumber); err != nil {
		return Car{}, err
	}

	return Car{
		Color:       color,
		RegisNumber: regisNumber,
	}, nil
}

// regis number validation, assumption format must be : ss-nnnn-sss
// s : string, n : number
func validateRegisNumber(regisNumber string) error {
	if regisNumber == "" {
		return errors.LogErr(errors.VALIDATION_REGIS_NUMBER_ERROR)
	}
	if strings.Count(regisNumber, "-") != 2 {
		return errors.LogErr(errors.VALIDATION_REGIS_NUMBER_ERROR, regisNumber)
	}

	splitted := strings.Split(regisNumber, "-")

	if len(splitted) != 3 {
		return errors.LogErr(errors.VALIDATION_REGIS_NUMBER_ERROR, regisNumber)
	}

	// can be BK, B
	if len(splitted[0]) > 2 || len(splitted[0]) <= 0 {
		return errors.LogErr(errors.VALIDATION_REGIS_NUMBER_ERROR, regisNumber)
	}

	// can be 1234,8080
	if len(splitted[1]) != 4 {
		return errors.LogErr(errors.VALIDATION_REGIS_NUMBER_ERROR, regisNumber)
	}

	for _, el := range splitted[1] {
		if !unicode.IsDigit(el) {
			return errors.LogErr(errors.VALIDATION_REGIS_NUMBER_ERROR, regisNumber)
		}
	}

	//can be A, BC, DEF
	if len(splitted[2]) <= 0 || len(splitted[2]) > 3 {
		return errors.LogErr(errors.VALIDATION_REGIS_NUMBER_ERROR, regisNumber)
	}

	for _, el := range splitted[2] {
		if unicode.IsDigit(el) {
			return errors.LogErr(errors.VALIDATION_REGIS_NUMBER_ERROR, regisNumber)
		}
	}

	return nil
}
