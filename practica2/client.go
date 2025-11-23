package main

import "fmt"

// Client functions

func (c *Client) create(w *Workshop) {
	fmt.Print("Introduzca el nombre del cliente: ")
	fmt.Scanln(&c.Name)
	fmt.Print("Introduzca el telefono del cliente: ")
	fmt.Scanln(&c.PhoneNumber)
	fmt.Print("Introduzca el email del cliente: ")
	fmt.Scanln(&c.Email)

	var registrationNumber string
	fmt.Print("Introduzca la matricula del vehiculo: ")
	fmt.Scanln(&registrationNumber)

	// Buscar el coche en el workshop
	car := w.findCarByRegistration(registrationNumber)
	if car != nil {
		c.Cars = append(c.Cars, *car)
		fmt.Println("Vehículo asociado correctamente.")
	} else {
		vehicleNotFound(registrationNumber)
	}
}

func (c Client) display() {
	fmt.Printf("\n===INFORMACION DEL CLIENTE===\n")
	fmt.Printf("ID: %d\n", c.Id)
	fmt.Printf("Nombre: %s\n", c.Name)
	fmt.Printf("Telefono: %s\n", c.PhoneNumber)
	fmt.Printf("Email: %s\n", c.Email)

	fmt.Printf("Vehiculos asociados: ")
	if len(c.Cars) == 0 {
		fmt.Println(" - Sin vehiculos asociados.")
	} else {
		for _, car := range c.Cars {
			fmt.Printf(" - %s %s (Matricula: %s)\n", car.Brand, car.Model, car.RegistrationNumber)
		}
	}
}

func (c *Client) modify() {

	var option int

	fmt.Printf("\n===MODIFICAR CLIENTE===\n")
	fmt.Printf("1- Modificar nombre")
	fmt.Printf("2- Modificar email")
	fmt.Printf("3- Modificar telefono")
	fmt.Print("Opcion: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		fmt.Print("Introduzca el nuevo nombre: ")
		fmt.Scan(&c.Name)
		fmt.Printf("Nombre actualizado: %s\n", c.Name)
	case 2:
		fmt.Print("Introduzca el nuevo telefono: ")
		fmt.Scan(&c.PhoneNumber)
		fmt.Printf("Teléfono actualizado: %s\n", c.PhoneNumber)
	case 3:
		fmt.Print("Introduzca el nuevo email: ")
		fmt.Scan(&c.Email)
		fmt.Printf("Email actualizado: %s\n", c.Email)
	default:
		fmt.Println("Opción inválida.")
	}
}

func (c *Client) addCar(car Car) {
	c.Cars = append(c.Cars, car)
	fmt.Printf("Vehículo %s añadido al cliente %s\n", car.RegistrationNumber, c.Name)
}

func (c Client) listCars() {
	if len(c.Cars) == 0 {
		fmt.Println("El cliente no tiene vehículos asociados.")
		return
	}
	fmt.Printf("\n=== VEHÍCULOS DE %s ===\n", c.Name)
	for _, car := range c.Cars {
		fmt.Printf("- %s %s (Matrícula: %s)\n", car.Brand, car.Model, car.RegistrationNumber)
	}
	fmt.Println()
}
