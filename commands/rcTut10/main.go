package main

import "fmt"

type GeoLocation struct {
	Lat float64
	Lng float64
}

//Pointers
func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j


	//Structures
	location := GeoLocation{39.2, -119.5}
	locationAddress := &location
	fmt.Println(locationAddress.Lat)

}

func addGeoLocation (loc *GeoLocation) float64{
	return loc.Lat + loc.Lng
}
