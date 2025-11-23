package main

import "fmt"

// Car  functions

func (c *Car) create() {
	fmt.Print("Introduzca la matricula del vehiculo: ")
	fmt.Scan(&c.RegistrationNumber)
	fmt.Print("Introduzca la marca del vehiculo: ")
	fmt.Scan(&c.Brand)
	fmt.Print("Introduzca el modelo del vehiculo: ")
	fmt.Scan(&c.Model)
	fmt.Print("Introduzca la fecha de entrada (DD/MM/YYYY): ")
	fmt.Scan(&c.EntryDate)
	fmt.Print("Introduzca la fecha estimada de salida (DD/MM/YYYY): ")
	fmt.Scan(&c.DepartureDate)
}

func (c *Car) modify() {
	var option int

	fmt.Println("\n=== MODIFICAR VEHÍCULO ===")
	fmt.Println("1- Modificar matricula")
	fmt.Println("2- Modificar marca")
	fmt.Println("3- Modificar modelo")
	fmt.Println("4- Modificar fecha de entrada")
	fmt.Println("5- Modificar fecha de salida")
	fmt.Print("Opción: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		fmt.Print("Nueva matrícula: ")
		fmt.Scan(&c.RegistrationNumber)
	case 2:
		fmt.Print("Nueva marca: ")
		fmt.Scan(&c.Brand)
	case 3:
		fmt.Print("Nuevo modelo: ")
		fmt.Scan(&c.Model)
	case 4:
		fmt.Print("Nueva fecha de entrada: ")
		fmt.Scan(&c.EntryDate)
	case 5:
		fmt.Print("Nueva fecha de salida: ")
		fmt.Scan(&c.DepartureDate)
	default:
		fmt.Println("Opción inválida.")
	}
}

func (c Car) display() {
	fmt.Printf("\n=== INFORMACIÓN DEL VEHÍCULO ===\n")
	fmt.Printf("Matrícula: %s\n", c.RegistrationNumber)
	fmt.Printf("Marca: %s\n", c.Brand)
	fmt.Printf("Modelo: %s\n", c.Model)
	fmt.Printf("Fecha de entrada: %s\n", c.EntryDate)
	fmt.Printf("Fecha estimada de salida: %s\n", c.DepartureDate)
	fmt.Printf("Incidencias: %d\n", len(c.Incidences))
	fmt.Printf("=================================\n")
}

func (c *Car) addIncidence(in Incidence) {
	in.Id = len(c.Incidences) + 1
	c.Incidences = append(c.Incidences, in)
	fmt.Printf("Incidencia #%d añadida al vehículo %s\n", in.Id, c.RegistrationNumber)
}

func (c Car) listIncidences() {
	if len(c.Incidences) == 0 {
		fmt.Println("El vehículo no tiene incidencias registradas.")
		return
	}

	fmt.Printf("\n=== INCIDENCIAS DE %s ===\n", c.RegistrationNumber)
	for _, in := range c.Incidences {
		fmt.Printf("ID: %d | Tipo: %s | Prioridad: %s | Estado: %s\n", in.Id, in.Type, in.Priority, in.Status)
		if in.Description != "" {
			fmt.Printf("  Descripción: %s\n", in.Description)
		}
	}
	fmt.Println()
}
