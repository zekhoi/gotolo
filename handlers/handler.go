package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tolol/utils"
)

type Status struct {
	Water string
	Wind  string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	raw, err := ioutil.ReadFile("./constants/data.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(raw), &result)
	wtr := result["status"].(map[string]interface{})["water"]
	wnd := result["status"].(map[string]interface{})["wind"]

	water := utils.WaterStatus(wtr.(float64))
	wind := utils.WindStatus(wnd.(float64))

	status := fmt.Sprint(water, "\n", wind)
	w.Write([]byte(status))

}
