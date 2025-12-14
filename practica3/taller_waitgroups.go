package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type TallerChannels struct {
	NumPlazas    int
	NumMecanicos int

	canalPlazas    chan struct{}
	canalMecanicos chan struct{}
	canalLimpieza  chan struct{}
	canalRevision  chan struct{}

	tiempoInicio time.Time
}

func NuevoTallerChannels(numPlazas, numMecanicos int) *TallerChannels {
	return &TallerChannels{
		NumPlazas:      numPlazas,
		NumMecanicos:   numMecanicos,
		canalPlazas:    make(chan struct{}, numPlazas),
		canalMecanicos: make(chan struct{}, numMecanicos),
		canalLimpieza:  make(chan struct{}, 1),
		canalRevision:  make(chan struct{}, 1),
		tiempoInicio:   time.Now(),
	}
}

func (t *TallerChannels) TiempoTranscurrido() string {
	duracion := time.Since(t.tiempoInicio)
	return fmt.Sprintf("%.2f", duracion.Seconds())
}

func (t *TallerChannels) Log(cocheID int, tipo string, fase string, estado string) {
	fmt.Printf("Tiempo %s Coche %d Incidencia %s Fase %s Estado %s\n",
		t.TiempoTranscurrido(), cocheID, tipo, fase, estado)
}

func (t *TallerChannels) Fase1_Entrada(v Vehiculo) {
	t.Log(v.ID, v.Incidencia.Tipo, "Entrada", "Esperando")

	t.canalPlazas <- struct{}{}

	t.Log(v.ID, v.Incidencia.Tipo, "Entrada", "Iniciando")

	variacion := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(v.Incidencia.Tiempo + variacion)

	t.Log(v.ID, v.Incidencia.Tipo, "Entrada", "Completada")
}

func (t *TallerChannels) Fase2_Reparacion(v Vehiculo) {
	t.Log(v.ID, v.Incidencia.Tipo, "Reparacion", "Esperando")

	t.canalMecanicos <- struct{}{}

	t.Log(v.ID, v.Incidencia.Tipo, "Reparacion", "Iniciando")

	variacion := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(v.Incidencia.Tiempo + variacion)

	t.Log(v.ID, v.Incidencia.Tipo, "Reparacion", "Completada")

	<-t.canalMecanicos
}

func (t *TallerChannels) Fase3_Limpieza(v Vehiculo) {
	t.Log(v.ID, v.Incidencia.Tipo, "Limpieza", "Esperando")

	t.canalLimpieza <- struct{}{}

	t.Log(v.ID, v.Incidencia.Tipo, "Limpieza", "Iniciando")

	variacion := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(v.Incidencia.Tiempo + variacion)

	t.Log(v.ID, v.Incidencia.Tipo, "Limpieza", "Completada")

	<-t.canalLimpieza
}

func (t *TallerChannels) Fase4_Revision(v Vehiculo) {
	t.Log(v.ID, v.Incidencia.Tipo, "Revision", "Esperando")

	t.canalRevision <- struct{}{}

	t.Log(v.ID, v.Incidencia.Tipo, "Revision", "Iniciando")

	variacion := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(v.Incidencia.Tiempo + variacion)

	t.Log(v.ID, v.Incidencia.Tipo, "Revision", "Completada")

	<-t.canalRevision

	<-t.canalPlazas
}

func (t *TallerChannels) ProcesarVehiculo(v Vehiculo, wg *sync.WaitGroup) {
	defer wg.Done()

	t.Fase1_Entrada(v)
	t.Fase2_Reparacion(v)
	t.Fase3_Limpieza(v)
	t.Fase4_Revision(v)
}

func ejecutarSimulacionChannels(cantMecanica, cantElectrica, cantCarroceria, numPlazas, numMecanicos int) time.Duration {
	taller := NuevoTallerChannels(numPlazas, numMecanicos)
	vehiculos := generarVehiculos(cantMecanica, cantElectrica, cantCarroceria)

	var wg sync.WaitGroup

	fmt.Printf("\n========================================\n")
	fmt.Printf("SIMULACIÓN CON CHANNELS\n")
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
