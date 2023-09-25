package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Status Data `json:"status"`
}

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	go reload()   // <--- running otomatis melakukan hitting ke startAPIServer
	startServer() // <-- sebatas penyedia web server dan service buat ke logic database dan update datanya.
}

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/reload", autoReloadHandler)

	fmt.Println("Listening to port 3030")
	err := http.ListenAndServe(":3030", mux)
	log.Fatal(err)
}

func autoReloadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Auto-reload started")

	// script koneksi database, script update data, logging water dan wind

	//...

}

func reload() {
	for {
		min := 1
		max := 100
		rand.Seed(time.Now().Unix())
		numberWater := rand.Intn(max - min)
		numberWind := rand.Intn(max - min)

		updateData(numberWater, numberWind)
		logWaterandWind(numberWater, numberWind)

		makeAPIRequest(numberWater, numberWind)

		time.Sleep(time.Second * 15)
	}
}

func updateData(numberWater int, numberWind int) {
	data := Data{}
	data.Water = numberWater
	data.Wind = numberWind
	dataWaterWind := Status{
		Status: data,
	}

	jsonprint, err := json.MarshalIndent(dataWaterWind, "", "    ")
	if err != nil {
		fmt.Println("Error print JSON")
		return
	}

	err = os.WriteFile("data.json", jsonprint, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonprint))
}

func logWaterandWind(numberWater int, numberWind int) {
	if numberWater < 6 {
		result := "Aman"
		fmt.Println("Status Water: ", result)
	} else if numberWater > 6 && numberWater <= 15 {
		result := "Siaga"
		fmt.Println("Status Water: ", result)
	} else if numberWater > 15 {
		result := "Bahaya"
		fmt.Println("Status Water: ", result)
	} else {
		fmt.Println("Not found")
	}

	if numberWind < 6 {
		result := "Aman"
		fmt.Println("Status Wind: ", result)
	} else if numberWind > 6 && numberWind <= 15 {
		result := "Siaga"
		fmt.Println("Status Wind: ", result)
	} else if numberWind > 15 {
		result := "Bahaya"
		fmt.Println("Status Wind: ", result)
	} else {
		fmt.Println("Measurement not found")
	}
}

func makeAPIRequest(numberWater int, numberWind int) {
	data := Data{
		Water: numberWater,
		Wind:  numberWind,
	}
	status := Status{
		Status: data,
	}

	requestJson, err := json.Marshal(status)
	if err != nil {
		fmt.Println("json error:", err)
		return
	}

	resp, err := http.Post("http://localhost:3030/reload", "application/json", bytes.NewBuffer(requestJson))
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("response error:", err)
		return
	}

	fmt.Println(string(res))

	fmt.Println("API request made")
}
