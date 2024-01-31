package data

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name  string
	Buses []Bus
}

type Bus struct {
	gorm.Model
	GroupID      uint
	LicencePlate string
	PhoneNumber  string
	Seats        []Seat
	Shedule      string
}

type Seat struct {
	gorm.Model
	BusID       uint
	SeatNumber  string
	PassengerID *uint
	IsBooked    bool
}

type Passenger struct {
	//TODO: Some How Make It So That PassengerID Cannot Be 0
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement:1"`
	Name        string
	PhoneNumber string
	Seats       []Seat
}
