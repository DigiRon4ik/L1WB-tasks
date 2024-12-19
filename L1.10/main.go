package main

import (
	"fmt"
	"sort"
	"strconv" // I could have done with just math, but I'm more comfortable that way. At least here you can.
)

func main() {
	// Define a slice of float64 representing temperatures
	temperatures := []float64{-25.4, -27.0, 13.0, -0.0, 19.0, -4.7, 15.5, 7.0, 24.5, -21.0, 32.5}

	// Group temperatures using the temperatureGrouping function
	temperGroups := temperatureGrouping(temperatures)

	// Create a slice to store keys from the map
	var keys []int
	// Iterate over the map keys and append them to the keys slice
	for key, _ := range temperGroups {
		keys = append(keys, key)
	}
	// Sort the keys slice in ascending order
	sort.Ints(keys)
	// Iterate over sorted keys and print the grouped temperatures
	for _, key := range keys {
		fmt.Printf("%4d: %v\n", key, temperGroups[key])
	}
}

// Function to group temperatures based on their integer part
func temperatureGrouping(tempers []float64) map[int][]float64 {
	// Initialize a map to store grouped temperatures
	temperGroups := make(map[int][]float64)

	// Iterate over each temperature in the slice
	for _, temp := range tempers {
		// Здесь можно использовать горутину с мютексом или с sync.Map
		// Convert the temperature to an integer string
		key := strconv.Itoa(int(temp))
		// Determine the key prefix based on the sign of the temperature
		if key[0] == '-' {
			key = key[:2] // Use first two characters for negative numbers
		} else {
			key = key[:1] // Use first character for non-negative numbers
		}
		// Convert the key prefix back to an integer
		keyInt, err := strconv.Atoi(key)
		// Panic if there's an error in conversion
		if err != nil {
			panic(err)
		}
		// Adjust the key for temperatures below -9 or above 9
		if temp < -9 || temp > 9 {
			keyInt *= 10
		}
		// Append the temperature to the corresponding group in the map
		temperGroups[keyInt] = append(temperGroups[keyInt], temp)
	}

	// Return the map with grouped temperatures
	return temperGroups
}

/*
- Input : -
[-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5]
- Output: -
 -20: [-25.4 -27 -21]
  10: [13 19 15.5]
  20: [24.5]
  30: [32.5]
*/

/*
- Input : -
[-25.4, -27.0, 13.0, -0.0, 19.0, -4.7, 15.5, 7.0, 24.5, -21.0, 32.5]
- Output: -
 -20: [-25.4 -27 -21]
  -4: [-4.7]
   0: [0]
   7: [7]
  10: [13 19 15.5]
  20: [24.5]
  30: [32.5]
*/
