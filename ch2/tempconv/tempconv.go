package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
const (
	AbsoluteZeroC Celsius = -273.15
	Freezingc     Celsius = 0
	Boilingc      Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g-degree C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g-degree F", f) }