//package lenconv performs mile and kilometer conversions.

package lenconv

import "fmt"

type Mile float64
type Kilometer float64

func (m Mile) String() string { return fmt.Sprintf("%.1f Miles", m)}
func (km Kilometer) String() string { return fmt.Sprintf("%.1f Kilometers", km)}

func MToKm(m Mile) Kilometer  {return Kilometer(m/1.609)}
func KmToM(km Kilometer) Mile {return Mile(km * 1.609)}