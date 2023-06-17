package errors

import (
	"fmt"
	"log"
	"runtime"
)

const PARKING_LOT_FULL_ERR = "Sorry, parking lot is full"
const NOT_FOUND = "Not found"
const UNIT_TEST_ERR_TEMPLATE = "expected err not match, expected : %v, actual : %v"
const PARKING_LOT_IS_NOT_INTIALIZED = "Please initialize Lot first"

var VALIDATION_REGIS_NUMBER_ERROR = fmt.Errorf("Registration number is invalid")

//easier error tracking
func LogErr(err error, msgToDev ...string) error {
	pc := make([]uintptr, 5)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	log.Println(err)
	if len(msgToDev) > 0 {
		log.Printf("debug msg : %s", msgToDev)
	}
	for {
		frame, more := frames.Next()
		log.Printf("%s:%d", frame.File, frame.Line)
		if !more {
			break
		}
	}
	fmt.Println("")
	return err
}
