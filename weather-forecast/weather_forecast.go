//Package weather provides a tool for showing the wheather
//condition for a specific location.
package weather

//CurrentCondition represents the current wheather condition.
var CurrentCondition string

//CurrentLocation represents the current Location.
var CurrentLocation string

//Forecast returns a string which shows current location and its corresponding weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
