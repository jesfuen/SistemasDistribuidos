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

func vehicleNotFound(r string) {
	fmt.Printf("El vehiculo con matricula %s no esta registrado.\n", r)
}

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

// Main function

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

// Menu de Clientes
func menuClientes(w *Workshop) {
	var option int
	for {
		fmt.Println("\n=== GESTION DE CLIENTES ===")
		fmt.Println("1 - Crear cliente")
		fmt.Println("2 - Visualizar cliente")
		fmt.Println("3 - Modificar cliente")
		fmt.Println("4 - Eliminar cliente")
		fmt.Println("0 - Volver al menu principal")
		fmt.Print("Seleccione una opcion: ")
		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			var newClient Client
			newClient.create(w)
			w.addClient(newClient)

		case 2:
			var id int
			fmt.Print("ID del cliente: ")
			fmt.Scanf("%d", &id)
			client := w.findClientByID(id)
			if client != nil {
				client.display()
			} else {
				fmt.Println("Cliente no encontrado.")
			}

		case 3:
			var id int
			fmt.Print("ID del cliente: ")
			fmt.Scanf("%d", &id)
			client := w.findClientByID(id)
			if client != nil {
				client.modify()
			} else {
				fmt.Println("Cliente no encontrado.")
			}

		case 4:
			var id int
			fmt.Print("ID del cliente a eliminar: ")
			fmt.Scanf("%d", &id)
			w.removeClient(id)

		case 0:
			return

		default:
			fmt.Println("Opcion invalida.")
		}
	}
}

// Menu de Vehiculos
func menuVehiculos(w *Workshop) {
	var option int
	for {
		fmt.Println("\n=== GESTION DE VEHICULOS ===")
		fmt.Println("1 - Crear vehiculo")
		fmt.Println("2 - Visualizar vehiculo")
		fmt.Println("3 - Modificar vehiculo")
		fmt.Println("4 - Eliminar vehiculo")
		fmt.Println("5 - Asignar vehiculo al taller")
		fmt.Println("0 - Volver al menu principal")
		fmt.Print("Seleccione una opcion: ")
		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			var newCar Car
			newCar.create()
			w.addCar(newCar)

		case 2:
			var reg string
			fmt.Print("Matricula del vehiculo: ")
			fmt.Scan(&reg)
			car := w.findCarByRegistration(reg)
			if car != nil {
				car.display()
			} else {
				vehicleNotFound(reg)
			}

		case 3:
			var reg string
			fmt.Print("Matricula del vehiculo: ")
			fmt.Scan(&reg)
			car := w.findCarByRegistration(reg)
			if car != nil {
				car.modify()
			} else {
				vehicleNotFound(reg)
			}

		case 4:
			var reg string
			fmt.Print("Matricula del vehiculo a eliminar: ")
			fmt.Scan(&reg)
			w.removeCar(reg)

		case 5:
			var newCar Car
			newCar.create()
			w.addCar(newCar)

		case 0:
			return

		default:
			fmt.Println("Opcion invalida.")
		}
	}
}

// Menu de Mecanicos
func menuMecanicos(w *Workshop) {
	var option int
	for {
		fmt.Println("\n=== GESTION DE MECANICOS ===")
		fmt.Println("1 - Crear mecanico")
		fmt.Println("2 - Visualizar mecanico")
		fmt.Println("3 - Modificar mecanico")
		fmt.Println("4 - Eliminar mecanico")
		fmt.Println("5 - Dar de alta/baja a mecanico")
		fmt.Println("0 - Volver al menu principal")
		fmt.Print("Seleccione una opcion: ")
		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			var newMechanic Mechanic
			newMechanic.create()
			w.addMechanic(newMechanic)

		case 2:
			var id int
			fmt.Print("ID del mecanico: ")
			fmt.Scanf("%d", &id)
			mechanic := w.findMechanicByID(id)
			if mechanic != nil {
				mechanic.display()
			} else {
				fmt.Println("Mecanico no encontrado.")
			}

		case 3:
			var id int
			fmt.Print("ID del mecanico: ")
			fmt.Scanf("%d", &id)
			mechanic := w.findMechanicByID(id)
			if mechanic != nil {
				mechanic.modify()
			} else {
				fmt.Println("Mecanico no encontrado.")
			}

		case 4:
			var id int
			fmt.Print("ID del mecanico a eliminar: ")
			fmt.Scanf("%d", &id)
			w.removeMechanic(id)

		case 5:
			var id int
			fmt.Print("ID del mecanico: ")
			fmt.Scanf("%d", &id)
			mechanic := w.findMechanicByID(id)
			if mechanic != nil {
				mechanic.toggleActivity()
			} else {
				fmt.Println("Mecanico no encontrado.")
			}

		case 0:
			return

		default:
			fmt.Println("Opcion invalida.")
		}
	}
}

// Menu de Incidencias
func menuIncidencias(w *Workshop) {
	var option int
	for {
		fmt.Println("\n=== GESTION DE INCIDENCIAS ===")
		fmt.Println("1 - Crear incidencia")
		fmt.Println("2 - Visualizar incidencia")
		fmt.Println("3 - Modificar incidencia")
		fmt.Println("4 - Eliminar incidencia")
		fmt.Println("5 - Cambiar estado de incidencia")
		fmt.Println("6 - Asignar mecanico a incidencia")
		fmt.Println("0 - Volver al menu principal")
		fmt.Print("Seleccione una opcion: ")
		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			var reg string
			fmt.Print("Matricula del vehiculo: ")
			fmt.Scan(&reg)
			car := w.findCarByRegistration(reg)
			if car != nil {
				var newIncidence Incidence
				newIncidence.create()
				car.addIncidence(newIncidence)
			} else {
				vehicleNotFound(reg)
			}

		case 2:
			visualizeIncidence(w)

		case 3:
			w.modifyIncidence()

		case 4:
			w.removeIncidence()

		case 5:
			changeIncidenceStatus(w)

		case 6:
			w.assignMechanicToIncidence()

		case 0:
			return

		default:
			fmt.Println("Opcion invalida.")
		}
	}
}

// Menu de Consultas
func menuConsultas(w *Workshop) {
	var option int
	for {
		fmt.Println("\n=== CONSULTAS Y LISTADOS ===")
		fmt.Println("1 - Listar incidencias de un vehiculo")
		fmt.Println("2 - Listar vehiculos de un cliente")
		fmt.Println("3 - Listar mecanicos disponibles")
		fmt.Println("4 - Listar incidencias de un mecanico")
		fmt.Println("5 - Listar clientes con vehiculos en taller")
		fmt.Println("6 - Listar todas las incidencias del taller")
		fmt.Println("0 - Volver al menu principal")
		fmt.Print("Seleccione una opcion: ")
		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			w.listCarIncidences()

		case 2:
			w.listClientCars()

		case 3:
			w.listAvailableMechanics()

		case 4:
			w.listMechanicIncidences()

		case 5:
			w.listClientsWithCarsInWorkshop()

		case 6:
			w.listAllIncidences()

		case 0:
			return

		default:
			fmt.Println("Opcion invalida.")
		}
	}
}

// Visualizar incidencia
func visualizeIncidence(w *Workshop) {
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
		fmt.Printf("ID: %d - %s (%s)\n", inc.Id, inc.Type, inc.Status)
	}

	fmt.Print("\nID de la incidencia a visualizar: ")
	fmt.Scanf("%d", &incId)

	for _, inc := range car.Incidences {
		if inc.Id == incId {
			inc.display()
			return
		}
	}
	fmt.Println("Incidencia no encontrada.")
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
