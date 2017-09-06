package main

import (
	"os"

	azan "github.com/trihatmaja/Azan-Schedule"

	"github.com/urfave/cli"
)

var (
	latitude  float64
	longitude float64
	timezone  float64
	city      string
)

func main() {
	app := cli.NewApp()
	app.Name = "azan schedule"
	app.Usage = "generate files for azan schedule each day in a year"
	app.Commands = []cli.Command{
		cli.Command{
			Name:  "generate",
			Usage: "generate azan files",
			Action: func(c *cli.Context) error {
				az := azan.New(latitude, longitude, timezone, city)
				az.Calculation()
				return nil
			},
			Flags: []cli.Flag{
				cli.Float64Flag{
					Name:        "latitude",
					Destination: &latitude,
				},
				cli.Float64Flag{
					Name:        "longitude",
					Destination: &longitude,
				},
				cli.Float64Flag{
					Name:        "timezone",
					Destination: &timezone,
				},
				cli.StringFlag{
					Name:        "city",
					Destination: &city,
				},
			},
		},
	}

	app.Run(os.Args)
}
