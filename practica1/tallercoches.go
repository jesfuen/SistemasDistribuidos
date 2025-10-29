package main

import (
	"fmt"
)

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

// Generic functions

func remove[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	copy(s[i:], s[i+1:])
	var cero T
	s[len(s)-1] = cero
	return s[:len(s)-1]
}

// Workshop functions

func (w *Workshop) assignCar(c Car) {
	for i := range w.Cars {
		if w.Cars[i].Brand == "" {
			w.Cars[i] = c
			fmt.Printf("El coche %s %s asignado correctamente al taller\n", c.Brand, c.Model)
			return
		}
	}

	fmt.Println("No hay espacio en el taller.")
}

func (w Workshop) showStatus() {

	var occupied int = 0
	var free int = 0

	for i := range w.Cars {
		if w.Cars[i].Brand == "" {
			free++
		} else {
			occupied++
		}
	}

	fmt.Printf("Hay %d plazas libres y %d plazas ocupadas en el taller\n", free, occupied)

}

// Client manage functions

func (w *Workshop) createClient() {

	var matricula string
	var c Client

	fmt.Print("Introduzca el nombre del cliente: ")
	fmt.Scanln(&c.Name)
	fmt.Print("Introduzca el telefono del cliente: ")
	fmt.Scanln(&c.PhoneNumber)
	fmt.Print("Introduzca el email del cliente: ")
	fmt.Scanln(&c.Email)
	fmt.Print("Introduzca la matricula del vehiculo: ")
	fmt.Scanln(&matricula)

	var carExists bool = false

	for i := range w.Cars {
		if w.Cars[i].RegistrationNumber == matricula {
			c.Cars = append(c.Cars, w.Cars[i])
			carExists = true
			break
		}
	}

	if !carExists {
		fmt.Printf("El coche con matricula %s no existe en el taller.", matricula)
	} else {
		fmt.Println("Cliente creado correctamente.")
		c.Id = 1 + len(w.Clients)
		w.Clients = append(w.Clients, c)
	}

}

func (w *Workshop) showClientInfo(id int) {
	var found bool
	for _, c := range w.Clients {
		if c.Id == id {
			fmt.Printf("Cliente: %s\nTelefono: %s\nEmail: %s\nId: %d\n", c.Name, c.PhoneNumber, c.Email, c.Id)
			fmt.Println("Coches asociados:")
			for _, car := range c.Cars {
				fmt.Printf("- %s %s (Matricula: %s)\n", car.Brand, car.Model, car.RegistrationNumber)
			}
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Cliente no encontrado.")
	}
}

func (w *Workshop) removeClient() {

	var id int
	var found bool

	fmt.Print("Introduzca el id del cliente: ")
	fmt.Scanf("%d", &id)

	for i := range w.Clients {
		if w.Clients[i].Id == id {
			w.Clients = remove(w.Clients, i)
			fmt.Println("Cliente eliminado correctamente")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Id incorrecto. No se encontro el cliente")
	}
}

func (w *Workshop) modifyClient() {
	var id int
	var option int
	var found bool

	fmt.Print("Introduzca el id del cliente: ")
	fmt.Scanf("%d", &id)

	for i := range w.Clients {
		if w.Clients[i].Id == id {
			found = true
			fmt.Println("Selecciona una opcion: ")
			fmt.Println("1- Modificar nombre")
			fmt.Println("2- Modificar telefono")
			fmt.Println("3- Modificar email")
			fmt.Print("Opcion: ")
			fmt.Scanf("%d", &option)

			switch option {
			case 1:
				fmt.Print("Introduzca el nombre: ")
				fmt.Scan(&w.Clients[i].Name)
				fmt.Printf("Nombre cambiado correctamente: %s\n", w.Clients[i].Name)
			case 2:
				fmt.Print("Introduzca el telefono: ")
				fmt.Scan(&w.Clients[i].PhoneNumber)
				fmt.Printf("Telefono cambiado correctamente: %s\n", w.Clients[i].PhoneNumber)
			case 3:
				fmt.Print("Introduzca el email: ")
				fmt.Scan(&w.Clients[i].Email)
				fmt.Printf("Email cambiado correctamente: %s\n", w.Clients[i].Email)
			}

			break
		}
	}

	if !found {
		fmt.Println("Id incorrecto. Cliente no encontrado")
	}
}

// Car manage functions

// Main function

func main() {

	workshop := Workshop{}

	workshop.Mecanics = []Mechanic{
		{Id: 1, Name: "Juan Perez", Speciality: "Motor", YearsExperience: 5, Status: "Disponible"},
		{Id: 2, Name: "Ana Gomez", Speciality: "Electricidad", YearsExperience: 3, Status: "Disponible"},
		{Id: 3, Name: "Luis Martinez", Speciality: "Neumáticos", YearsExperience: 4, Status: "Disponible"},
	}

	workshop.Cars = make([]Car, 2*len(workshop.Mecanics))

	car1 := Car{
		RegistrationNumber: "1234ABC",
		Brand:              "Toyota",
		Model:              "Corolla",
		EntryDate:          "2024-06-01",
	}
	workshop.assignCar(car1)

	workshop.showStatus()

	workshop.createClient()

	workshop.showClientInfo(1)

	fmt.Printf("El taller tiene %d clientes registrados.\n", len(workshop.Clients))

	workshop.modifyClient()

	workshop.showClientInfo(1)

	workshop.removeClient()

	fmt.Printf("El taller tiene %d clientes registrados.\n", len(workshop.Clients))

	workshop.showStatus()
}
