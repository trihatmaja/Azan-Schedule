package main

import (
	"fmt"
	"os"

	azan "github.com/trihatmaja/Azan-Schedule"

	"github.com/trihatmaja/Azan-Schedule/calculation"
	"github.com/trihatmaja/Azan-Schedule/database"

	"github.com/urfave/cli"
)

// apps var
var (
	latitude  float64
	longitude float64
	timezone  float64
	city      string
	outputdir string
	filename  string
)

// main var
var (
	Version string
	Build   string
)

func main() {
	app := cli.NewApp()
	app.HideVersion = true
	app.Name = "azan schedule"
	app.Usage = "generate files for azan schedule each day in a year"
	app.Commands = []cli.Command{
		cli.Command{
			Name:  "generate",
			Usage: "generate azan files",
			Action: func(c *cli.Context) error {
				opt := database.OptionFiles{
					OutputDir: outputdir,
					FileName:  filename,
				}

				db := database.NewFiles(opt)

				calc := calculation.NewTDjamaluddin()

				// in cli apps, no need cache mechanism
				// the cache mechanism is nil
				az := azan.New(db, nil, calc)
				az.Generate(latitude, longitude, timezone, city)
				return nil
			},
			Flags: []cli.Flag{
				cli.Float64Flag{
					Name:        "latitude",
					Destination: &latitude,
					Value:       -6.18,
				},
				cli.Float64Flag{
					Name:        "longitude",
					Destination: &longitude,
					Value:       106.83,
				},
				cli.Float64Flag{
					Name:        "timezone",
					Destination: &timezone,
					Value:       +7,
				},
				cli.StringFlag{
					Name:        "city",
					Destination: &city,
					Value:       "Jakarta",
				},
				cli.StringFlag{
					Name:        "outputdir",
					Destination: &outputdir,
					Value:       ".",
				},
				cli.StringFlag{
					Name:        "filename",
					Destination: &filename,
					Value:       "schedule.json",
				},
			},
		},
		cli.Command{
			Name:  "version",
			Usage: "azan schedule version",
			Action: func(c *cli.Context) error {
				fmt.Printf("{Version: %s, Build: %s}", Version, Build)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
