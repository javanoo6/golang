
package main

import (
	"fmt"
)

func main() {
	jewelry := make(map[string]float64)
	jewelry["neclace"] = 89.99
	jewelry["earrings"] = 79.99
	clothing := map[string]float64{"pants": 59.99, "shirt": 39.99}
	fmt.Println("Earrings", jewelry["earrings"])
	fmt.Println("Necklace", jewelry["neclace"])
	fmt.Println("Shirt",clothing["shirt"])
	fmt.Println("Pants",clothing["pants"])
}
