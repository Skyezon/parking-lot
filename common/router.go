package common

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/skyezon/parking-lot/handler"
)

func Serve() {
	router := chi.NewRouter()

	router.Post("/park/{platNumber}/{color}", handler.ParkHandler)

	router.Post("/create_parking_lot/{totalParkingLot}", handler.CreateParkingHandler)

	router.Post("/leave/{slotNumber}", handler.LeaveHandler)

	router.Get("/status", handler.StatusHandler)

	router.Get("/cars_registration_numbers/colour/{color}", handler.FindRegisNumberByColorHandler)

	router.Get("/cars_slot/colour/{color}", handler.FindCarSlotsByColorHandler)

	router.Get("/slot_number/car_registration_number/{regisNumber}", handler.FindSlotNumberbyRegisNumberHandler)

	router.Post("/bulk", handler.BulkCommandHandler)

	fmt.Println("Server running on port :8080")

	http.ListenAndServe(":8080", router)
}
