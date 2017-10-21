/*
	this is implementation azan in the form of cli
*/

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
	outputdir string
	filename  string
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
				opt := database.OptionFiles{
					OutputDir: outputdir,
					FileName:  filename,
				}

				db := database.NewFiles(opt)

				calc := calculation.NewTDjamaluddin()

				az := azan.New(db, calc)
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
	}

	app.Run(os.Args)
}
