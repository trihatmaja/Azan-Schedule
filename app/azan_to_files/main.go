package main

import (
	"os"

	azan "github.com/trihatmaja/Azan-Schedule"

	"github.com/trihatmaja/Azan-Schedule/calculation"
	"github.com/trihatmaja/Azan-Schedule/database"

	"github.com/urfave/cli"
)

var (
	latitude  float64
	longitude float64
	timezone  float64
	city      string
)

func main() {
	opt := database.OptionFiles{
		OutputDir: ".",
		FileName:  "schedule.json",
	}

	db := database.NewFiles(opt)

	app := cli.NewApp()
	app.Name = "azan schedule"
	app.Usage = "generate files for azan schedule each day in a year"
	app.Commands = []cli.Command{
		cli.Command{
			Name:  "generate",
			Usage: "generate azan files",
			Action: func(c *cli.Context) error {
				calc := calculation.NewTDjamaluddin(latitude, longitude, timezone, city)
				az := azan.New(db, calc)
				az.Generate()
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
