package main

import (
	"fmt"
	"github.com/qinrui/go-design-model/builder-model/pkg"
)

func main() {
	opts := make([]pkg.ChangeCarConfigurators, 0)
	opts = append(opts, RedColor, SetSeats)
	car := pkg.NewCar("BMW", pkg.SUV, 1000, opts...)
	fmt.Println(*car)
}

func RedColor(car *pkg.Car) {
	car.Color = "red"
}

func SetSeats(car *pkg.Car) {
	car.Seat = 4
}
