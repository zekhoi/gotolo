package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

type Data struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func WaterStatus(number float64) string {
	tipe := "[Air] "
	satuan := fmt.Sprint(number, " meter")
	if number < 5 {
		return tipe + "Aman: " + satuan
	} else if number >= 6 && number < 10 {
		return tipe + "Siaga: " + satuan
	} else {
		return tipe + "Bahaya: " + satuan
	}
}

func WindStatus(number float64) string {
	tipe := "[Angin] "
	satuan := fmt.Sprint(number, " meter per detik")
	if number < 5 {
		return tipe + "Aman: " + satuan
	} else if number >= 6 && number < 10 {
		return tipe + "Siaga: " + satuan
	} else {
		return tipe + "Bahaya: " + satuan
	}
}

func UpdateJson() {
	const filedir = "./constants/data.json"
	file, err := os.OpenFile(filedir, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	defer file.Close()
	byteJson, _ := ioutil.ReadAll(file)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteJson), &result)
	result["status"].(map[string]interface{})["water"] = randomNumber(1, 100)
	result["status"].(map[string]interface{})["wind"] = randomNumber(1, 100)
	fmt.Println(result)

	newFile, _ := json.Marshal(result)
	err = ioutil.WriteFile(filedir, newFile, 0644)
	if err != nil {
		log.Fatal("Error when saving file: ", err)
	}
}

func randomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}
