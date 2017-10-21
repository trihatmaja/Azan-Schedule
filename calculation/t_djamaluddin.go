// this calculation based on Dr. T. Djamaluddin azan calculation
// Lembaga Penerbangan dan Antariksa Nasional Bandung
// t_djamal@lapan.go.id

package calculation

import (
	"fmt"
	"io/ioutil"
	"math"
)

const (
	Pi  = 3.14159
	Rad = Pi / 180.0
)

var (
	TheDate  = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	TheMonth = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
)

type Calculation struct {
	Latitude  float64
	Longitude float64
	Timezone  float64
	City      string
	T         [7]float64
}

type CalcResult struct {
	Month    string
	Schedule []AzanSchedule
}

type AzanSchedule struct {
	Date    int
	Fajr    string
	Sunrise string
	Zuhr    string
	Asr     string
	Maghrib string
	Isya    string
}

func New(latitude, longitude, timezone float64, city string) *Calculation {
	return &Calculation{
		Latitude:  latitude,
		Longitude: longitude,
		Timezone:  timezone,
		City:      city,
	}
}

func (az *Calculation) Calculate() []CalcResult {
	lamd := az.Longitude / 15.0
	phi := az.Latitude * Rad
	tdif := az.Timezone - lamd

	var retVal []CalcResult

	h := 0.0
	zd := 0.0
	n := 0.0
	for i := 0; i < 12; i++ {
		cr := CalcResult{}
		cr.Month = TheMonth[i]
		for k := 0; k < TheDate[i]; k++ {
			n = n + 1.0
			a := 6.0
			z := 110.0 * Rad
			for w := 1; w < 7; w++ {
				st := n + (a-lamd)/24.0
				L := (0.9856*st - 3.289) * Rad
				L = L + 1.916*Rad*math.Sin(L) + 0.02*Rad*math.Sin(2*L) + 282.634*Rad
				RA := float64(int(((L/Pi)*12.0)/6.0) + 1)
				if int(RA/2)*2-int(RA) != 0 {
					RA--
				}
				RA = (math.Atan(0.91746*math.Tan(L)) / Pi * 12.0) + float64(RA*6.0)
				X := 0.39782 * math.Sin(L)
				ATNX := math.Sqrt(1 - X*X)
				dek := math.Atan(X / ATNX)
				if a == 15 {
					z = math.Atan(math.Tan(zd) + 1)
				}
				X = (math.Cos(z) - X*math.Sin(phi)) / (ATNX * math.Cos(phi))
				if X <= 1.0 && X >= -1.0 {
					ATNX = math.Atan(math.Sqrt(1-X*X)/X) / Rad
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
						az.T[1] = st // t[1] = fajr
					}
					z = (90.0 + 5.0/6.0) * Rad
				case 2:
					az.T[2] = st // t[2] = sunrise
					a = 18.0
					z = (90.0 + 5.0/6.0) * Rad
				case 3:
					az.T[5] = st + 2.0/60.0 // t[5] = maghrib
					z = 108.0 * Rad
				case 4:
					if math.Abs(X) <= 1.0 {
						az.T[6] = st // t[6] = isya'
					}
					a = 12.0
				case 5:
					az.T[3] = st + 2.0/60.0 // t[3] = zuhr
					zd = math.Abs((dek - phi))
					a = 15.0
				case 6:
					az.T[4] = st // t[4] = asr
				}

				if n == 59.0 {
					if k == 27 {
						n--
					}
				}
			}
			as := AzanSchedule{}
			as.Date = k + 1
			for j := 1; j < 7; j++ {
				var buff string
				th := int32(az.T[j])                        // hour
				tm := int32((az.T[j] - float64(th)) * 60.0) // minute
				if tm < 10 {
					buff = fmt.Sprintf("%d:0%d", th, tm)
				} else {
					buff = fmt.Sprintf("%d:%d", th, tm)
				}
				switch j {
				case 1:
					as.Fajr = buff
					break
				case 2:
					as.Sunrise = buff
					break
				case 3:
					as.Zuhr = buff
					break
				case 4:
					as.Asr = buff
					break
				case 5:
					as.Maghrib = buff
					break
				case 6:
					as.Isya = buff
					break
				}
			}
			if int(n) == 59 {
				if k == 27 {
					n--
				}
			}
			cr.Schedule = append(cr.Schedule, as)
		}

		retVal = append(retVal, cr)
	}

	return retVal
}
