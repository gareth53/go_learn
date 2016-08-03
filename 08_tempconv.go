// package that handles temperatir conversions - Celsius, Farenheit & Kelvin
package tempconv

import "fmt"

type Celsius float64
type Farenheit float64
type Kelvin float64

const (
  AbsoluteZeroC Celsius = -273.15
  FreezingC     Celsius = 0
  BoilingC      Celsius = 100
)

func (c Celsius) String() string {
  return fmt.Sprintf("%g˚C", c)
}

func (f Farenheit) String() string {
  return fmt.Sprintf("%g˚F", f)
}

func (k Kelvin) String() string {
  return fmt.Sprintf("%g˚K", k)
}

func CToF(c Celsius) Farenheit {
  return Farenheit(c*9/5 + 32)
}

func CToK(c Celsius) Kelvin {
    return Kelvin(c + Celsius(AbsoluteZeroC))
}

func FToC(f Farenheit) Celsius {
  return Celsius((f - 32) * 5/9)
}

func FToK(f Farenheit) Kelvin {
  return CToK(FToC(f))
}

func KToC(k Kelvin) Celsius {
    return Celsius(k - Kelvin(AbsoluteZeroC))
}

func KToF(k Kelvin) Farenheit {
    return Farenheit(CToF(KToC(k)))
}
