//package massconv converts between kg and lbs

package massconv

import "fmt"

type Kilogram float64
type Pound float64

func (k Kilogram) String() string { return fmt.Sprintf("%.1fkg", k)}
func (p Pound) String() string { return fmt.Sprintf("%.1flbs", p)}

func KgToLb(k Kilogram) Pound { return Pound(k/2.205) }
func LbToKg(p Pound) Kilogram { return Kilogram(p*2.205) }