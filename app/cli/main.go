//go:generate esc -o ../../player/mp3/mp3.go -pkg mp3 ../../player/mp3

/*
	this is implementation azan in the form of cli
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	azan "github.com/trihatmaja/Azan-Schedule"

	"github.com/trihatmaja/Azan-Schedule/calculation"
	"github.com/trihatmaja/Azan-Schedule/database"
	"github.com/trihatmaja/Azan-Schedule/player"
	"github.com/trihatmaja/Azan-Schedule/player/mp3"

	"github.com/jasonlvhit/gocron"
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

func checkSchedule(k *azan.CalcResult) {
	t := time.Now()
	tgl := t.Format("2006-January-2")
	jam := t.Format("15:04")

	for _, v := range k.Schedule {
		if tgl == v.Date {
			switch {
			case jam == v.Fajr:
				fmt.Println("Fajr Pray Time")
				player.Play(mp3.FSMustByte(false, "/player/mp3/fajr.mp3"))
			case jam == v.Zuhr:
				fmt.Println("Zuhr Pray Time")
				player.Play(mp3.FSMustByte(false, "/player/mp3/azan.mp3"))
			case jam == v.Asr:
				fmt.Println("Asr Pray Time")
				player.Play(mp3.FSMustByte(false, "/player/mp3/azan.mp3"))
			case jam == v.Maghrib:
				fmt.Println("Maghrib Pray Time")
				player.Play(mp3.FSMustByte(false, "/player/mp3/azan.mp3"))
			case jam == v.Isya:
				fmt.Println("Isya' Pray Time")
				player.Play(mp3.FSMustByte(false, "/player/mp3/azan.mp3"))
			default:
				break
			}
		}
	}
}

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
		cli.Command{
			Name:  "play",
			Usage: "play azan",
			Action: func(c *cli.Context) error {
				player.Play(mp3.FSMustByte(false, "/player/mp3/azan.mp3"))
				return nil
			},
		},
		cli.Command{
			Name:  "start",
			Usage: "starting praying schedule",
			Action: func(c *cli.Context) error {
				fmt.Println("Starting schedule...")

				//gocron.Every(1).Minute().Do(checkSchedule, &k)
				s := gocron.NewScheduler()

				var gracefulStop = make(chan os.Signal)
				signal.Notify(gracefulStop, syscall.SIGTERM)
				signal.Notify(gracefulStop, syscall.SIGINT)
				go func(s *gocron.Scheduler) {
					sig := <-gracefulStop
					fmt.Println("\nCaught sig:", sig)
					fmt.Println("Wait for apps to gracefully stop")
					s.Remove(checkSchedule)
					s.Clear()
					os.Exit(0)
				}(s)

				fmt.Println("Check file schedule.json...")

				if _, err := os.Stat("schedule.json"); os.IsNotExist(err) {
					fmt.Println("Cannot find file schedule.json, please run generate first!")
					return nil
				}

				fmt.Println("Ok!")

				f, err := ioutil.ReadFile("schedule.json")
				if err != nil {
					fmt.Println(err)
				}
				k := azan.CalcResult{}

				err = json.Unmarshal(f, &k)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println("Running job..")
				s.Every(1).Minute().Do(checkSchedule, &k)
				<-s.Start()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
