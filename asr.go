package main

import "github.com/evalphobia/google-home-client-go/googlehome"

func Asr() {
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

		cli.SetVolume(config.Volume.Asr)
		cli.Play(config.Audio.Athan)
	}
}
