package main

import (
	"fmt"
	"time"
)

//ConnectedTo gives information of connected google home and its basic paramaters
func ConnectedTo() {
	config, _ := LoadConfig("config.json")

	fmt.Println("Device Connected:", time.Now())
	fmt.Println("Connected to Device:", config.Settings.Name)
	fmt.Println("Using Lanuage:", config.Settings.Language)
	fmt.Println("Using Accent:", config.Settings.Accent)
	fmt.Println("Default Volume Set at", config.Volume.Default)

	//Calculation Method
	MethodV()

}

//MethodV Find out what Calculation method is being used
func MethodV() {
	config, _ := LoadConfig("config.json")

	switch config.Calculation.Method {
	case 0:
		//fmt.Println("Using Calculation Method: Shia Ithna-Ansari")
	case 1:
		fmt.Println("Using Calculation Method: University of Islamic Sciences, Karachi")
	case 2:
		fmt.Println("Using Calculation Method: Islamic Society of North America")
	case 3:
		fmt.Println("Using Calculation Method: Muslim World League")
	case 4:
		fmt.Println("Using Calculation Method: Umm Al-Qura University, Makkah")
	case 5:
		fmt.Println("Using Calculation Method: Egyptian General Authority of Survey")
	case 7:
		fmt.Println("Using Calculation Method: Institute of Geophysics, University of Tehran")
	case 8:
		fmt.Println("Using Calculation Method: Gulf Region")
	case 9:
		fmt.Println("Using Calculation Method: Kuwait")
	case 10:
		fmt.Println("Using Calculation Method: Qatar")
	case 11:
		fmt.Println("Using Calculation Method: Majlis Ugama Islam Singapura, Singapore")
	case 12:
		fmt.Println("Using Calculation Method: Union Organization islamic de France")
	case 13:
		fmt.Println("Using Calculation Method: Diyanet İşleri Başkanlığı, Turkey")
	default:
		fmt.Println("Other option choosen")
	}
}
