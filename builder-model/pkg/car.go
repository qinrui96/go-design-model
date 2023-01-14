package pkg

type CarType int

const (
	Convertible CarType = iota
	SUV
	Roadster
)

type Car struct {
	Logo    string
	CarType CarType
	Price   int
	Seat    int
	Color   string
	Stereo  string
}

type ChangeCarConfigurators func(car *Car)

func NewCar(logo string, carType CarType, price int, opts ...ChangeCarConfigurators) *Car {
	myCar := &Car{
		Logo:    logo,
		CarType: carType,
		Price:   price,
	}
	for _, opt := range opts {
		opt(myCar)
	}
	return myCar
}
