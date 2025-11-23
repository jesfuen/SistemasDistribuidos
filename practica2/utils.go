package main

import "fmt"

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
