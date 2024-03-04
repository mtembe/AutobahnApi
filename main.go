package main

import (
	"flag"
	"fmt"
	"strings"

	"Stellplatzanzahl/Stellplatzanzahl"
)

func main() {
	totalpkwbuchten := 0
	totallkwbuchten := 0
	totalNoHighways := 0

	HighwayNo := flag.String("roads", "", "")
	flag.Parse()
	str1slice := strings.Split(*HighwayNo, ",")

	for i, x := range str1slice {

		a, b, err := Stellplatzanzahl.ParkinglorrySum(x)
		fmt.Println("Autobahn :", x)
		fmt.Println("Anzahl PKWbuchten: ", a)
		fmt.Println("Anzahl LKWbuchten: ", b)
		fmt.Println("_________________________")
		if err != nil {
			fmt.Println("Error: ", err)

		}
		totalNoHighways = (i + 1)
		totalpkwbuchten += a
		totallkwbuchten += b

	}
	fmt.Println("_________________________")
	fmt.Println("Anzahl der ausgewerteten Autobahnen: ", totalNoHighways)
	fmt.Println("Gesamtanzahl PKW Buchten: ", totalpkwbuchten)
	fmt.Println("Gesamtanzahl LKW Buchten: ", totallkwbuchten)
}

// import (
// 	"flag"
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	str1 := flag.String("roads", "", "")

// 	flag.Parse()
// 	fmt.Println(*str1)

// 	str1slice := strings.Split(*str1, ",")
// 	fmt.Println(str1slice)

// 	for _, bahn := range str1slice {
// 		fmt.Println(bahn)
// 	}

// }

//original

// package main

// import (
// 	"fmt"

// 	"Stellplatzanzahl/Stellplatzanzahl"
// )

// func main() {
// 	HighwayNo := "A1"
// 	a, b, err := Stellplatzanzahl.ParkinglorrySum(HighwayNo)
// 	fmt.Println("Anzahl PKWbuchten: ", a)
// 	fmt.Println("Anzahl LKWbuchten: ", b)
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 	}
// }
