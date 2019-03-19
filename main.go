package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	var err error
	//Connect to Json file for settings and paramaters
	config, err = LoadConfig("config.json")
	if err != nil {
		log.Fatal("Error importing config.json file", err)
	}

	fmt.Println("Connected")
	//cli.Notify("test")
	Y := ACal()

	for range time.Tick(time.Second * 8) {
		//Get Local time test
		t := time.Now()
		location, err := time.LoadLocation(config.Location.TimeZone)
		if err != nil {
			fmt.Println(err)
		}

		CurrentTime := fmt.Sprint(t.In(location).Format("15:04"))

		//Check if friday
		day := time.Now().Weekday()
		CurrentDay := fmt.Sprint(day)

		//cli.Notify("test")

		//Checks if its time for Fajir
		if Y.Data.Timings.F == CurrentTime {
			if config.Prayers.Fajir == true {
				//fmt.Println("Time for Fajir")
				Fajr()
			}
		}

		//Checks if its time for Duhur
		if Y.Data.Timings.D == CurrentTime {
			if config.Prayers.Duhur == true {
				//fmt.Println("Time for Duhur")
				Duhur()
			}
		}

		//Checks if the day is Friday
		if config.Options.Recite == true {
			if CurrentDay == "Friday" {
				//cli.Notify("I will begin reciting Quran.")
				time.Sleep(5 * time.Second)
				//cli.Play(config.Audio.Recite)
			}
		}

		if config.Prayers.Duhur == true {
			Duhur()
			time.Sleep(4 * time.Minute)
		}

		//Checks if its time for Asr
		if Y.Data.Timings.A == CurrentTime {
			if config.Prayers.Asr == true {
				//fmt.Println("Time for Asr")
				Asr()
			}
		}

		//Checks if its time for Magrib
		if Y.Data.Timings.M == CurrentTime {
			if config.Prayers.Magrib == true {
				//fmt.Println("Time for Magrib")
				Magrib()

			}
		}

		//Checks if time for Isha
		if Y.Data.Timings.I == CurrentTime {
			if config.Prayers.Isha == true {
				//fmt.Println("Time for Isha")
				Isha()
				ACal() //Recall Json Data
			}

		}
	} // End Loop

}
