package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	TipoMecanica   = "Mecanica"
	TipoElectrica  = "Electrica"
	TipoCarroceria = "Carroceria"
)

const (
	PrioridadAlta  = 3
	PrioridadMedia = 2
	PrioridadBaja  = 1
)

type Incidencia struct {
	Tipo      string
	Prioridad int
	Tiempo    time.Duration
}

type Vehiculo struct {
	ID         int
	Incidencia Incidencia
}

type TallerRWMutex struct {
	NumPlazas       int
	NumMecanicos    int
	plazasOcupadas  int
	mecanicosLibres int

	muPlazas    sync.RWMutex
	muMecanicos sync.RWMutex

	muLimpieza sync.Mutex
	muRevision sync.Mutex

	tiempoInicio time.Time
}

func NuevaIncidencia(tipo string) Incidencia {
	var inc Incidencia
	inc.Tipo = tipo

	switch tipo {
	case TipoMecanica:
		inc.Prioridad = PrioridadAlta
		inc.Tiempo = 5 * time.Second
	case TipoElectrica:
		inc.Prioridad = PrioridadMedia
		inc.Tiempo = 3 * time.Second
	case TipoCarroceria:
		inc.Prioridad = PrioridadBaja
		inc.Tiempo = 1 * time.Second
	default:
		inc.Prioridad = PrioridadBaja
		inc.Tiempo = 1 * time.Second
	}

	return inc
}

func NuevoTallerRWMutex(numPlazas, numMecanicos int) *TallerRWMutex {
	return &TallerRWMutex{
		NumPlazas:       numPlazas,
		NumMecanicos:    numMecanicos,
		plazasOcupadas:  0,
		mecanicosLibres: numMecanicos,
		tiempoInicio:    time.Now(),
	}
}

func (t *TallerRWMutex) TiempoTranscurrido() string {
	duracion := time.Since(t.tiempoInicio)
	return fmt.Sprintf("%.2f", duracion.Seconds())
}

func (t *TallerRWMutex) Log(cocheID int, tipo string, fase string, estado string) {
	fmt.Printf("Tiempo %s Coche %d Incidencia %s Fase %s Estado %s\n",
		t.TiempoTranscurrido(), cocheID, tipo, fase, estado)
}

func (t *TallerRWMutex) Fase1_Entrada(v Vehiculo) {
	t.Log(v.ID, v.Incidencia.Tipo, "Entrada", "Esperando")

	for {
		t.muPlazas.Lock()
		if t.plazasOcupadas < t.NumPlazas {
			t.plazasOcupadas++
			t.muPlazas.Unlock()
			break
		}
		t.muPlazas.Unlock()
		time.Sleep(100 * time.Millisecond)
	}

	t.Log(v.ID, v.Incidencia.Tipo, "Entrada", "Iniciando")

	variacion := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(v.Incidencia.Tiempo + variacion)

	t.Log(v.ID, v.Incidencia.Tipo, "Entrada", "Completada")
}

func (t *TallerRWMutex) Fase2_Reparacion(v Vehiculo) {
	t.Log(v.ID, v.Incidencia.Tipo, "Reparacion", "Esperando")

	for {
		t.muMecanicos.Lock()
		if t.mecanicosLibres > 0 {
			t.mecanicosLibres--
			t.muMecanicos.Unlock()
			break
		}
		t.muMecanicos.Unlock()
		time.Sleep(100 * time.Millisecond)
	}

	t.Log(v.ID, v.Incidencia.Tipo, "Reparacion", "Iniciando")

	variacion := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(v.Incidencia.Tiempo + variacion)

	t.Log(v.ID, v.Incidencia.Tipo, "Reparacion", "Completada")

	t.muMecanicos.Lock()
	t.mecanicosLibres++
	t.muMecanicos.Unlock()
}

func (t *TallerRWMutex) Fase3_Limpieza(v Vehiculo) {
	t.Log(v.ID, v.Incidencia.Tipo, "Limpieza", "Esperando")

	t.muLimpieza.Lock()

	t.Log(v.ID, v.Incidencia.Tipo, "Limpieza", "Iniciando")

	variacion := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(v.Incidencia.Tiempo + variacion)

	t.Log(v.ID, v.Incidencia.Tipo, "Limpieza", "Completada")

	t.muLimpieza.Unlock()
}

func (t *TallerRWMutex) Fase4_Revision(v Vehiculo) {
	t.Log(v.ID, v.Incidencia.Tipo, "Revision", "Esperando")

	t.muRevision.Lock()

	t.Log(v.ID, v.Incidencia.Tipo, "Revision", "Iniciando")

	variacion := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(v.Incidencia.Tiempo + variacion)

	t.Log(v.ID, v.Incidencia.Tipo, "Revision", "Completada")

	t.muRevision.Unlock()

	t.muPlazas.Lock()
	t.plazasOcupadas--
	t.muPlazas.Unlock()
}

func (t *TallerRWMutex) ProcesarVehiculo(v Vehiculo, wg *sync.WaitGroup) {
	defer wg.Done()

	t.Fase1_Entrada(v)
	t.Fase2_Reparacion(v)
	t.Fase3_Limpieza(v)
	t.Fase4_Revision(v)
}

func generarVehiculos(cantMecanica, cantElectrica, cantCarroceria int) []Vehiculo {
	vehiculos := make([]Vehiculo, 0)
	id := 1

	// Categoría A: Mecánica (Prioridad Alta)
	for i := 0; i < cantMecanica; i++ {
		vehiculos = append(vehiculos, Vehiculo{
			ID:         id,
			Incidencia: NuevaIncidencia(TipoMecanica),
		})
		id++
	}

	// Categoría B: Eléctrica (Prioridad Media)
	for i := 0; i < cantElectrica; i++ {
		vehiculos = append(vehiculos, Vehiculo{
			ID:         id,
			Incidencia: NuevaIncidencia(TipoElectrica),
		})
		id++
	}

	// Categoría C: Carrocería (Prioridad Baja)
	for i := 0; i < cantCarroceria; i++ {
		vehiculos = append(vehiculos, Vehiculo{
			ID:         id,
			Incidencia: NuevaIncidencia(TipoCarroceria),
		})
		id++
	}

	rand.Shuffle(len(vehiculos), func(i, j int) {
		vehiculos[i], vehiculos[j] = vehiculos[j], vehiculos[i]
	})

	return vehiculos
}

func ejecutarSimulacionRWMutex(cantMecanica, cantElectrica, cantCarroceria, numPlazas, numMecanicos int) time.Duration {
	taller := NuevoTallerRWMutex(numPlazas, numMecanicos)
	vehiculos := generarVehiculos(cantMecanica, cantElectrica, cantCarroceria)

	var wg sync.WaitGroup

	fmt.Printf("\n========================================\n")
	fmt.Printf("SIMULACIÓN CON RWMUTEX\n")
	fmt.Printf("========================================\n")
	fmt.Printf("Categoría A (Mecánica):   %d vehículos\n", cantMecanica)
	fmt.Printf("Categoría B (Eléctrica):  %d vehículos\n", cantElectrica)
	fmt.Printf("Categoría C (Carrocería): %d vehículos\n", cantCarroceria)
	fmt.Printf("Plazas: %d | Mecánicos: %d\n", numPlazas, numMecanicos)
	fmt.Printf("========================================\n\n")

	inicio := time.Now()

	for _, vehiculo := range vehiculos {
		wg.Add(1)
		go taller.ProcesarVehiculo(vehiculo, &wg)
		time.Sleep(50 * time.Millisecond)
	}

	wg.Wait()

	duracion := time.Since(inicio)

	fmt.Printf("\n========================================\n")
	fmt.Printf("SIMULACIÓN COMPLETADA\n")
	fmt.Printf("Tiempo total: %v\n", duracion)
	fmt.Printf("========================================\n\n")

	return duracion
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Test 1: 10-10-10
	ejecutarSimulacionRWMutex(10, 10, 10, 5, 3)
	ejecutarSimulacionChannels(10, 10, 10, 5, 3)

	// Test 2: 20-5-5
	ejecutarSimulacionRWMutex(20, 5, 5, 5, 3)
	ejecutarSimulacionChannels(20, 5, 5, 5, 3)

	// Test 3: 5-5-20
	ejecutarSimulacionRWMutex(5, 5, 20, 5, 3)
	ejecutarSimulacionChannels(5, 5, 20, 5, 3)
}
