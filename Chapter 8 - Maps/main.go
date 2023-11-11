package main

/*

Maps are a way to associate a Key with a Value (K:V)

The zero value of a map is nil


make(map[string]int)
Creates a map of string to integer
We are mapping string into integers

ages := make(map[string]int)
ages["John"] = 37

or

ages = map[string]int {
	"John": 37
}


Mutations:

-- Insert an element

m[key] = elem

-- Get an element

elem = m[key]

-- Delete an element
delete(m, key)

-- check if a key exists
elem, ok := m[key]
if `key` is in `m` `ok` is true otherwise is false



We can create a nesting maps but we can just use a struct

ages := make(map[string]map[string]int)

ages["john"]["vski"] = 11


Instead of this we could do

type Key struct {
	Path, Country string
}


hits := make(map[Key]int)



======== Effective Go


Like slices, maps hold references

So when we pass a map to a function that changes the contents of the map, the change will be visible in the caller


*/

func main() {

}
