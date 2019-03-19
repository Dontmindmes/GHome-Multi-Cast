package main

import "github.com/evalphobia/google-home-client-go/googlehome"

func Fajr() {
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

		cli.SetVolume(config.Volume.Fajir)
		cli.Play(config.Audio.Athan)
	}
}
