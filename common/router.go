package common

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/skyezon/parking-lot/handler"
)

func Router(){
    router:= chi.NewRouter()

    router.Get("/",handler.HelloWorldHandler)

    router.Post("/park/{platNumber}/{color}",handler.ParkHandler)
    
    router.Post("/create_parking_lot/{totalParkingLot}",handler.CreateParkingHandler)

    router.Post("/leave/{slotNumber}",handler.LeaveHandler)

    router.Get("/status",handler.StatusHandler)

    router.Get("/cars_registration/colour/{color}",handler.FindRegisNumberByColor)

    router.Get("/cars_slot/colour/{color}", handler.FindCarSlotsByColor)

    router.Get("/slot_number/car_registration_number/{regisNumber}",handler.FindSlotNumberbyRegisNumber)

    http.ListenAndServe(":8080",router)
}
