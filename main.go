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
		Addr:     "my-release-redis-master.default.svc.cluster.local:6379",
		Password: "13371337",
		DB:       0,
	})

	totalpkwbuchten := 0
	totallkwbuchten := 0
	totalNoHighways := 0

	HighwayName := flag.String("roads", "", "")
	flag.Parse()
	str1slice := strings.Split(*HighwayName, ",")
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

		key := x

		data := fmt.Sprintf(`{"PKW":%d,"LKW":%d,"Total":%d}`, bab.PKW, bab.LKW, bab.Sum())

		err = client.Set(context.Background(), key, data, 0).Err()
		if err != nil {
			fmt.Println("Error setting value in Redis for Autobahn", x, ":", err)
		} else {
			fmt.Println("Success setting value in Redis for Autobahn", x)
		}

	}
	fmt.Println("_________________________")
	fmt.Println("Anzahl der ausgewerteten Autobahnen: ", totalNoHighways)
	fmt.Println("Gesamtanzahl PKW Buchten: ", totalpkwbuchten)
	fmt.Println("Gesamtanzahl LKW Buchten: ", totallkwbuchten)

}
