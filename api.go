package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

//Athan grabs main data header from Json
type Athan struct {
	Data YtData `json:"data"`
}

//YtData grabs secondary header under "Athan"
type YtData struct {
	Timings YtTime `json:"timings"`
}

//YtTime Gets 3rd Header from Json file, which is where the athan times are located
type YtTime struct {
	F string `json:"Fajr"`
	D string `json:"Dhuhr"`
	A string `json:"Asr"`
	M string `json:"Maghrib"`
	I string `json:"Isha"`
}

//Config Get Config settings from config.json file
type Config struct {
	Settings struct {
		Name     string `json:"Name"`
		Language string `json:"Language"`
		Accent   string `json:"Accent"`
		Athan    string `json:"Athan"`
	}

	Prayers struct {
		Fajir  bool `json:"Fajir"`
		Duhur  bool `json:"Duhur"`
		Asr    bool `json:"Asr"`
		Magrib bool `json:"Magrib"`
		Isha   bool `json:"Isha"`
	}

	Audio struct {
		Athan  string `json:"Athan"`
		Recite string `json:"Recite"`
	}

	Location struct {
		City     string `json:"City"`
		Country  string `json:"Country"`
		State    string `json:"State"`
		TimeZone string `json:"TimeZone"`
	}

	Calculation struct {
		Method int `json:"Method"`
	}

	Volume struct {
		Connection bool    `json:"Connection"`
		Default    float64 `json:"Default"`
		Fajir      float64 `json:"Fajir"`
		Duhur      float64 `json:"Duhur"`
		Asr        float64 `json:"Asr"`
		Magrib     float64 `json:"Magrib"`
		Isha       float64 `json:"Isha"`
	}

	Options struct {
		Whisper bool `json:"Whisper"`
		Recite  bool `json:"Recite"`
	}
}

//Split API
const (
	MainAPI    string = "http://api.aladhan.com/v1/timingsByCity?city="
	CountryAPI string = "&country="
	StateAPI   string = "&state="
	MethodAPI  string = "&method="
)

var Meth = strconv.Itoa(config.Calculation.Method)

var config Config
var entry string

func ACal() Athan {
	var AthanAPI = MainAPI + config.Location.City + CountryAPI + config.Location.Country + StateAPI + config.Location.State + MethodAPI + Meth
	FormatAPI := fmt.Sprintf(AthanAPI)

	//fmt.Println(AthanAPI)

	resp, err := http.Get(FormatAPI)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	var Y Athan
	err = json.Unmarshal(body, &Y)
	if err != nil {
		log.Fatal(err)
	}

	return Y
}

//LoadConfig file
func LoadConfig(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)

	defer configFile.Close()
	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}
