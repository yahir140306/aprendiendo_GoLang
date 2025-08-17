package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// las interfaces no se implementan al menos manualmente
// las interfaces se cumplen

type forma interface {
	Area() float64
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func imprimirArea(f forma) {
	fmt.Printf("El area es: %.2f \n", f.Area())
}

func interfaces() {
	c := Circulo{Radio: 5}

	imprimirArea(c)

}

// concurrencia

// go => GoRoutine
// GoRoutine => es un hilo de ejecucion ligero virtual
// channel => para comunicar cosas entre GoRoutine

func decirHola(canal chan<- string) {
	time.Sleep(1 * time.Second)
	canal <- "Hola desde la GoRoutine"
}

func imprimirMensaje(canal <-chan string) {
	fmt.Println("Esperando mensaje...")
	msg := <-canal
	fmt.Println(msg)
}

func concurrencia() {
	canal := make(chan string)
	go decirHola(canal)
	imprimirMensaje(canal)

	canal2 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			canal2 <- i
		}
		close(canal2)
	}()

	for num := range canal2 {
		fmt.Println("Numero recibido:", num)
	}

	// Ejemplo Mutex
	var contador int
	var mu sync.RWMutex

	// writer
	go func() {
		for i := 0; i < 5; i++ {
			mu.Lock()
			contador++
			mu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// reader
	for i := 0; i < 3; i++ {
		go func(id int) {
			for j := 0; j < 5; j++ {
				mu.RLock()
				fmt.Printf("Leyendo desde la GoRoutine %d: %d\n", id, contador)
				mu.RUnlock()
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
}
