//package tempconv performs Celcius, Fahrenheit, 
//and Kelvin conversions.

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const(
	AbsoluteZeroC Celsius = -273.15
	FreezingC = 0
	BoilingC = 100
)

const(
	AbsoluteZeroK = 0
)

func (c Celsius) String() string { return fmt.Sprintf("%gdeg C", c)}
func (f Fahrenheit) String() string { return fmt.Sprintf("%gdeg F", f)}
func (k Kelvin) String() string { return fmt.Sprintf("%gdeg K", k)}

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32)}

//FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}

//CToK converts Celsius to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15)}

//KToC converts Kelvin to Celsius 
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15)}
