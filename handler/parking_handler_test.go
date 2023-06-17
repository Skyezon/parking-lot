package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/skyezon/parking-lot/common/errors"
	"github.com/skyezon/parking-lot/db/model"
	"github.com/skyezon/parking-lot/service"
)

func Test_CreateParkingHandler(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		expectedBody string
	}{
		{
			name:         "valid",
			url:          "/create_parking_lot/10",
			expectedBody: "Created a parking lot with 10 slots",
		},
		{
			name:         "invalid too low",
			url:          "/create_parking_lot/-1",
			expectedBody: "Maximum lot number is invalid",
		},
		{
			name:         "invalid not a number",
			url:          "/create_parking_lot/a",
			expectedBody: "Maximum lot number is invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := chi.NewRouter()
			r.Post("/create_parking_lot/{totalParkingLot}", CreateParkingHandler)
			req := httptest.NewRequest("POST", tt.url, nil)
			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			if rr.Body.String() != tt.expectedBody {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.expectedBody, rr.Body.String())
			}
		})
	}
}

func Test_ParkHandler(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		expectedBody string
	}{
		{
			name:         "valid",
			url:          "/park/bk-1234-abc/Red",
			expectedBody: "Allocated slot number: 1",
		},
		{
			name:         "invalid registration number",
			url:          "/park/bk--abc/red",
			expectedBody: "Registration number is invalid",
		},
		{
			name:         "invalid maximum lot",
			url:          "/park/DD-8080-YUK/Brown",
			expectedBody: "Sorry, parking lot is full",
		},
	}
	service.CreateParkingLot(1)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", tt.url, nil)
			rr := httptest.NewRecorder()
			r := chi.NewRouter()
			r.Post("/park/{platNumber}/{color}", ParkHandler)
			r.ServeHTTP(rr, req)
			if rr.Body.String() != tt.expectedBody {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.expectedBody, rr.Body.String())
			}
		})
	}
}

func Test_LeaveHandler(t *testing.T) {
	tests := []struct {
		name         string
		slotToLeave  string
		expectedBody string
	}{
		{
			name:         "valid",
			slotToLeave:  "1",
			expectedBody: "Slot number 1 is free",
		},
		{
			name:         "invalid over max",
			slotToLeave:  "2",
			expectedBody: "slot number invalid",
		},
		{
			name:         "invalid too low",
			slotToLeave:  "-1",
			expectedBody: "slot number invalid",
		},
		{
			name:         "invalid not number",
			slotToLeave:  "a",
			expectedBody: "slot number invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", fmt.Sprintf("/leave/%s", tt.slotToLeave), nil)
			rr := httptest.NewRecorder()
			r := chi.NewRouter()
			r.Post("/leave/{slotNumber}", LeaveHandler)
			r.ServeHTTP(rr, req)
			if rr.Body.String() != tt.expectedBody {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.expectedBody, rr.Body.String())
			}
		})
	}
}

func Test_StatusHandler(t *testing.T) {
	tests := []struct {
		name         string
		setup        func()
		expectedBody string
	}{
		{
			name: "invalid",
			setup: func() {

			},
			expectedBody: `Slot No. Registration No Colour
`,
		},
		{
			name: "valid normal",
			setup: func() {
				service.CreateParkingLot(10)
				service.ParkParkingLot("Bk-1234-abc", "hijau")
				service.ParkParkingLot("Bk-2312-cde", "merah")
			},
			expectedBody: `Slot No. Registration No Colour
1 Bk-1234-abc hijau
2 Bk-2312-cde merah
`,
		},
		{
			name: "valid removed",
			setup: func() {
				service.CreateParkingLot(10)
				service.ParkParkingLot("Bk-1234-abc", "hijau")
				service.ParkParkingLot("Bk-2312-cde", "merah")
				service.LeaveParkingLot(1)
			},
			expectedBody: `Slot No. Registration No Colour
1 Bk-1234-abc hijau
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			r := chi.NewRouter()
			r.Get("/status", StatusHandler)
			req, _ := http.NewRequest("GET", "/status", nil)
			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)
			if rr.Body.String() != tt.expectedBody {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.expectedBody, rr.Body.String())
			}

		})
	}
}

func Test_FindRegisNumberByColorHandler(t *testing.T) {
	tests := []struct {
		name         string
		color        string
		setup        func()
		expectedBody string
	}{
		{
			name:  "valid",
			color: "blue",
			setup: func() {
				service.CreateParkingLot(10)
				service.ParkParkingLot("bk-1234-abc", "blue")
			},
			expectedBody: "bk-1234-abc",
		},
		{
			name:  "valid case insensitive",
			color: "blue",
			setup: func() {
				service.CreateParkingLot(10)
				service.ParkParkingLot("bk-1234-abc", "blue")

			},
			expectedBody: "bk-1234-abc",
		},
		{
			name:  "multi",
			color: "blue",
			setup: func() {
				service.CreateParkingLot(10)
				service.ParkParkingLot("bk-1234-abc", "Blue")
				service.ParkParkingLot("ac-1234-abc", "blue")
			},
			expectedBody: "bk-1234-abc, ac-1234-abc",
		},
		{
			name:  "not found",
			color: "blue",
			setup: func() {
				service.CreateParkingLot(10)
				service.ParkParkingLot("bk-1234-abc", "Gray")
				service.ParkParkingLot("bk-1234-kbc", "Yello")
			},
			expectedBody: "",
		},
		{
			name:  "lot not initialized",
			color: "red",
			setup: func() {
				model.GlobalLot = nil
			},
			expectedBody: "Please initialize Lot first",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			r := chi.NewRouter()
			r.Get("/cars_registration_numbers/colour/{color}", FindRegisNumberByColorHandler)
			req, _ := http.NewRequest("GET", fmt.Sprintf("/cars_registration_numbers/colour/%s", tt.color), nil)
			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)
			if rr.Body.String() != tt.expectedBody {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.expectedBody, rr.Body.String())
			}

		})
	}
}

func Test_FindCarSlotsByColorHandler(t *testing.T) {
	tests := []struct {
		name         string
		color        string
		setup        func()
		expectedBody string
	}{
		{
			name:  "valid",
			color: "blue",
			setup: func() {
				service.CreateParkingLot(10)
				service.ParkParkingLot("bk-1234-abc", "blue")
			},
			expectedBody: "1",
		},
		{
			name:  "valid case insensitive",
			color: "blue",
			setup: func() {
				service.CreateParkingLot(10)
				service.ParkParkingLot("bk-1234-abc", "blue")

			},
			expectedBody: "1",
		},
		{
			name:  "multi",
			color: "blue",
			setup: func() {
				service.CreateParkingLot(10)
				service.ParkParkingLot("bk-1234-abc", "Blue")
				service.ParkParkingLot("ac-1234-abc", "blue")
			},
			expectedBody: "1, 2",
		},
		{
			name:  "not found",
			color: "blue",
			setup: func() {
				service.CreateParkingLot(10)
				service.ParkParkingLot("bk-1234-abc", "Gray")
				service.ParkParkingLot("bk-1234-kbc", "Yello")
			},
			expectedBody: "",
		},
        {
			name:        "lot not initialized",
			color: "tes",
			setup: func() {
				model.GlobalLot = nil
			},
			expectedBody: "Please initialize Lot first",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			r := chi.NewRouter()
			r.Get("/cars_slot/colour/{color}", FindCarSlotsByColorHandler)
			req, _ := http.NewRequest("GET", fmt.Sprintf("/cars_slot/colour/%s", tt.color), nil)
			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)
			if rr.Body.String() != tt.expectedBody {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.expectedBody, rr.Body.String())
			}

		})
	}
}

func Test_FindSlotNumberByRegisNumberHandler(t *testing.T) {
	tests := []struct {
		name         string
		regisNumber  string
		setup        func()
		expectedBody string
	}{
		{
			name:        "valid",
			regisNumber: "bk-1234-abc",
			setup: func() {
				service.CreateParkingLot(10)
				service.ParkParkingLot("bk-1234-abc", "blue")
			},
			expectedBody: "1",
		},
		{
			name:        "valid case insensitive",
			regisNumber: "bk-1234-abC",
			setup: func() {
				service.CreateParkingLot(10)
				service.ParkParkingLot("bK-1234-abc", "blue")

			},
			expectedBody: "1",
		},
		{
			name:        "not found",
			regisNumber: "bk-1233-abc",
			setup: func() {
				service.CreateParkingLot(10)
				service.ParkParkingLot("bk-1234-abc", "Gray")
				service.ParkParkingLot("bk-1234-kbc", "Yello")
			},
			expectedBody: "Not found",
		},
		{
			name:        "lot not initialized",
			regisNumber: "tes",
			setup: func() {
				model.GlobalLot = nil
			},
			expectedBody: "Please initialize Lot first",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			r := chi.NewRouter()
			r.Get("/slot_number/car_registration_number/{regisNumber}", FindSlotNumberbyRegisNumberHandler)
			req, _ := http.NewRequest("GET", fmt.Sprintf("/slot_number/car_registration_number/%s", tt.regisNumber), nil)
			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)
			if rr.Body.String() != tt.expectedBody {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.expectedBody, rr.Body.String())
			}

		})
	}
}

func Test_BulkHandler(t *testing.T) {
	tests := []struct {
		name         string
		payload      string
		expectedBody string
	}{
		{
			name: "valid",
			payload: `create_parking_lot 6
park B-1234-RFS Black
park B-1999-RFD Green
park B-1000-RFS Black
park B-1777-RFU BlueSky
park B-1701-RFL Blue
park B-1141-RFS Black
leave 4
status
park B-1333-RFS Black
park B-1989-RFU BlueSky
registration_numbers_for_cars_with_colour Black
slot_numbers_for_cars_with_colour Black
slot_number_for_registration_number B-1701-RFL
slot_number_for_registration_number RI-1
`,

			expectedBody: `Created a parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Allocated slot number: 4
Allocated slot number: 5
Allocated slot number: 6
Slot number 4 is free
Slot No. Registration No Colour
1 B-1234-RFS Black
2 B-1999-RFD Green
3 B-1000-RFS Black
5 B-1701-RFL Blue
6 B-1141-RFS Black
Allocated slot number: 4
Sorry, parking lot is full
B-1234-RFS, B-1000-RFS, B-1333-RFS, B-1141-RFS
1, 3, 4, 6
5
Not found
`,
		},
		{
			name: "error",
			payload: `create_parking_lot 6
park B-1234-RFS Black
park B-1999-RFDb Green`,
			expectedBody: `Created a parking lot with 6 slots
Allocated slot number: 1
Registration number is invalid`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := chi.NewRouter()
			r.Post("/bulk", BulkCommandHandler)
			rr := httptest.NewRecorder()
			reqBody := bytes.NewBufferString(tt.payload)
			req, _ := http.NewRequest("POST", "/bulk", reqBody)
			r.ServeHTTP(rr, req)

			if rr.Body.String() != tt.expectedBody {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.expectedBody, rr.Body.String())
			}
		})
	}
}
