package tempconv

func CToF(c Celsius) Farenheit {
  return Farenheit(c*9/5 + 32)
}

func CToK(c Celsius) Kelvin {
    return Kelvin(c+AbsoluteZeroC)
}

func FToC(f Farenheit) Celsius {
  return Celsius((f - 32) * 5/9)
}

func FToK(f Farenheit) Kelvin {
  return CToK(FToC(f))
}

func KToC(k Kelvin) Celsius {
    return Celsius(k - AbsoluteZeroC)
}

func KToF(k Kelvin) Farenheit {
    return Farenheit(CToF(KtoC(k)))
}
