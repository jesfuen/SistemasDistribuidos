package main

import "fmt"

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
	Car         Car
}

func (w *Workshop) assigncar(c Car) {
	for i := 0; i < len(w.Capacity); i++ {
		if w.Capacity[i].Brand == "" {
			w.Capacity[i] = c
			fmt.Printf("El coche %s %s asignado correctamente al taller\n", c.Brand, c.Model)
			return
		}
	}

	fmt.Println("No hay espacio en el taller.")
}

func main() {

}
