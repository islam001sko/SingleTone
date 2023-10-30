package main

import "fmt"

type micro struct {
	isRunning bool
}

func (m *micro) on() {
	m.isRunning = true
	fmt.Println("Turning micro on")
}

func (m *micro) off() {
	m.isRunning = false
	fmt.Println("Turning micro off")
}
