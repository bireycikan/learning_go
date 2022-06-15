package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	type Sample struct {
		field string
		a, b  int
	}

	data := Sample{"word", 1, 2}
	// or
	data2 := Sample{
		field: "hello world",
		a:     3,
		b:     4,
	}
	// or
	data3 := Sample{}
	// or
	var data4 struct {
		field string
		a, b  float32
	}
	// or
	data5 := struct {
		field string
		a, b  float32
	}{
		"hello",
		7.7, 8.8,
	}
	fmt.Println("data: ", data)
	fmt.Println("data2: ", data2)
	fmt.Println("data3: ", data3)
	fmt.Println("data3 a: ", data3.a)
	fmt.Println("data4: ", data4)
	fmt.Println("data5: ", data5)

	jessica := Passenger{"Jessica", 1, false}
	fmt.Println(jessica)

	var (
		michael = Passenger{Name: "Michael", TicketNumber: 2}
		kate    = Passenger{Name: "Kate", TicketNumber: 3}
	)

	fmt.Println(michael, kate)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	jessica.Boarded = true
	michael.Boarded = true
	if michael.Boarded {
		fmt.Println("Michael boarded the bus")
	}
	if jessica.Boarded {
		fmt.Println(jessica.Name, "has boarded the bus")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
