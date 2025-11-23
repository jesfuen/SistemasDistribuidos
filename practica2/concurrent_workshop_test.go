package main

import (
	"fmt"
	"testing"
	"time"
)

// Test 1: Comparativa duplicando cantidad de coches
// Caso A: 10 coches con incidencia mecánica
// Caso B: 20 coches con incidencia mecánica
func TestDoubleCarLoad(t *testing.T) {
	fmt.Println()
	fmt.Println("======================================================================")
	fmt.Println("TEST 1: DUPLICAR CANTIDAD DE COCHES")
	fmt.Println("======================================================================")
	fmt.Println()

	// Caso A: 10 coches
	t.Run("10_coches_mecanica", func(t *testing.T) {
		workshop := createTestWorkshop(3, 1, 1, 1) // 3 mecánicos: 1 de cada
		cw := NewConcurrentWorkshop(&workshop)
		cw.StartDispatcher()
		cw.StartMonitor()
		cw.StartStatsHandler()
		defer cw.Stop()

		startTime := time.Now()

		// Añadir 10 coches con incidencia mecánica
		for i := 1; i <= 10; i++ {
			car := Car{
				RegistrationNumber: fmt.Sprintf("TEST-%03d", i),
				Brand:              "TestBrand",
				Model:              "TestModel",
			}
			workshop.Cars = append(workshop.Cars, car)
			carPtr := &workshop.Cars[len(workshop.Cars)-1]

			incidence := Incidence{
				Id:          1,
				Type:        "Mecánica",
				Priority:    "Media",
				Description: "Test mecánica",
				Status:      "Abierta",
			}
			carPtr.Incidences = append(carPtr.Incidences, incidence)
			incPtr := &carPtr.Incidences[0]

			cw.AddCarToQueue(carPtr, incPtr)
			time.Sleep(100 * time.Millisecond) // Simular llegada escalonada
		}

		// Esperar a que terminen todos
		time.Sleep(60 * time.Second)

		duration := time.Since(startTime)
		total, busy, available := cw.GetStats()

		t.Logf("\n RESULTADOS - 10 COCHES:")
		t.Logf("   Tiempo total: %.2f segundos", duration.Seconds())
		t.Logf("   Mecánicos totales: %d", total)
		t.Logf("   Mecánicos contratados: %d", total-3)
		t.Logf("   Mecánicos ocupados al final: %d", busy)
		t.Logf("   Mecánicos disponibles: %d\n", available)
	})

	// Caso B: 20 coches
	t.Run("20_coches_mecanica", func(t *testing.T) {
		workshop := createTestWorkshop(3, 1, 1, 1)
		cw := NewConcurrentWorkshop(&workshop)
		cw.StartDispatcher()
		cw.StartMonitor()
		cw.StartStatsHandler()
		defer cw.Stop()

		startTime := time.Now()

		// Añadir 20 coches con incidencia mecánica
		for i := 1; i <= 20; i++ {
			car := Car{
				RegistrationNumber: fmt.Sprintf("TEST-%03d", i),
				Brand:              "TestBrand",
				Model:              "TestModel",
			}
			workshop.Cars = append(workshop.Cars, car)
			carPtr := &workshop.Cars[len(workshop.Cars)-1]

			incidence := Incidence{
				Id:          1,
				Type:        "Mecánica",
				Priority:    "Media",
				Description: "Test mecánica",
				Status:      "Abierta",
			}
			carPtr.Incidences = append(carPtr.Incidences, incidence)
			incPtr := &carPtr.Incidences[0]

			cw.AddCarToQueue(carPtr, incPtr)
			time.Sleep(100 * time.Millisecond)
		}

		// Esperar a que terminen todos
		time.Sleep(120 * time.Second)

		duration := time.Since(startTime)
		total, busy, available := cw.GetStats()

		t.Logf("\n RESULTADOS - 20 COCHES:")
		t.Logf("   Tiempo total: %.2f segundos", duration.Seconds())
		t.Logf("   Mecánicos totales: %d", total)
		t.Logf("   Mecánicos contratados: %d", total-3)
		t.Logf("   Mecánicos ocupados al final: %d", busy)
		t.Logf("   Mecánicos disponibles: %d\n", available)
	})
}

// Test 2: Comparativa duplicando plantilla de mecánicos
// Caso A: 3 mecánicos (1 de cada especialidad)
// Caso B: 6 mecánicos (2 de cada especialidad)
func TestDoubleMechanicStaff(t *testing.T) {
	fmt.Println()
	fmt.Println("======================================================================")
	fmt.Println("TEST 2: DUPLICAR PLANTILLA DE MECÁNICOS")
	fmt.Println("======================================================================")
	fmt.Println()

	numCars := 15 // Misma carga de coches

	// Caso A: 3 mecánicos
	t.Run("3_mecanicos", func(t *testing.T) {
		workshop := createTestWorkshop(3, 1, 1, 1)
		cw := NewConcurrentWorkshop(&workshop)
		cw.StartDispatcher()
		cw.StartMonitor()
		cw.StartStatsHandler()
		defer cw.Stop()

		startTime := time.Now()

		// Añadir coches mezclados (5 de cada tipo)
		for i := 1; i <= numCars; i++ {
			car := Car{
				RegistrationNumber: fmt.Sprintf("CAR-%03d", i),
				Brand:              "TestBrand",
				Model:              "TestModel",
			}
			workshop.Cars = append(workshop.Cars, car)
			carPtr := &workshop.Cars[len(workshop.Cars)-1]

			// Alternar tipos de incidencia
			incType := []string{"Mecánica", "Eléctrica", "Carrocería"}[i%3]
			incidence := Incidence{
				Id:          1,
				Type:        incType,
				Priority:    "Media",
				Description: fmt.Sprintf("Test %s", incType),
				Status:      "Abierta",
			}
			carPtr.Incidences = append(carPtr.Incidences, incidence)
			incPtr := &carPtr.Incidences[0]

			cw.AddCarToQueue(carPtr, incPtr)
			time.Sleep(100 * time.Millisecond)
		}

		time.Sleep(120 * time.Second)

		duration := time.Since(startTime)
		total, busy, available := cw.GetStats()

		t.Logf("\n RESULTADOS - 3 MECÁNICOS:")
		t.Logf("   Tiempo total: %.2f segundos", duration.Seconds())
		t.Logf("   Mecánicos totales: %d", total)
		t.Logf("   Mecánicos contratados: %d", total-3)
		t.Logf("   Mecánicos ocupados al final: %d", busy)
		t.Logf("   Mecánicos disponibles: %d\n", available)
	})

	// Caso B: 6 mecánicos
	t.Run("6_mecanicos", func(t *testing.T) {
		workshop := createTestWorkshop(6, 2, 2, 2)
		cw := NewConcurrentWorkshop(&workshop)
		cw.StartDispatcher()
		cw.StartMonitor()
		cw.StartStatsHandler()
		defer cw.Stop()

		startTime := time.Now()

		// Añadir misma cantidad de coches
		for i := 1; i <= numCars; i++ {
			car := Car{
				RegistrationNumber: fmt.Sprintf("CAR-%03d", i),
				Brand:              "TestBrand",
				Model:              "TestModel",
			}
			workshop.Cars = append(workshop.Cars, car)
			carPtr := &workshop.Cars[len(workshop.Cars)-1]

			incType := []string{"Mecánica", "Eléctrica", "Carrocería"}[i%3]
			incidence := Incidence{
				Id:          1,
				Type:        incType,
				Priority:    "Media",
				Description: fmt.Sprintf("Test %s", incType),
				Status:      "Abierta",
			}
			carPtr.Incidences = append(carPtr.Incidences, incidence)
			incPtr := &carPtr.Incidences[0]

			cw.AddCarToQueue(carPtr, incPtr)
			time.Sleep(100 * time.Millisecond)
		}

		time.Sleep(90 * time.Second)

		duration := time.Since(startTime)
		total, busy, available := cw.GetStats()

		t.Logf("\n RESULTADOS - 6 MECÁNICOS:")
		t.Logf("   Tiempo total: %.2f segundos", duration.Seconds())
		t.Logf("   Mecánicos totales: %d", total)
		t.Logf("   Mecánicos contratados: %d", total-6)
		t.Logf("   Mecánicos ocupados al final: %d", busy)
		t.Logf("   Mecánicos disponibles: %d\n", available)
	})
}

// Test 3: Comparativa con distribuciones diferentes de especialidades
// Caso A: 3 mecánica, 1 eléctrica, 1 carrocería (5 total)
// Caso B: 1 mecánica, 3 eléctrica, 3 carrocería (7 total)
func TestDifferentSpecialityDistribution(t *testing.T) {
	fmt.Println()
	fmt.Println("======================================================================")
	fmt.Println("TEST 3: DIFERENTES DISTRIBUCIONES DE ESPECIALIDADES")
	fmt.Println("======================================================================")
	fmt.Println()

	numCars := 15 // Misma carga

	// Caso A: 3 mecánica, 1 eléctrica, 1 carrocería
	t.Run("3_mecanica_1_electrica_1_carroceria", func(t *testing.T) {
		workshop := createTestWorkshop(5, 3, 1, 1)
		cw := NewConcurrentWorkshop(&workshop)
		cw.StartDispatcher()
		cw.StartMonitor()
		cw.StartStatsHandler()
		defer cw.Stop()

		startTime := time.Now()

		// Añadir coches (5 de cada tipo)
		for i := 1; i <= numCars; i++ {
			car := Car{
				RegistrationNumber: fmt.Sprintf("CAR-%03d", i),
				Brand:              "TestBrand",
				Model:              "TestModel",
			}
			workshop.Cars = append(workshop.Cars, car)
			carPtr := &workshop.Cars[len(workshop.Cars)-1]

			incType := []string{"Mecánica", "Eléctrica", "Carrocería"}[i%3]
			incidence := Incidence{
				Id:          1,
				Type:        incType,
				Priority:    "Media",
				Description: fmt.Sprintf("Test %s", incType),
				Status:      "Abierta",
			}
			carPtr.Incidences = append(carPtr.Incidences, incidence)
			incPtr := &carPtr.Incidences[0]

			cw.AddCarToQueue(carPtr, incPtr)
			time.Sleep(100 * time.Millisecond)
		}

		time.Sleep(120 * time.Second)

		duration := time.Since(startTime)
		total, busy, available := cw.GetStats()

		t.Logf("\n RESULTADOS - 3 MECÁNICA, 1 ELÉCTRICA, 1 CARROCERÍA:")
		t.Logf("   Tiempo total: %.2f segundos", duration.Seconds())
		t.Logf("   Mecánicos totales: %d", total)
		t.Logf("   Mecánicos contratados: %d", total-5)
		t.Logf("   Mecánicos ocupados al final: %d", busy)
		t.Logf("   Mecánicos disponibles: %d\n", available)
	})

	// Caso B: 1 mecánica, 3 eléctrica, 3 carrocería
	t.Run("1_mecanica_3_electrica_3_carroceria", func(t *testing.T) {
		workshop := createTestWorkshop(7, 1, 3, 3)
		cw := NewConcurrentWorkshop(&workshop)
		cw.StartDispatcher()
		cw.StartMonitor()
		cw.StartStatsHandler()
		defer cw.Stop()

		startTime := time.Now()

		// Añadir mismos coches
		for i := 1; i <= numCars; i++ {
			car := Car{
				RegistrationNumber: fmt.Sprintf("CAR-%03d", i),
				Brand:              "TestBrand",
				Model:              "TestModel",
			}
			workshop.Cars = append(workshop.Cars, car)
			carPtr := &workshop.Cars[len(workshop.Cars)-1]

			incType := []string{"Mecánica", "Eléctrica", "Carrocería"}[i%3]
			incidence := Incidence{
				Id:          1,
				Type:        incType,
				Priority:    "Media",
				Description: fmt.Sprintf("Test %s", incType),
				Status:      "Abierta",
			}
			carPtr.Incidences = append(carPtr.Incidences, incidence)
			incPtr := &carPtr.Incidences[0]

			cw.AddCarToQueue(carPtr, incPtr)
			time.Sleep(100 * time.Millisecond)
		}

		time.Sleep(120 * time.Second)

		duration := time.Since(startTime)
		total, busy, available := cw.GetStats()

		t.Logf("\n RESULTADOS - 1 MECÁNICA, 3 ELÉCTRICA, 3 CARROCERÍA:")
		t.Logf("   Tiempo total: %.2f segundos", duration.Seconds())
		t.Logf("   Mecánicos totales: %d", total)
		t.Logf("   Mecánicos contratados: %d", total-7)
		t.Logf("   Mecánicos ocupados al final: %d", busy)
		t.Logf("   Mecánicos disponibles: %d\n", available)
	})
}

// Función helper para crear workshop de prueba
func createTestWorkshop(total, motor, electricidad, carroceria int) Workshop {
	workshop := Workshop{
		Mecanics: []Mechanic{},
		Cars:     []Car{},
		Clients:  []Client{},
	}

	id := 1

	// Añadir mecánicos de Motor
	for i := 0; i < motor; i++ {
		workshop.Mecanics = append(workshop.Mecanics, Mechanic{
			Id:              id,
			Name:            fmt.Sprintf("MecMotor-%d", i+1),
			Speciality:      "Motor",
			YearsExperience: 5,
			Status:          "Disponible",
			Activity:        "Alta",
		})
		id++
	}

	// Añadir mecánicos de Electricidad
	for i := 0; i < electricidad; i++ {
		workshop.Mecanics = append(workshop.Mecanics, Mechanic{
			Id:              id,
			Name:            fmt.Sprintf("MecElec-%d", i+1),
			Speciality:      "Electricidad",
			YearsExperience: 5,
			Status:          "Disponible",
			Activity:        "Alta",
		})
		id++
	}

	// Añadir mecánicos de Carrocería
	for i := 0; i < carroceria; i++ {
		workshop.Mecanics = append(workshop.Mecanics, Mechanic{
			Id:              id,
			Name:            fmt.Sprintf("MecCarr-%d", i+1),
			Speciality:      "Carrocería",
			YearsExperience: 5,
			Status:          "Disponible",
			Activity:        "Alta",
		})
		id++
	}

	return workshop
}

// Test adicional: Verificar contratación automática cuando se acumula > 15s
func TestAutoHiringOver15Seconds(t *testing.T) {
	fmt.Println()
	fmt.Println("======================================================================")
	fmt.Println("TEST ADICIONAL: CONTRATACIÓN AUTOMÁTICA > 15 SEGUNDOS")
	fmt.Println("======================================================================")
	fmt.Println()

	// Workshop con solo 1 mecánico de carrocería (11s de trabajo)
	workshop := createTestWorkshop(1, 0, 0, 1)
	cw := NewConcurrentWorkshop(&workshop)
	cw.StartDispatcher()
	cw.StartMonitor()
	cw.StartStatsHandler()
	defer cw.Stop()

	startTime := time.Now()
	initialMechanics := len(workshop.Mecanics)

	// Añadir un coche que necesitará >15s de trabajo
	car := Car{
		RegistrationNumber: "TEST-LONG",
		Brand:              "TestBrand",
		Model:              "TestModel",
	}
	workshop.Cars = append(workshop.Cars, car)
	carPtr := &workshop.Cars[0]

	// Incidencia de carrocería (11s cada una)
	incidence := Incidence{
		Id:          1,
		Type:        "Carrocería",
		Priority:    "Alta",
		Description: "Trabajo largo",
		Status:      "Abierta",
	}
	carPtr.Incidences = append(carPtr.Incidences, incidence)
	incPtr := &carPtr.Incidences[0]

	cw.AddCarToQueue(carPtr, incPtr)

	// Esperar suficiente tiempo para que acumule >15s
	time.Sleep(25 * time.Second)

	duration := time.Since(startTime)
	total, _, _ := cw.GetStats()
	hired := total - initialMechanics

	t.Logf("\n RESULTADOS - CONTRATACIÓN AUTOMÁTICA:")
	t.Logf("   Tiempo total: %.2f segundos", duration.Seconds())
	t.Logf("   Mecánicos iniciales: %d", initialMechanics)
	t.Logf("   Mecánicos finales: %d", total)
	t.Logf("   Mecánicos contratados: %d", hired)

	if hired > 0 {
		t.Logf("    Se contrataron mecánicos cuando se superaron 15s\n")
	} else {
		t.Logf("     No se contrataron mecánicos adicionales\n")
	}
}
