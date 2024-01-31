package data

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

var DB *gorm.DB

func InitDatabase() {
	command := "host=localhost user=zed password=zed dbname=ticket_booking port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(command), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	DB = db
	db.AutoMigrate(&Passenger{}, &Group{}, &Bus{}, &Seat{})

}
func Populate() {
	groupA := Group{
		Name: "Baba Rides",
	}
	DB.Create(&groupA)
	for i := 1; i <= 5; i++ {
		bus := Bus{
			GroupID:      groupA.ID,
			LicencePlate: fmt.Sprintf("LP%d", i),
			PhoneNumber:  fmt.Sprintf("PN%d", i),
			Shedule:      "12:00am - 2:00pm",
		}
		tx := DB.Begin()
		tx.Create(&bus)
		tx.Commit()
		for j := 0; j < i; j++ {
			seat := Seat{
				BusID:       bus.ID,
				SeatNumber:  strconv.Itoa(j + 1),
				IsBooked:    false,
				PassengerID: nil,
			}
			DB.Create(&seat)
		}
	}
}
