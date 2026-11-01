// Package weather provides weather forecast functions.
package weather

var (
	// CurrentCondition tells about weather conditions.
	CurrentCondition string
	// CurrentLocation tells about weather location.
	CurrentLocation string
)

// Forecast format weather forecast for given location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
