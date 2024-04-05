package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"

	"AutobahnApiGo/webserver/stellplatzanzahl"
)

func Bundesapi(MissingHighwayNo string) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	totalpkwbuchten := 0
	totallkwbuchten := 0
	totalNoHighways := 0

	str1slice := strings.Split(MissingHighwayNo, ",")
	fmt.Printf("%v", str1slice)

	for i, x := range str1slice {

		bab, err := stellplatzanzahl.ParkinglorrySum(x)
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
	fmt.Println("Gesamtanzahl der PKW Buchten: ", totalpkwbuchten)
	fmt.Println("Gesamtanzahl der LKW Buchten: ", totallkwbuchten)

}
