// Package tempconv realiza conversiones Celsius y Farhenheit.
// dejar esa descripción para documentar el paquete
package tempconv

import "fmt"

// ya descrito en ../variables/main.go

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)
