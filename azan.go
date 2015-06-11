package main

/*
# azan.go
# Anda boleh menggunakan dan menyebarkan file ini dengan menyebutkan sumbernya:
# Nara Sumber awal:
# Dr. T. Djamaluddin
# Lembaga Penerbangan dan Antariksa Nasional (LAPAN) Bandung
# Phone 022-6012602. Fax 022-6014998
# e-mail: t_djamal@lapan.go.id  t_djamal@hotmail.com
# Porting ke Perl:
# Wastono ST
# Jl Taman Cilandak Rt:001 Rw:04 No.4 Jakarta 12430
# Phone 021-75909268. was.tono@gmail.com
# Porting ke Golang:
# Wicaksono Trihatmaja
# trihatmaja@gmail.com
*/

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

import (
	"github.com/droundy/goopt"
)

var latitiude float64
var longitude float64
var timezone float64
var city string
var t [7]float64

var mydate = [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var mymonth = [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

const PI float64 = 3.14159
const rad float64 = PI / 180.0

var strbuf string

func init() {
	goopt.ReqArg([]string{"--latitude"}, "LAT", "Latitude Value",
		func(to string) error {
			var err error
			latitiude, err = strconv.ParseFloat(to, 64)
			if err != nil {
				return errors.New("value not an float")
			}
			return nil
		})
	goopt.ReqArg([]string{"--longitude"}, "LONG", "Longitude Value",
		func(to string) error {
			var err error
			longitude, err = strconv.ParseFloat(to, 64)
			if err != nil {
				return errors.New("value not an float")
			}
			return nil
		})

	goopt.ReqArg([]string{"--timezone"}, "TZ", "Time Zone Value",
		func(to string) error {
			var err error
			timezone, err = strconv.ParseFloat(to, 64)
			if err != nil {
				return errors.New("value not an float")
			}
			return nil
		})

	goopt.ReqArg([]string{"--city"}, "CT", "City",
		func(to string) error {
			city = strings.ToUpper(to)
			return nil
		})

	goopt.Author = "Wicaksono Trihatmaja <trihatmaja@gmail.com>"

	goopt.Version = "1.0"

	goopt.Suite = "Azan Schedule"

	goopt.Summary = "Azan Schedule based on location latitude and longitude"

	goopt.Description = func() string {
		return "Azan Schedule"
	}
}

func calculation() {
	strbuf += fmt.Sprintln("Jadwal Waktu Azan untuk wilayah", city)
	if timezone > 0 {
		strbuf += fmt.Sprintf("GMT+%v Latitude=%v Longitude=%v\n", timezone, latitiude, longitude)
	} else {
		strbuf += fmt.Sprintf("GMT%v Latitude=%v Longitude=%v\n", timezone, latitiude, longitude)
	}

	lamd := longitude / 15.0
	phi := latitiude * rad
	tdif := timezone - lamd

	h := 0.0
	zd := 0.0
	n := 0.0
	for i := 0; i < 12; i++ {
		strbuf += fmt.Sprintln("\n" + mymonth[i] + "\nTgl\tSubuh\tTerbit\tZuhur\tAshar\tMagrib\tIsya")
		for k := 0; k < mydate[i]; k++ {
			n = n + 1.0
			a := 6.0
			z := 110.0 * rad
			for w := 1; w < 7; w++ {
				st := n + (a-lamd)/24.0
				L := (0.9856*st - 3.289) * rad
				L = L + 1.916*rad*math.Sin(L) + 0.02*rad*math.Sin(2*L) + 282.634*rad
				RA := float64(int(((L/PI)*12.0)/6.0) + 1)
				if int(RA/2)*2-int(RA) != 0 {
					RA--
				}
				RA = (math.Atan(0.91746*math.Tan(L)) / PI * 12.0) + float64(RA*6.0)
				X := 0.39782 * math.Sin(L)
				ATNX := math.Sqrt(1 - X*X)
				dek := math.Atan(X / ATNX)
				if a == 15 {
					z = math.Atan(math.Tan(zd) + 1)
				}
				X = (math.Cos(z) - X*math.Sin(phi)) / (ATNX * math.Cos(phi))
				if X <= 1.0 && X >= -1.0 {
					ATNX = math.Atan(math.Sqrt(1-X*X)/X) / rad
					if ATNX < 0.0 {
						ATNX = ATNX + 180.0
					}
					h = (360.0 - ATNX) * 24.0 / 360.0
					if a == 18 {
						h = 24.0 - h
					}
					if a == 12 {
						h = 0.0
					}
				}
				if a == 15 {
					h = 24.0 - h
				}
				st = h + RA - 0.06571*st - 6.622 + 24.0
				st = st - float64(int(st/24.0)*24.0)
				st = st + tdif
				switch w {
				case 1:
					if math.Abs(X) <= 1.0 {
						t[1] = st // t[1] = subuh
					}
					z = (90.0 + 5.0/6.0) * rad
				case 2:
					t[2] = st // t[2] = sunrise
					a = 18.0
					z = (90.0 + 5.0/6.0) * rad
				case 3:
					t[5] = st + 2.0/60.0 // t[5] = maghrib
					z = 108.0 * rad
				case 4:
					if math.Abs(X) <= 1.0 {
						t[6] = st // t[6] = isya
					}
					a = 12.0
				case 5:
					t[3] = st + 2.0/60.0 // t[3] = dhuhur
					zd = math.Abs((dek - phi))
					a = 15.0
				case 6:
					t[4] = st // t[4] = ashar
				}

				if n == 59.0 {
					if k == 27 {
						n = n - 1.0
					}
				}
			}
			strbuf += fmt.Sprintf("%d\t", k+1)
			for j := 1; j < 7; j++ {
				th := int32(t[j])
				tm := int32((t[j] - float64(th)) * 60.0)
				if tm < 10 {
					strbuf += fmt.Sprintf("%d:0%d\t", th, tm)
				} else {
					strbuf += fmt.Sprintf("%d:%d\t", th, tm)
				}
				if j == 6 {
					strbuf += fmt.Sprintf("\n")
				}
			}
			if int(n) == 59 {
				if k == 27 {
					n--
				}
			}
		}
	}
	err := ioutil.WriteFile(city+".txt", []byte(strbuf), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	goopt.Parse(nil)
	calculation()
}
