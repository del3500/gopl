package tempconv

// CToF converts a Celsius temperator to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK convets a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(C + 273.15) }

// FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts a Fahrenheit temperature to Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin(f-32)*(5/9) + 273.15 }

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celcius { return Celius(k - 273.15) }

// KToF converts a Kelvin temperature to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*(9/5) + 32) }
