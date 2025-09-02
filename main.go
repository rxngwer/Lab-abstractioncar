// Lab: Abstraction with Go (Car Example)
// --------------------------------------
// This lab demonstrates abstraction via the Car interface.
// Each concrete type (Sedan, Pickup, EV) implements the same
// behaviors defined by the interface, allowing polymorphic use.
//
// Run:
//   go run main.go
//
// Author: You 🧑‍💻
package main

import "fmt"

// Car is the abstraction (interface) that defines common behaviors.
type Car interface {
    Drive()              // behavior when the car is driven
    Info() string        // formatted car info
}

// Sedan represents a passenger car.
type Sedan struct {
    Brand string
    Model string
    Color string
}

func (s Sedan) Drive() {
    fmt.Println("Driving sedan ->", s.Brand, s.Model, "Color:", s.Color)
}

func (s Sedan) Info() string {
    return fmt.Sprintf("Sedan(%s %s, %s)", s.Brand, s.Model, s.Color)
}

// Pickup represents a pickup truck.
type Pickup struct {
    Brand   string
    Payload int // kg
    FourWD  bool
}

func (p Pickup) Drive() {
    mode := "2WD"
    if p.FourWD {
        mode = "4WD"
    }
    fmt.Println("Driving pickup ->", p.Brand, "Payload:", p.Payload, "kg", "| Mode:", mode)
}

func (p Pickup) Info() string {
    return fmt.Sprintf("Pickup(%s, %dkg, FourWD:%t)", p.Brand, p.Payload, p.FourWD)
}

// EV represents an electric vehicle.
type EV struct {
    Brand    string
    Battery  int // kWh
    RangeKM  int // km per charge
    FastCharge bool
}

func (e EV) Drive() {
    fc := "Standard charge"
    if e.FastCharge {
        fc = "Fast charge ready"
    }
    fmt.Println("Driving EV ->", e.Brand, "| Battery:", e.Battery, "kWh | Range:", e.RangeKM, "km |", fc)
}

func (e EV) Info() string {
    return fmt.Sprintf("EV(%s, %dkWh, %dkm, FastCharge:%t)", e.Brand, e.Battery, e.RangeKM, e.FastCharge)
}

func main() {
    // Create concrete cars
    sedan := Sedan{Brand: "Toyota", Model: "Camry", Color: "Black"}
    pickup := Pickup{Brand: "Isuzu", Payload: 1000, FourWD: true}
    ev := EV{Brand: "Tesla", Battery: 75, RangeKM: 500, FastCharge: true}

    // Use abstraction via the Car interface
    cars := []Car{sedan, pickup, ev}

    for _, car := range cars {
        fmt.Println(car.Info())
        car.Drive()
        fmt.Println("---")
    }
}
