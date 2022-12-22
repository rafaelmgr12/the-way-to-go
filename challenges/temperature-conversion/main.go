/*
Challenge: Temperature Conversion
Implement function toFahrenheit(t Celsius)to convert from Celsius to Fahrenheit. Types of temperatures are float32 which are aliased to Celsius and Fahrenheit for you.

Temperature in ˚F = (Temp in °C * 9/5) + 32
(Tempin°C∗9/5)+32

*/
package main
import "fmt"
import "strconv"
import "encoding/json"

// aliasing type
type Celsius float32
type Fahrenheit float32

// Function to convert celsius to fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {
	
	var temp Fahrenheit
	
	temp = Fahrenheit((t * 9/5) + 32)

	return temp

}