package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// al usar desde algún lugar tempconv se accede a todos los métodos exportados cómo un todo aunque estén en archivos separados
