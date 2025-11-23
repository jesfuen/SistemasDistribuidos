package main

import (
	"fmt"
)

// Main function

func main() {
	var workshop Workshop

	workshop.Mecanics = append(workshop.Mecanics, Mechanic{Id: 1, Name: "Juan Perez", Speciality: "Motor", YearsExperience: 5, Status: "Disponible", Activity: "Alta"})
	workshop.Mecanics = append(workshop.Mecanics, Mechanic{Id: 2, Name: "Ana Gomez", Speciality: "Electricidad", YearsExperience: 3, Status: "Disponible", Activity: "Alta"})

	var option int
	for {
		fmt.Println("\n=== TALLER MECANICO - MENU PRINCIPAL ===")
		fmt.Println("1  - Gestion de Clientes")
		fmt.Println("2  - Gestion de Vehiculos")
		fmt.Println("3  - Gestion de Mecanicos")
		fmt.Println("4  - Gestion de Incidencias")
		fmt.Println("5  - Consultas y Listados")
		fmt.Println("6  - Estado del Taller")
		fmt.Println("0  - Salir")
		fmt.Print("Seleccione una opcion: ")
		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			menuClientes(&workshop)
		case 2:
			menuVehiculos(&workshop)
		case 3:
			menuMecanicos(&workshop)
		case 4:
			menuIncidencias(&workshop)
		case 5:
			menuConsultas(&workshop)
		case 6:
			workshop.showStatus()
		case 0:
			fmt.Println("\nGracias por usar el sistema. Hasta luego.")
			return
		default:
			fmt.Println("Opcion invalida. Intente nuevamente.")
		}
	}
}
