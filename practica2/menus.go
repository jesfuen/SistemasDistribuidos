package main

import "fmt"

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
