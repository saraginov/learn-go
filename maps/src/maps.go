package main

import(
	"fmt"
)

type Vertex struct {
	Latitude, Longitude float64
}

/*
	Maps

		A map, maps keys to values.

		The zero value of a map is nil.
		A nil map has no keys, nor can keys be added.

		The make function returns a map of the given type, 
		initialized and ready for use.

		when printing a map it has the same structure as an array and slice,
		the difference being that the elements in a map are key value:pairs
		rather than just values
*/
var m map[string]Vertex

/*
	Map Literals

		Map literals are like struct literals,
		but the keys are required

		If the top-level type is just a type name,
		you can omit it from the elements of the literal.

		thus:
			var newM = map[string]Vertex{
				"Bell Labs": { 40.68433, -74.39967,},
				"Google": { 37.42202, -122.08408,},
			}
*/

var newM = map[string]Vertex{
	"Bell Labs": Vertex{ 40.68433, -74.39967,},
	"Google": Vertex{ 37.42202, -122.08408,},
}

func main(){
	fmt.Println(m) // output: map[] ie nil

	m = make(map[string]Vertex) // we re-initialize m in main(){...} scope

	fmt.Println(m) // output: map[] ie nil, but because of make() we can now add k:v pairs
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

	// printing map literal
	fmt.Println(newM)
}