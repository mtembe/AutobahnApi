package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/pkg/browser"
)

type ApiAntwort struct {
	Raststaetten []Raststaette `json:"parking_lorry"`
}

type Raststaette struct {
	Extent                   string
	Identifier               string
	Title                    string
	Point                    string
	Subtitle                 string                     `json:"subtitle"`
	Coordinates              Coordinates                `json:"coordinate"`
	LorryParkingFeatureIcons []LorryParkingFeatureIcons `json:"lorryParkingFeatureIcons"`
}

type Coordinates struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"long"`
}

type LorryParkingFeatureIcons struct {
	Zusatzinfos1 Picknickmoeglichkeit `json:"0"`
	Zusatzinfos2 Toiletten            `json:"1"`
}

type Picknickmoeglichkeit struct {
	DescriptionP string `json:"description"`
}

type Toiletten struct {
	DescriptionT string `json:"description"`
}

func main() {

	//Eingabeaufforderung
	var Input1 string
	fmt.Println("Hello User!")
	fmt.Println("Zu welcher Autobahn möchtest du eine Liste der Rastplaetze?")
	fmt.Scan(&Input1)

	hostname := "https://verkehr.autobahn.de/o/autobahn/"
	path1 := "//"
	path2 := Input1
	path3 := "//services/parking_lorry"

	result, err := url.JoinPath(hostname, path1, path2, path3)
	if err != nil {
		log.Fatal(err)
	}

	apiurl := result
	resp, err := http.Get(apiurl)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var Liste ApiAntwort
	err = json.Unmarshal(body, &Liste)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for i, p := range Liste.Raststaetten {
		fmt.Println("Raststaette", (i + 1), ":", p.Subtitle)
	}

	//Eingabeaufforderung Rastplatzspezifizierung

	var Input2 int
	fmt.Println("Zu welchem Rastplatz möchtest du mehr Informationen? Tippe die Rastplatznummer!")
	fmt.Scan(&Input2)

	//Weitere Details zum spezifizierten Rastplatz ausgeben
	x := (Input2 - 1)
	fmt.Println("Der Name der Raststätte lautet: ", Liste.Raststaetten[x].Subtitle)
	fmt.Println("Die Koordinaten lauten: ", Liste.Raststaetten[x].Coordinates.Latitude)

	//Ausgegebene Koordinaten in google maps url einfügen

	hostname2 := "https://www.google.com/maps/place/" + Liste.Raststaetten[x].Coordinates.Latitude + "," + Liste.Raststaetten[x].Coordinates.Longitude

	fmt.Print(hostname2)

	_ = browser.OpenURL(hostname2) //https://www.google.com/maps/place/lat,lng

}
