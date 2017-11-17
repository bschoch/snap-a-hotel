package main

import (
	"encoding/json"
	"errors"
	"github.com/kellydunn/golang-geo"
	"sort"
	"strconv"
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

var bearingThreshold float64 = 60

type Hotel struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	LatitudeStr  string  `json:"lat"`
	LongitudeStr string  `json:"lng"`
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
		if dist < 0.5 {
			var matchLower, matchHigher bool
			lowerLimit := bearing - bearingThreshold
			if lowerLimit < -180 {
				if targ > lowerLimit + 360 {
					matchLower = true
				}
			} else if targ > lowerLimit {
				matchLower = true
			}
			upperLimit := bearing + bearingThreshold
			if upperLimit > 180 {
				if targ <  upperLimit - 360 {
					matchHigher = true
				}
			} else if targ < upperLimit {
				matchHigher = true
			}
			if matchLower && matchHigher {
				return hotels[i], nil
			}
		}
	}
	return nil, errors.New("No hotel found")
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
