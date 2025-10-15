package main

type Mechanic struct {
	Id              int
	Name            string
	Speciality      string
	YearsExperience int
	Status          string
}

type Incidence struct {
	Id          int
	Mecanic     Mechanic
	Type        string
	Priority    string
	Description string
	Status      string
}

type Car struct {
	RegistrationNumber int
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
