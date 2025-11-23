package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Sistema concurrente usando SOLO goroutines y channels (sin mutex ni map)
type ConcurrentWorkshop struct {
	workshop      *Workshop
	carQueue      chan *CarWork
	motorPool     chan *Mechanic
	electricPool  chan *Mechanic
	bodyPool      chan *Mechanic
	completedWork chan WorkResult
	quit          chan bool
	statsRequest  chan chan Stats
}

// Trabajo de un coche
type CarWork struct {
	Car           *Car
	Incidence     *Incidence
	ArrivalTime   time.Time
	TotalWorkTime float64
}

// Resultado del trabajo
type WorkResult struct {
	Car              *Car
	Incidence        *Incidence
	Mechanic         *Mechanic
	WorkDuration     time.Duration
	CompletionTime   time.Time
	TotalCarWorkTime float64
}

// Estadísticas del sistema
type Stats struct {
	TotalMechanics     int
	MotorMechanics     int
	ElectricMechanics  int
	BodyMechanics      int
	BusyMechanics      int
	AvailableMechanics int
}

// Inicializar el sistema concurrente
func NewConcurrentWorkshop(w *Workshop) *ConcurrentWorkshop {
	cw := &ConcurrentWorkshop{
		workshop:      w,
		carQueue:      make(chan *CarWork, 100), // Cola prácticamente ilimitada
		motorPool:     make(chan *Mechanic, 50),
		electricPool:  make(chan *Mechanic, 50),
		bodyPool:      make(chan *Mechanic, 50),
		completedWork: make(chan WorkResult, 50),
		quit:          make(chan bool),
		statsRequest:  make(chan chan Stats),
	}

	// Inicializar pools con mecánicos disponibles
	for i := range w.Mecanics {
		if w.Mecanics[i].Status == "Disponible" && w.Mecanics[i].Activity == "Alta" {
			switch w.Mecanics[i].Speciality {
			case "Motor":
				cw.motorPool <- &w.Mecanics[i]
			case "Electricidad":
				cw.electricPool <- &w.Mecanics[i]
			case "Carrocería":
				cw.bodyPool <- &w.Mecanics[i]
			}
		}
	}

	return cw
}

// Iniciar el dispatcher que asigna trabajos
func (cw *ConcurrentWorkshop) StartDispatcher() {
	go func() {
		fmt.Println("Dispatcher iniciado - Esperando coches...")
		for {
			select {
			case work := <-cw.carQueue:
				// Lanzar goroutine para procesar el trabajo
				go cw.processCarWork(work)
			case <-cw.quit:
				fmt.Println("Dispatcher detenido")
				return
			}
		}
	}()
}

// Procesar trabajo de un coche
func (cw *ConcurrentWorkshop) processCarWork(work *CarWork) {
	// Determinar pool según tipo de incidencia
	var pool chan *Mechanic

	switch work.Incidence.Type {
	case "Mecánica":
		pool = cw.motorPool
	case "Eléctrica":
		pool = cw.electricPool
	case "Carrocería":
		pool = cw.bodyPool
	default:
		pool = cw.motorPool
	}

	// Esperar por mecánico disponible (BLOQUEANTE)
	mechanic := <-pool

	// Marcar mecánico como ocupado
	mechanic.Status = "Ocupado"

	// Determinar tiempo de trabajo según especialidad
	workTime := cw.getWorkTimeForType(work.Incidence.Type)

	waitTime := time.Since(work.ArrivalTime)
	fmt.Printf("Coche %s asignado a %s (%s) - Esperó: %.2fs - Trabajo: %.1fs\n",
		work.Car.RegistrationNumber, mechanic.Name, mechanic.Speciality,
		waitTime.Seconds(), workTime.Seconds())

	// Asignar mecánico a la incidencia
	work.Incidence.assignMechanic(*mechanic)
	work.Incidence.Status = "En proceso"

	// Simular trabajo
	time.Sleep(workTime)

	// Actualizar tiempo acumulado
	work.TotalWorkTime += workTime.Seconds()

	// Verificar si necesita más tiempo (> 15 segundos)
	if work.TotalWorkTime > 15.0 {
		fmt.Printf("Coche %s acumula %.1fs - NECESITA MECÁNICO ADICIONAL\n",
			work.Car.RegistrationNumber, work.TotalWorkTime)

		// Intentar asignar otro mecánico de cualquier especialidad
		go cw.tryAssignAdditionalMechanic(work)
	}

	// Completar trabajo
	result := WorkResult{
		Car:              work.Car,
		Incidence:        work.Incidence,
		Mechanic:         mechanic,
		WorkDuration:     workTime,
		CompletionTime:   time.Now(),
		TotalCarWorkTime: work.TotalWorkTime,
	}

	// Si el trabajo está completo
	if work.TotalWorkTime <= 15.0 || work.Incidence.Status == "Cerrada" {
		work.Incidence.Status = "Cerrada"
	}

	// Liberar mecánico
	mechanic.Status = "Disponible"
	pool <- mechanic

	// Enviar resultado
	cw.completedWork <- result

	fmt.Printf("%s terminó trabajo en %s (%.1fs) - Total coche: %.1fs\n",
		mechanic.Name, work.Car.RegistrationNumber,
		workTime.Seconds(), work.TotalWorkTime)
}

// Intentar asignar mecánico adicional
func (cw *ConcurrentWorkshop) tryAssignAdditionalMechanic(work *CarWork) {
	// Intentar obtener mecánico de cualquier pool
	var mechanic *Mechanic
	var pool chan *Mechanic
	found := false

	// Intentar motor
	select {
	case mechanic = <-cw.motorPool:
		pool = cw.motorPool
		found = true
	default:
	}

	// Intentar eléctrico si no encontró
	if !found {
		select {
		case mechanic = <-cw.electricPool:
			pool = cw.electricPool
			found = true
		default:
		}
	}

	// Intentar carrocería si no encontró
	if !found {
		select {
		case mechanic = <-cw.bodyPool:
			pool = cw.bodyPool
			found = true
		default:
		}
	}

	if found {
		fmt.Printf("Mecánico adicional %s (%s) asignado a %s\n",
			mechanic.Name, mechanic.Speciality, work.Car.RegistrationNumber)

		// Trabajar en paralelo
		mechanic.Status = "Ocupado"
		work.Incidence.assignMechanic(*mechanic)

		// Trabajo adicional más corto
		workTime := cw.getWorkTimeForType(work.Incidence.Type) / 2
		time.Sleep(workTime)

		work.TotalWorkTime += workTime.Seconds()
		fmt.Printf("%s terminó trabajo adicional en %s (%.1fs)\n",
			mechanic.Name, work.Car.RegistrationNumber, workTime.Seconds())

		// Liberar
		mechanic.Status = "Disponible"
		pool <- mechanic

		work.Incidence.Status = "Cerrada"
	} else {
		// No hay mecánicos, contratar uno nuevo
		fmt.Printf("📞 CONTRATANDO nuevo mecánico de %s\n",
			work.Incidence.Type)

		newMechanic := cw.hireMechanic(work.Incidence.Type)

		// Trabajar con el nuevo mecánico
		newMechanic.Status = "Ocupado"
		work.Incidence.assignMechanic(*newMechanic)

		workTime := cw.getWorkTimeForType(work.Incidence.Type) / 2
		time.Sleep(workTime)

		work.TotalWorkTime += workTime.Seconds()
		fmt.Printf("Nuevo mecánico %s terminó trabajo en %s (%.1fs)\n",
			newMechanic.Name, work.Car.RegistrationNumber, workTime.Seconds())

		// Añadir el nuevo mecánico al pool
		newMechanic.Status = "Disponible"
		switch newMechanic.Speciality {
		case "Motor":
			cw.motorPool <- newMechanic
		case "Electricidad":
			cw.electricPool <- newMechanic
		case "Carrocería":
			cw.bodyPool <- newMechanic
		}

		work.Incidence.Status = "Cerrada"
	}
}

// Contratar nuevo mecánico (sincronizado con channel)
func (cw *ConcurrentWorkshop) hireMechanic(incidenceType string) *Mechanic {
	// Determinar especialidad
	var speciality string
	switch incidenceType {
	case "Mecánica":
		speciality = "Motor"
	case "Eléctrica":
		speciality = "Electricidad"
	case "Carrocería":
		speciality = "Carrocería"
	default:
		speciality = "Motor"
	}

	newId := len(cw.workshop.Mecanics) + 1
	newMechanic := Mechanic{
		Id:              newId,
		Name:            fmt.Sprintf("Mecánico-%d", newId),
		Speciality:      speciality,
		YearsExperience: 1,
		Status:          "Disponible",
		Activity:        "Alta",
	}

	cw.workshop.Mecanics = append(cw.workshop.Mecanics, newMechanic)
	mechPtr := &cw.workshop.Mecanics[len(cw.workshop.Mecanics)-1]

	return mechPtr
}

// Obtener tiempo de trabajo según tipo
func (cw *ConcurrentWorkshop) getWorkTimeForType(incidenceType string) time.Duration {
	// Tiempo medio +/- variación aleatoria
	rand.Seed(time.Now().UnixNano())
	variation := float64(rand.Intn(2000)-1000) / 1000.0 // +/- 1 segundo

	switch incidenceType {
	case "Mecánica":
		return time.Duration((5.0 + variation) * float64(time.Second))
	case "Eléctrica":
		return time.Duration((7.0 + variation) * float64(time.Second))
	case "Carrocería":
		return time.Duration((11.0 + variation) * float64(time.Second))
	default:
		return 5 * time.Second
	}
}

// Añadir coche a la cola
func (cw *ConcurrentWorkshop) AddCarToQueue(car *Car, incidence *Incidence) {
	work := &CarWork{
		Car:           car,
		Incidence:     incidence,
		ArrivalTime:   time.Now(),
		TotalWorkTime: 0,
	}

	queueLen := len(cw.carQueue)

	// Determinar pool para ver disponibilidad
	var pool chan *Mechanic
	switch incidence.Type {
	case "Mecánica":
		pool = cw.motorPool
	case "Eléctrica":
		pool = cw.electricPool
	case "Carrocería":
		pool = cw.bodyPool
	default:
		pool = cw.motorPool
	}

	availableMechs := len(pool)

	if availableMechs > 0 {
		fmt.Printf("Coche %s añadido - Mecánicos disponibles: %d\n",
			car.RegistrationNumber, availableMechs)
	} else {
		fmt.Printf("Coche %s en cola de espera (Posición: %d)\n",
			car.RegistrationNumber, queueLen+1)
	}

	cw.carQueue <- work
}

// Monitor de trabajos completados
func (cw *ConcurrentWorkshop) StartMonitor() {
	go func() {
		for {
			select {
			case <-cw.completedWork:
				// Ya se imprime en processCarWork
			case <-cw.quit:
				return
			}
		}
	}()
}

// Goroutine para manejar consultas de estadísticas
func (cw *ConcurrentWorkshop) StartStatsHandler() {
	go func() {
		for {
			select {
			case responseChan := <-cw.statsRequest:
				// Calcular estadísticas
				stats := Stats{
					TotalMechanics:     len(cw.workshop.Mecanics),
					MotorMechanics:     len(cw.motorPool),
					ElectricMechanics:  len(cw.electricPool),
					BodyMechanics:      len(cw.bodyPool),
					AvailableMechanics: len(cw.motorPool) + len(cw.electricPool) + len(cw.bodyPool),
				}
				stats.BusyMechanics = stats.TotalMechanics - stats.AvailableMechanics

				// Enviar respuesta
				responseChan <- stats
			case <-cw.quit:
				return
			}
		}
	}()
}

// Obtener estadísticas (usando channel)
func (cw *ConcurrentWorkshop) GetStats() (int, int, int) {
	responseChan := make(chan Stats)
	cw.statsRequest <- responseChan
	stats := <-responseChan

	return stats.TotalMechanics, stats.BusyMechanics, stats.AvailableMechanics
}

// Detener el sistema
func (cw *ConcurrentWorkshop) Stop() {
	close(cw.quit)
}
