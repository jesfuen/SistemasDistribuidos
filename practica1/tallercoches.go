package main

import "fmt"

// Data structures

type Workshop struct {
	Mecanics []Mechanic
	Capacity [2]Car
}

type Mechanic struct {
	Id              int
	Name            string
	Speciality      string
	YearsExperience int
	Status          string
}

type Incidence struct {
	Id          int
	Mecanics    []Mechanic
	Type        string
	Priority    string
	Description string
	Status      string
}

type Car struct {
	RegistrationNumber string
	Brand              string
	Model              string
	EntryDate          string
	DepartureDate      string
	Incidence          Incidence
}

type Client struct {
	Id          int
	Name        string
	PhoneNumber string
	Email       string
	Cars        []Car
}

// Workshop functions

func (w *Workshop) assignCar(c Car) {
	for i := 0; i < len(w.Capacity); i++ {
		if w.Capacity[i].Brand == "" {
			w.Capacity[i] = c
			fmt.Printf("El coche %s %s asignado correctamente al taller\n", c.Brand, c.Model)
			return
		}
	}

	fmt.Println("No hay espacio en el taller.")
}

func (w Workshop) showStatus() {

	// Local vars
	var occupied int = 0
	var free int = 0

	for i := 0; i < len(w.Capacity); i++ {
		if w.Capacity[i].Brand == "" {
			free++
		} else {
			occupied++
		}
	}

	fmt.Printf("Hay %d plazas libres y %d plazas ocupadas en el taller", free, occupied)

}

func main() {

}
