/*
ESTE ES EL ÚNICO ARCHIVO QUE SE PUEDE MODIFICAR

RECOMENDACIÓN: Solo modicar a partir de la parte
				donde se encuentran la explicación
				de las otras variables.

*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
	msg    string
)

// ============================================================================
// VARIABLES Y ESTRUCTURAS ADICIONALES
// ============================================================================

// Configuración del taller
var (
	numPlazas    = 4
	numMecanicos = 4
	cantidadA    = 5
	cantidadB    = 5
	cantidadC    = 20
)

// Estados del taller
const (
	TALLER_INACTIVO       = 0
	SOLO_CATEGORIA_A      = 1
	SOLO_CATEGORIA_B      = 2
	SOLO_CATEGORIA_C      = 3
	PRIORIDAD_CATEGORIA_A = 4
	PRIORIDAD_CATEGORIA_B = 5
	PRIORIDAD_CATEGORIA_C = 6
	NO_DEFINIDO_7         = 7
	NO_DEFINIDO_8         = 8
	TALLER_CERRADO        = 9
)

// Categorías
const (
	CATEGORIA_A = iota
	CATEGORIA_B
	CATEGORIA_C
)

// Fases
const (
	FASE_ENTRADA    = "Entrada"
	FASE_REPARACION = "Reparación"
	FASE_LIMPIEZA   = "Limpieza"
	FASE_ENTREGA    = "Entrega"
)

// Coche representa un vehículo en el taller
type Coche struct {
	ID         int
	Categoria  int
	TiempoFase int
}

func (c Coche) GetCategoriaNombre() string {
	switch c.Categoria {
	case CATEGORIA_A:
		return "Mecánica"
	case CATEGORIA_B:
		return "Eléctrica"
	case CATEGORIA_C:
		return "Carrocería"
	}
	return "Desconocida"
}

// GestorTaller maneja el estado y recursos del taller
type GestorTaller struct {
	estadoActual         int
	estadoMutex          sync.RWMutex
	plazasDisponibles    chan struct{}
	mecanicosDisponibles chan struct{}
	tiempoInicio         time.Time
	wg                   sync.WaitGroup
	cerrado              bool
	cerradoMutex         sync.RWMutex
}

var gestor *GestorTaller

func inicializarGestor() {
	gestor = &GestorTaller{
		estadoActual:         TALLER_INACTIVO,
		plazasDisponibles:    make(chan struct{}, numPlazas),
		mecanicosDisponibles: make(chan struct{}, numMecanicos),
		tiempoInicio:         time.Now(),
		cerrado:              false,
	}

	// Inicializar recursos
	for i := 0; i < numPlazas; i++ {
		gestor.plazasDisponibles <- struct{}{}
	}
	for i := 0; i < numMecanicos; i++ {
		gestor.mecanicosDisponibles <- struct{}{}
	}
}

func (gt *GestorTaller) cambiarEstado(nuevoEstado int) {
	gt.estadoMutex.Lock()
	defer gt.estadoMutex.Unlock()

	if nuevoEstado == NO_DEFINIDO_7 || nuevoEstado == NO_DEFINIDO_8 {
		return
	}

	gt.estadoActual = nuevoEstado
	fmt.Printf("[TALLER] Estado cambiado a: %d\n", nuevoEstado)

	if nuevoEstado == TALLER_CERRADO {
		gt.cerradoMutex.Lock()
		gt.cerrado = true
		gt.cerradoMutex.Unlock()
	}
}

func (gt *GestorTaller) obtenerEstado() int {
	gt.estadoMutex.RLock()
	defer gt.estadoMutex.RUnlock()
	return gt.estadoActual
}

func (gt *GestorTaller) estaCerrado() bool {
	gt.cerradoMutex.RLock()
	defer gt.cerradoMutex.RUnlock()
	return gt.cerrado
}

func (gt *GestorTaller) puedeEntrar(coche Coche) bool {
	estado := gt.obtenerEstado()

	switch estado {
	case TALLER_INACTIVO, TALLER_CERRADO:
		return false
	case SOLO_CATEGORIA_A:
		return coche.Categoria == CATEGORIA_A
	case SOLO_CATEGORIA_B:
		return coche.Categoria == CATEGORIA_B
	case SOLO_CATEGORIA_C:
		return coche.Categoria == CATEGORIA_C
	case PRIORIDAD_CATEGORIA_A, PRIORIDAD_CATEGORIA_B, PRIORIDAD_CATEGORIA_C:
		return true
	default:
		return false
	}
}

func (gt *GestorTaller) obtenerPrioridad(coche Coche) int {
	estado := gt.obtenerEstado()
	prioridad := 0

	switch coche.Categoria {
	case CATEGORIA_A:
		prioridad = 30
	case CATEGORIA_B:
		prioridad = 20
	case CATEGORIA_C:
		prioridad = 10
	}

	switch estado {
	case PRIORIDAD_CATEGORIA_A:
		if coche.Categoria == CATEGORIA_A {
			prioridad += 100
		}
	case PRIORIDAD_CATEGORIA_B:
		if coche.Categoria == CATEGORIA_B {
			prioridad += 100
		}
	case PRIORIDAD_CATEGORIA_C:
		if coche.Categoria == CATEGORIA_C {
			prioridad += 100
		}
	}

	return prioridad
}

func (gt *GestorTaller) imprimirMensaje(coche Coche, fase string, estado string) {
	tiempoTranscurrido := time.Since(gt.tiempoInicio).Seconds()
	fmt.Printf("Tiempo %.2fs Coche %d Incidencia %s Fase %s Estado %s\n",
		tiempoTranscurrido,
		coche.ID,
		coche.GetCategoriaNombre(),
		fase,
		estado)
}

// Procesar un coche a través de las 4 fases
func procesarCoche(coche Coche) {
	// Asegurar que siempre se llama Done al terminar
	defer gestor.wg.Done()

	// FASE 1: ENTRADA - Esperar hasta que pueda entrar
	gestor.imprimirMensaje(coche, FASE_ENTRADA, "Esperando")

	// Esperar hasta que el taller permita entrada
	for {
		if gestor.puedeEntrar(coche) {
			break
		}
		if gestor.estaCerrado() {
			// Si el taller está cerrado DESPUÉS de haber intentado entrar
			// verificar si ya había empezado el proceso
			estado := gestor.obtenerEstado()
			if estado == TALLER_CERRADO {
				// Esperar un poco para ver si cambia
				time.Sleep(500 * time.Millisecond)
				continue
			}
		}
		time.Sleep(100 * time.Millisecond)
	}

	// Obtener plaza
	<-gestor.plazasDisponibles
	gestor.imprimirMensaje(coche, FASE_ENTRADA, "Entrando")
	time.Sleep(time.Duration(coche.TiempoFase) * time.Second)
	gestor.imprimirMensaje(coche, FASE_ENTRADA, "Completado")

	// FASE 2: REPARACIÓN
	gestor.imprimirMensaje(coche, FASE_REPARACION, "Esperando")
	<-gestor.mecanicosDisponibles
	gestor.imprimirMensaje(coche, FASE_REPARACION, "Reparando")
	time.Sleep(time.Duration(coche.TiempoFase) * time.Second)
	gestor.imprimirMensaje(coche, FASE_REPARACION, "Completado")
	gestor.mecanicosDisponibles <- struct{}{}

	// FASE 3: LIMPIEZA
	gestor.imprimirMensaje(coche, FASE_LIMPIEZA, "Esperando")
	gestor.imprimirMensaje(coche, FASE_LIMPIEZA, "Limpiando")
	time.Sleep(time.Duration(coche.TiempoFase) * time.Second)
	gestor.imprimirMensaje(coche, FASE_LIMPIEZA, "Completado")

	// FASE 4: ENTREGA
	gestor.imprimirMensaje(coche, FASE_ENTREGA, "Esperando")
	gestor.imprimirMensaje(coche, FASE_ENTREGA, "Entregando")
	time.Sleep(time.Duration(coche.TiempoFase) * time.Second)
	gestor.imprimirMensaje(coche, FASE_ENTREGA, "Completado")

	// Liberar plaza
	gestor.plazasDisponibles <- struct{}{}
}

func generarCoches() {
	var todosLosCoches []Coche
	cocheID := 1

	for i := 0; i < cantidadA; i++ {
		todosLosCoches = append(todosLosCoches, Coche{
			ID:         cocheID,
			Categoria:  CATEGORIA_A,
			TiempoFase: 5,
		})
		cocheID++
	}

	for i := 0; i < cantidadB; i++ {
		todosLosCoches = append(todosLosCoches, Coche{
			ID:         cocheID,
			Categoria:  CATEGORIA_B,
			TiempoFase: 3,
		})
		cocheID++
	}

	for i := 0; i < cantidadC; i++ {
		todosLosCoches = append(todosLosCoches, Coche{
			ID:         cocheID,
			Categoria:  CATEGORIA_C,
			TiempoFase: 1,
		})
		cocheID++
	}

	// Mezclar aleatoriamente
	rand.Shuffle(len(todosLosCoches), func(i, j int) {
		todosLosCoches[i], todosLosCoches[j] = todosLosCoches[j], todosLosCoches[i]
	})

	fmt.Printf("[INFO] Generando %d coches...\n", len(todosLosCoches))

	// Incrementar el WaitGroup ANTES de lanzar las goroutines
	gestor.wg.Add(len(todosLosCoches))

	// Lanzar goroutines
	for _, coche := range todosLosCoches {
		go func(c Coche) {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			procesarCoche(c)
		}(coche)
	}

	fmt.Printf("[INFO] %d coches lanzados y esperando estados del taller\n", len(todosLosCoches))
}

func procesarMensaje(mensaje string) {
	// Limpiar el mensaje (eliminar saltos de línea y espacios)
	mensaje = strings.TrimSpace(mensaje)

	// Debug (comentado para salida limpia)
	// fmt.Printf("[DEBUG] Procesando: '%s'\n", mensaje)

	// Intentar extraer un número del mensaje
	// El mensaje puede venir solo o con texto adicional
	for _, parte := range strings.Fields(mensaje) {
		numero, err := strconv.Atoi(parte)
		if err == nil && numero >= 0 && numero <= 9 {
			// Es un número válido entre 0-9
			gestor.cambiarEstado(numero)
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()

	// Inicializar gestor del taller
	inicializarGestor()

	fmt.Println("===========================================")
	fmt.Printf("TALLER INICIADO\n")
	fmt.Printf("Plazas: %d | Mecánicos: %d\n", numPlazas, numMecanicos)
	fmt.Printf("Coches: A=%d, B=%d, C=%d\n", cantidadA, cantidadB, cantidadC)
	fmt.Println("===========================================")

	// Generar coches (NO en goroutine, para que wg.Add se ejecute primero)
	generarCoches()

	// Canal para señalar que debemos terminar
	done := make(chan struct{})

	// Goroutine para leer mensajes del servidor
	go func() {
		buf := make([]byte, 512)
		for {
			n, err := conn.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			if n > 0 {
				msg = string(buf[:n])
				/*
					Desde aquí debería salir la información a una goroutine o a una función ordinaria según se requiera

					n 	:=	Longitud del string enviado por el servidor
					msg := 	Mensaje recibido por la conexión del server

					Recomendación: Para usar la conversión entre string e int se recomienda usar strconv.Atoi
					Más info en: https://pkg.go.dev/strconv#Atoi

				*/
				// Debug: ver qué llega (comentado para salida limpia)
				// fmt.Printf("[DEBUG] Mensaje recibido: '%s' (len=%d)\n", msg, n)

				// Procesar el mensaje recibido
				procesarMensaje(msg)
			}
		}
	}()

	// Goroutine para esperar a que terminen todos los coches
	go func() {
		gestor.wg.Wait()
		close(done)
	}()

	// Esperar a que terminen todos los coches o timeout
	select {
	case <-done:
		fmt.Println("===========================================")
		fmt.Println("TALLER FINALIZADO - Todos los coches procesados")
		fmt.Println("===========================================")
	case <-time.After(120 * time.Second):
		fmt.Println("===========================================")
		fmt.Println("TALLER FINALIZADO - Timeout alcanzado")
		fmt.Println("===========================================")
	}
}
