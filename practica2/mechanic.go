package main

import "fmt"

// Mechanic functions

func (m *Mechanic) create() {
	fmt.Print("Introduzca el nombre del mecánico: ")
	fmt.Scan(&m.Name)

	var option int
	fmt.Println("Seleccione la especialidad:")
	fmt.Println("1- Motor")
	fmt.Println("2- Electricidad")
	fmt.Println("3- Carrocería")
	fmt.Print("Opción: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		m.Speciality = "Motor"
	case 2:
		m.Speciality = "Electricidad"
	case 3:
		m.Speciality = "Carrocería"
	default:
		m.Speciality = "Motor"
	}

	fmt.Print("Años de experiencia: ")
	fmt.Scan(&m.YearsExperience)

	m.Status = "Disponible"
	m.Activity = "Alta"
}

func (m Mechanic) display() {
	fmt.Printf("\n=== INFORMACIÓN DEL MECÁNICO ===\n")
	fmt.Printf("ID: %d\n", m.Id)
	fmt.Printf("Nombre: %s\n", m.Name)
	fmt.Printf("Especialidad: %s\n", m.Speciality)
	fmt.Printf("Años de experiencia: %d\n", m.YearsExperience)
	fmt.Printf("Estado: %s\n", m.Status)
	fmt.Printf("Actividad: %s\n", m.Activity)
	fmt.Println("=================================")
}

func (m *Mechanic) modify() {
	var option int

	fmt.Println("\n=== MODIFICAR MECÁNICO ===")
	fmt.Println("1- Modificar nombre")
	fmt.Println("2- Modificar especialidad")
	fmt.Println("3- Modificar años de experiencia")
	fmt.Print("Opción: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		fmt.Print("Nuevo nombre: ")
		fmt.Scan(&m.Name)
	case 2:
		var spec int
		fmt.Println("Especialidad:")
		fmt.Println("1- Motor")
		fmt.Println("2- Electricidad")
		fmt.Println("3- Carrocería")
		fmt.Scan(&spec)
		switch spec {
		case 1:
			m.Speciality = "Motor"
		case 2:
			m.Speciality = "Electricidad"
		case 3:
			m.Speciality = "Carrocería"
		}
	case 3:
		fmt.Print("Nuevos años de experiencia: ")
		fmt.Scan(&m.YearsExperience)
	default:
		fmt.Println("Opción inválida.")
	}
}

func (m *Mechanic) toggleActivity() {
	if m.Activity == "Baja" {
		m.Activity = "Alta"
		m.Status = "Disponible"
		fmt.Printf("Mecánico %s dado de ALTA.\n", m.Name)
	} else {
		m.Activity = "Baja"
		m.Status = "No disponible"
		fmt.Printf("Mecánico %s dado de BAJA.\n", m.Name)
	}
}
