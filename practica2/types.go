package main

// Data structures

type Workshop struct {
	Mecanics []Mechanic
	Cars     []Car
	Clients  []Client
}

type Mechanic struct {
	Id              int
	Name            string
	Speciality      string
	YearsExperience int
	Status          string
	Activity        string
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
	Incidences         []Incidence
}

type Client struct {
	Id          int
	Name        string
	PhoneNumber string
	Email       string
	Cars        []Car
}
