package main

import (
	"errors"
	"fmt"
)

type Vehicle struct {
	Brand string
	Model string
	Year int
	Color string
}

type VehicleError struct {
	VehicleType string
	Err error
}


func (ve VehicleError) Error() string {
	return fmt.Sprintf("%s vehicle error: %v", ve.VehicleType, ve.Err)
}

type VehicleInterface interface {
	Start()
	Stop()
	Steer()
}

type Car struct {
	Vehicle
	NumDoors int
	EngineType string
}

type Boat struct {
	Vehicle
	Length int
	PropulsionType string
}

type Motorcycle struct {
	Vehicle
	NumWheels int
	HasSidecar bool
}
