package main

import "fmt"

// Workshop functions

func (w *Workshop) findClientByID(id int) *Client {
	for i := range w.Clients {
		if w.Clients[i].Id == id {
			return &w.Clients[i]
		}
	}
	return nil
}

func (w *Workshop) findCarByRegistration(reg string) *Car {
	for i := range w.Cars {
		if w.Cars[i].RegistrationNumber == reg {
			return &w.Cars[i]
		}
	}
	return nil
}

func (w *Workshop) findMechanicByID(id int) *Mechanic {
	for i := range w.Mecanics {
		if w.Mecanics[i].Id == id {
			return &w.Mecanics[i]
		}
	}
	return nil
}

func (w *Workshop) addClient(c Client) {
	c.Id = len(w.Clients) + 1
	w.Clients = append(w.Clients, c)
	fmt.Printf("Cliente %s añadido con ID %d.\n", c.Name, c.Id)
}

func (w *Workshop) addCar(car Car) bool {

	if !w.canAssignCar() {
		fmt.Printf("No hay espacio. Capacidad: %d/%d ocupada.\n",
			len(w.Cars), w.getTotalCapacity())
		return false
	}

	for _, c := range w.Cars {
		if c.RegistrationNumber == car.RegistrationNumber {
			fmt.Printf("Ya existe un vehículo con matrícula %s.\n", car.RegistrationNumber)
			return false
		}
	}

	w.Cars = append(w.Cars, car)
	fmt.Printf("Vehículo %s %s añadido al taller. Plazas: %d/%d\n",
		car.Brand, car.Model, len(w.Cars), w.getTotalCapacity())
	return true
}

func (w *Workshop) addMechanic(m Mechanic) {
	m.Id = len(w.Mecanics) + 1
	w.Mecanics = append(w.Mecanics, m)
	fmt.Printf("Mecánico %s añadido con ID %d.\n", m.Name, m.Id)
}

func (w *Workshop) removeClient(id int) bool {
	for i := range w.Clients {
		if w.Clients[i].Id == id {
			w.Clients = remove(w.Clients, i)
			fmt.Println("Cliente eliminado correctamente.")
			return true
		}
	}
	fmt.Println("Cliente no encontrado.")
	return false
}

func (w *Workshop) removeCar(reg string) bool {
	for i := range w.Cars {
		if w.Cars[i].RegistrationNumber == reg {
			w.Cars = remove(w.Cars, i)
			fmt.Printf("Vehículo con matrícula %s eliminado.\n", reg)
			return true
		}
	}
	vehicleNotFound(reg)
	return false
}

func (w *Workshop) removeMechanic(id int) bool {
	for i := range w.Mecanics {
		if w.Mecanics[i].Id == id {
			w.Mecanics = remove(w.Mecanics, i)
			fmt.Println("Mecánico eliminado correctamente.")
			return true
		}
	}
	fmt.Println("Mecánico no encontrado.")
	return false
}

func (w *Workshop) countActiveMechanics() int {
	count := 0
	for _, m := range w.Mecanics {
		if m.Activity == "Alta" {
			count++
		}
	}
	return count
}

func (w *Workshop) getTotalCapacity() int {
	return w.countActiveMechanics() * 2
}

func (w *Workshop) canAssignCar() bool {
	return len(w.Cars) < w.getTotalCapacity()
}

func (w Workshop) showStatus() {
	totalCapacity := w.getTotalCapacity()
	occupied := len(w.Cars)
	free := totalCapacity - occupied

	fmt.Printf("\n=== ESTADO DEL TALLER ===\n")
	fmt.Printf("Mecánicos activos: %d\n", w.countActiveMechanics())
	fmt.Printf("Capacidad total: %d plazas\n", totalCapacity)
	fmt.Printf("Plazas ocupadas: %d\n", occupied)
	fmt.Printf("Plazas libres: %d\n", free)
	fmt.Printf("=========================\n")
}

func (w *Workshop) listClientCars() {
	var id int
	fmt.Print("ID del cliente: ")
	fmt.Scanf("%d", &id)

	client := w.findClientByID(id)
	if client != nil {
		client.listCars()
	} else {
		fmt.Println("Cliente no encontrado.")
	}
}

func (w *Workshop) listCarIncidences() {
	var reg string
	fmt.Print("Matrícula del vehículo: ")
	fmt.Scan(&reg)

	car := w.findCarByRegistration(reg)
	if car != nil {
		car.listIncidences()
	} else {
		vehicleNotFound(reg)
	}
}

func (w *Workshop) listAvailableMechanics() {
	found := false
	fmt.Println("\n=== MECÁNICOS DISPONIBLES ===")
	for _, m := range w.Mecanics {
		if m.Activity == "Alta" && m.Status == "Disponible" {
			fmt.Printf("ID: %d | %s (%s) - %d años de experiencia\n", m.Id, m.Name, m.Speciality, m.YearsExperience)
			found = true
		}
	}
	if !found {
		fmt.Println("No hay mecánicos disponibles.")
	}
	fmt.Println()
}

func (w *Workshop) listMechanicIncidences() {
	var id int
	fmt.Print("ID del mecánico: ")
	fmt.Scanf("%d", &id)

	mechanic := w.findMechanicByID(id)
	if mechanic == nil {
		fmt.Println("Mecánico no encontrado.")
		return
	}

	found := false
	fmt.Printf("\n=== INCIDENCIAS DE %s ===\n", mechanic.Name)
	for _, car := range w.Cars {
		for _, inc := range car.Incidences {
			for _, m := range inc.Mecanics {
				if m.Id == id {
					fmt.Printf("Vehículo: %s | Incidencia #%d: %s (%s - %s)\n", car.RegistrationNumber, inc.Id, inc.Type, inc.Priority, inc.Status)
					found = true
				}
			}
		}
	}
	if !found {
		fmt.Println("No hay incidencias asignadas a este mecánico.")
	}
	fmt.Println()
}

func (w *Workshop) listClientsWithCarsInWorkshop() {
	// Crear mapa de matrículas en el taller
	carsInWorkshop := make(map[string]bool)
	for _, car := range w.Cars {
		if car.RegistrationNumber != "" {
			carsInWorkshop[car.RegistrationNumber] = true
		}
	}

	found := false
	fmt.Println("\n=== CLIENTES CON VEHÍCULOS EN EL TALLER ===")
	for _, client := range w.Clients {
		hasCarInWorkshop := false
		for _, car := range client.Cars {
			if carsInWorkshop[car.RegistrationNumber] {
				hasCarInWorkshop = true
				break
			}
		}
		if hasCarInWorkshop {
			fmt.Printf("ID: %d | %s (Tel: %s)\n", client.Id, client.Name, client.PhoneNumber)
			found = true
		}
	}
	if !found {
		fmt.Println("No hay clientes con vehículos en el taller.")
	}
	fmt.Println()
}

func (w *Workshop) listAllIncidences() {
	found := false
	fmt.Println("\n=== TODAS LAS INCIDENCIAS DEL TALLER ===")
	for _, car := range w.Cars {
		for _, inc := range car.Incidences {
			fmt.Printf("Vehículo: %s %s (%s) | Incidencia #%d: %s - %s - %s\n", car.Brand, car.Model, car.RegistrationNumber, inc.Id, inc.Type, inc.Priority, inc.Status)
			found = true
		}
	}
	if !found {
		fmt.Println("No hay incidencias registradas.")
	}
	fmt.Println()
}

// Eliminar incidencia de un vehículo
func (w *Workshop) removeIncidence() {
	var reg string
	var incId int

	fmt.Print("Matrícula del vehículo: ")
	fmt.Scan(&reg)

	car := w.findCarByRegistration(reg)
	if car == nil {
		vehicleNotFound(reg)
		return
	}

	if len(car.Incidences) == 0 {
		fmt.Println("El vehículo no tiene incidencias.")
		return
	}

	fmt.Println("Incidencias del vehículo:")
	for _, inc := range car.Incidences {
		fmt.Printf("ID: %d - %s (%s)\n", inc.Id, inc.Type, inc.Status)
	}

	fmt.Print("ID de la incidencia a eliminar: ")
	fmt.Scanf("%d", &incId)

	for i := range car.Incidences {
		if car.Incidences[i].Id == incId {
			car.Incidences = remove(car.Incidences, i)
			fmt.Println("Incidencia eliminada correctamente.")
			return
		}
	}
	fmt.Println("Incidencia no encontrada.")
}

// Modificar incidencia
func (w *Workshop) modifyIncidence() {
	var reg string
	var incId int

	fmt.Print("Matrícula del vehículo: ")
	fmt.Scan(&reg)

	car := w.findCarByRegistration(reg)
	if car == nil {
		vehicleNotFound(reg)
		return
	}

	if len(car.Incidences) == 0 {
		fmt.Println("El vehículo no tiene incidencias.")
		return
	}

	fmt.Println("Incidencias del vehículo:")
	for _, inc := range car.Incidences {
		fmt.Printf("ID: %d - %s (%s)\n", inc.Id, inc.Type, inc.Status)
	}

	fmt.Print("ID de la incidencia a modificar: ")
	fmt.Scanf("%d", &incId)

	for i := range car.Incidences {
		if car.Incidences[i].Id == incId {
			var option int
			fmt.Println("\n¿Qué desea modificar?")
			fmt.Println("1- Cambiar estado")
			fmt.Println("2- Cambiar prioridad")
			fmt.Println("3- Cambiar descripción")
			fmt.Print("Opción: ")
			fmt.Scanf("%d", &option)

			switch option {
			case 1:
				car.Incidences[i].changeStatus()
			case 2:
				fmt.Println("Nueva prioridad:")
				fmt.Println("1- Alta")
				fmt.Println("2- Media")
				fmt.Println("3- Baja")
				fmt.Scan(&option)
				switch option {
				case 1:
					car.Incidences[i].Priority = "Alta"
				case 2:
					car.Incidences[i].Priority = "Media"
				case 3:
					car.Incidences[i].Priority = "Baja"
				}
				fmt.Println("Prioridad actualizada.")
			case 3:
				fmt.Print("Nueva descripción: ")
				fmt.Scan(&car.Incidences[i].Description)
				fmt.Println("Descripción actualizada.")
			}
			return
		}
	}
	fmt.Println("Incidencia no encontrada.")
}

// Asignar mecánico a una incidencia existente
func (w *Workshop) assignMechanicToIncidence() {
	var reg string
	var incId, mechId int

	fmt.Print("Matrícula del vehículo: ")
	fmt.Scan(&reg)

	car := w.findCarByRegistration(reg)
	if car == nil {
		vehicleNotFound(reg)
		return
	}

	if len(car.Incidences) == 0 {
		fmt.Println("El vehículo no tiene incidencias.")
		return
	}

	fmt.Println("Incidencias del vehículo:")
	for _, inc := range car.Incidences {
		fmt.Printf("ID: %d - %s (%s)\n", inc.Id, inc.Type, inc.Status)
	}

	fmt.Print("ID de la incidencia: ")
	fmt.Scanf("%d", &incId)

	var incidenceFound *Incidence
	for i := range car.Incidences {
		if car.Incidences[i].Id == incId {
			incidenceFound = &car.Incidences[i]
			break
		}
	}

	if incidenceFound == nil {
		fmt.Println("Incidencia no encontrada.")
		return
	}

	// Mostrar mecánicos disponibles
	fmt.Println("\nMecánicos disponibles:")
	for _, m := range w.Mecanics {
		if m.Activity == "Alta" && m.Status == "Disponible" {
			fmt.Printf("ID: %d | %s (%s)\n", m.Id, m.Name, m.Speciality)
		}
	}

	fmt.Print("ID del mecánico: ")
	fmt.Scanf("%d", &mechId)

	mechanic := w.findMechanicByID(mechId)
	if mechanic == nil || mechanic.Activity != "Alta" {
		fmt.Println("Mecánico no disponible.")
		return
	}

	if incidenceFound.assignMechanic(*mechanic) {
		mechanic.Status = "Ocupado"
	}
}

// Cambiar estado de incidencia
func changeIncidenceStatus(w *Workshop) {
	var reg string
	var incId int

	fmt.Print("Matricula del vehiculo: ")
	fmt.Scan(&reg)

	car := w.findCarByRegistration(reg)
	if car == nil {
		vehicleNotFound(reg)
		return
	}

	if len(car.Incidences) == 0 {
		fmt.Println("El vehiculo no tiene incidencias.")
		return
	}

	fmt.Println("\nIncidencias del vehiculo:")
	for _, inc := range car.Incidences {
		fmt.Printf("ID: %d - %s (Estado: %s)\n", inc.Id, inc.Type, inc.Status)
	}

	fmt.Print("\nID de la incidencia: ")
	fmt.Scanf("%d", &incId)

	for i := range car.Incidences {
		if car.Incidences[i].Id == incId {
			car.Incidences[i].changeStatus()
			return
		}
	}
	fmt.Println("Incidencia no encontrada.")
}
