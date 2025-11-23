package main

import "fmt"

// Incidence functions

// Crear incidencia (pide datos interactivos)
func (in *Incidence) create() {
	// Tipo
	var option int
	fmt.Println("Seleccione el tipo de incidencia:")
	fmt.Println("1- Mecánica")
	fmt.Println("2- Eléctrica")
	fmt.Println("3- Carrocería")
	fmt.Print("Opción: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		in.Type = "Mecánica"
	case 2:
		in.Type = "Eléctrica"
	case 3:
		in.Type = "Carrocería"
	}

	// Prioridad
	fmt.Println("Seleccione la prioridad:")
	fmt.Println("1- Alta")
	fmt.Println("2- Media")
	fmt.Println("3- Baja")
	fmt.Print("Opción: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		in.Priority = "Alta"
	case 2:
		in.Priority = "Media"
	case 3:
		in.Priority = "Baja"
	}

	// Estado
	fmt.Println("Seleccione el estado:")
	fmt.Println("1- Abierta")
	fmt.Println("2- En proceso")
	fmt.Println("3- Cerrada")
	fmt.Print("Opción: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		in.Status = "Abierta"
	case 2:
		in.Status = "En proceso"
	case 3:
		in.Status = "Cerrada"
	}

	// Descripción
	fmt.Print("Descripción de la incidencia: ")
	fmt.Scan(&in.Description)
}

// Mostrar información de la incidencia
func (in Incidence) display() {
	fmt.Printf("\n=== INCIDENCIA #%d ===\n", in.Id)
	fmt.Printf("Tipo: %s\n", in.Type)
	fmt.Printf("Prioridad: %s\n", in.Priority)
	fmt.Printf("Estado: %s\n", in.Status)
	fmt.Printf("Descripción: %s\n", in.Description)
	fmt.Printf("Mecánicos asignados: %d\n", len(in.Mecanics))
	for _, m := range in.Mecanics {
		fmt.Printf("  - %s (%s)\n", m.Name, m.Speciality)
	}
	fmt.Println("=======================")
}

// Cambiar estado de la incidencia
func (in *Incidence) changeStatus() {
	var option int
	fmt.Println("Nuevo estado:")
	fmt.Println("1- Abierta")
	fmt.Println("2- En proceso")
	fmt.Println("3- Cerrada")
	fmt.Print("Opción: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		in.Status = "Abierta"
	case 2:
		in.Status = "En proceso"
	case 3:
		in.Status = "Cerrada"
	}
	fmt.Printf("Estado actualizado a: %s\n", in.Status)
}

// Asignar mecánico a la incidencia
func (in *Incidence) assignMechanic(m Mechanic) bool {
	// Validar especialidad
	if (in.Type == "Mecánica" && m.Speciality != "Motor") ||
		(in.Type == "Eléctrica" && m.Speciality != "Electricidad") ||
		(in.Type == "Carrocería" && m.Speciality != "Carrocería") {
		fmt.Printf("El mecánico %s no está especializado en %s.\n", m.Name, in.Type)
		return false
	}

	in.Mecanics = append(in.Mecanics, m)
	fmt.Printf("Mecánico %s asignado a la incidencia.\n", m.Name)
	return true
}
