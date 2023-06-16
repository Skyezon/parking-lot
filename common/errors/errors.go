package errors

import (
	"fmt"
	"log"
	"runtime"
)

const PARKING_LOT_FULL_ERR = "Sorry, parking lot is full"
const REGIS_NUMBER_NOT_FOUND = "Not found"
const UNIT_TEST_ERR_TEMPLATE = "expected err not match, expected : %v, actual : %t"

var VALIDATION_REGIS_NUMBER_ERROR = fmt.Errorf("Registration number is invalid")

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
		log.Printf("%s:%d %s", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
	fmt.Println("")
	return err
}
