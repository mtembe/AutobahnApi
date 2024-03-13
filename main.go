package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"

	"Stellplatzanzahl/Stellplatzanzahl"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	totalpkwbuchten := 0
	totallkwbuchten := 0
	totalNoHighways := 0

	HighwayNo := flag.String("roads", "", "")
	flag.Parse()
	str1slice := strings.Split(*HighwayNo, ",")
	fmt.Printf("%v", str1slice)

	for i, x := range str1slice {

		bab, err := Stellplatzanzahl.ParkinglorrySum(x)
		fmt.Println("Autobahn :", x)
		fmt.Println("Anzahl PKWbuchten: ", bab.PKW)
		fmt.Println("Anzahl LKWbuchten: ", bab.LKW)
		fmt.Println("Gesamtanzahl", bab.Sum())
		fmt.Println("_________________________")
		if err != nil {
			fmt.Println("Error: ", err)

		}
		totalNoHighways = (i + 1)
		totalpkwbuchten += bab.PKW
		totallkwbuchten += bab.LKW

		err = client.JSONSet(context.Background(), x, "$", bab).Err()
		if err != nil {
			fmt.Println("error setting value in redis", err)

		} else {
			fmt.Println("Great Succes while setting value in redis")
		}

	}
	fmt.Println("_________________________")
	fmt.Println("Anzahl der ausgewerteten Autobahnen: ", totalNoHighways)
	fmt.Println("Gesamtanzahl PKW Buchten: ", totalpkwbuchten)
	fmt.Println("Gesamtanzahl LKW Buchten: ", totallkwbuchten)

}
