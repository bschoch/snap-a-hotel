package main

import (
	"gopkg.in/stretchr/testify.v1/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	InitHotelCache()
	hotel, err := Search(37.79285169999994, -122.40037489777632, 347.6893005371094, HotelCache)
	assert.Nil(t, err)
	assert.Equal(t, "Le Meridien San Francisco", hotel.Name)
	assert.Equal(t, "https://pointshound.com/hotel/results/110021/?checkin=2017-11-21T00:00:00+00:00&checkout=2017-11-22T00:00:00+00:00&adults=1&flow=earn&locale=en-US&currency=USD&destination_id=a_6000581&program_id=16", hotel.Link)

	hotel, err = Search(37.79288317011038, -122.40053291718955, 244.24774169921875, HotelCache)
	assert.Nil(t, err)
	assert.Equal(t, "Loews Regency San Francisco", hotel.Name)
	assert.Equal(t, "https://pointshound.com/hotel/results/108742/?checkin=2017-11-21T00:00:00+00:00&checkout=2017-11-22T00:00:00+00:00&adults=1&flow=earn&locale=en-US&currency=USD&destination_id=a_6000581&program_id=16", hotel.Link)

	hotel, err = Search(37.79096485999155, -122.40051141823338, 148.4856719970703, HotelCache)
	assert.Nil(t, err)
	assert.Equal(t, "Palace Hotel, a Luxury Collection Hotel, San Francisco", hotel.Name)
	assert.Equal(t, "https://pointshound.com/hotel/results/141179/?checkin=2017-11-21T00:00:00+00:00&checkout=2017-11-22T00:00:00+00:00&adults=1&flow=earn&locale=en-US&currency=USD&destination_id=a_6000581&program_id=16", hotel.Link)

	hotel, err = Search(37.79286925938579, -122.4003320700139, 43.22932052612305, HotelCache)
	assert.Nil(t, err)
	assert.Equal(t, "Le Meridien San Francisco", hotel.Name)
	assert.Equal(t, "https://pointshound.com/hotel/results/110021/?checkin=2017-11-21T00:00:00+00:00&checkout=2017-11-22T00:00:00+00:00&adults=1&flow=earn&locale=en-US&currency=USD&destination_id=a_6000581&program_id=16", hotel.Link)
}
