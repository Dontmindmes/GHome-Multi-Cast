package main

import (
	"time"

	"github.com/evalphobia/google-home-client-go/googlehome"
)

func Friday() {
	homes := LookupHomeIP()

	for _, home := range homes {
		cli, err := googlehome.NewClientWithConfig(googlehome.Config{
			Hostname: home.Ip,
			Lang:     "en",
			Accent:   "GB",
		})
		if err != nil {
			panic(err)
		}

		cli.Notify("I will begin reciting Quran.")
		time.Sleep(5 * time.Second)
		cli.SetVolume(config.Volume.Default)
		cli.Play(config.Audio.Recite)
	}
}
