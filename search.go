package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kellydunn/golang-geo"
	"math"
	"sort"
	"strconv"
	"time"
)

var HotelCache []*Hotel

func InitHotelCache() error {
	if err := json.Unmarshal([]byte(hotelJSON), &HotelCache); err != nil {
		return err
	}
	for i := range HotelCache {
		var err error
		if HotelCache[i].Latitude, err = strconv.ParseFloat(HotelCache[i].LatitudeStr, 64); err != nil {
			return err
		}
		if HotelCache[i].Longitude, err = strconv.ParseFloat(HotelCache[i].LongitudeStr, 64); err != nil {
			return err
		}
	}
	return nil
}

var bearingThreshold float64 = 70
var distanceThreshold float64 = 0.3

type Hotel struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	LatitudeStr  string  `json:"lat"`
	LongitudeStr string  `json:"lng"`
	Link         string  `json:"link"`
	Latitude     float64 `json:"-"`
	Longitude    float64 `json:"-"`
}

func Search(latitude, longitude, bearing float64, hotels []*Hotel) (*Hotel, error) {
	o := geo.NewPoint(latitude, longitude)
	hotelSort := ByHotelDistance{
		Hotels:       hotels,
		UserLocation: o,
	}
	if bearing > 180 {
		bearing -= 360
	}
	sort.Sort(hotelSort)
	hotels = hotelSort.Hotels
	for i := range hotels {
		d := geo.NewPoint(hotels[i].Latitude, hotels[i].Longitude)
		dist := o.GreatCircleDistance(d)
		targ := o.BearingTo(d)
		if dist < distanceThreshold {
			if targ < 0 && bearing < 0 || targ > 0 && bearing > 0 {
				if math.Abs(targ-bearing) < bearingThreshold {
					return generateLink(hotels[i]), nil
				}
			} else {
				d := math.Abs(targ - bearing)
				if d > 180 {
					d = 360 - d
				}
				if d < bearingThreshold {
					return generateLink(hotels[i]), nil
				}
			}
		}
	}
	return nil, errors.New("No hotel found")
}

func generateLink(hotel *Hotel) *Hotel {
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	h := *hotel
	h.Link = fmt.Sprintf("https://pointshound.com/hotel/results/%d/?checkin=%s&checkout=%s&adults=1&flow=earn&locale=en-US&currency=USD&destination_id=a_6000581&program_id=16", hotel.ID, now.Format("2006-01-02T15:04:05-07:00"), now.Add(24*time.Hour).Format("2006-01-02T15:04:05-07:00"))
	return &h
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByHotelDistance struct {
	Hotels       []*Hotel
	UserLocation *geo.Point
}

func (a ByHotelDistance) Len() int      { return len(a.Hotels) }
func (a ByHotelDistance) Swap(i, j int) { a.Hotels[i], a.Hotels[j] = a.Hotels[j], a.Hotels[i] }
func (a ByHotelDistance) Less(i, j int) bool {
	return a.UserLocation.GreatCircleDistance(geo.NewPoint(a.Hotels[i].Latitude, a.Hotels[i].Longitude)) < a.UserLocation.GreatCircleDistance(geo.NewPoint(a.Hotels[j].Latitude, a.Hotels[j].Longitude))
}
