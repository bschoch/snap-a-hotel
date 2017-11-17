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

	hotel, err = Search(37.79288317011038, -122.40053291718955, 244.24774169921875, HotelCache)
	assert.Nil(t, err)
	assert.Equal(t, "Loews Regency San Francisco", hotel.Name)

	hotel, err = Search(37.79096485999155, -122.40051141823338, 148.4856719970703, HotelCache)
	assert.Nil(t, err)
	assert.Equal(t, "Loews Regency San Francisco", hotel.Name)
}
