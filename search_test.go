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
	assert.Equal(t, "https://pointshound.com/hotel/results/110021/?checkin%3D2017-11-21T00%3A00%3A00%2B00%3A00%26checkout%3D2017-11-22T00%3A00%3A00%2B00%3A00%26adults%3D1%26flow%3Dearn%26locale%3Den-US%26currency%3DUSD%26destination_id%3Da_6000581%26program_id%3D16", hotel.Link)

	hotel, err = Search(37.79288317011038, -122.40053291718955, 244.24774169921875, HotelCache)
	assert.Nil(t, err)
	assert.Equal(t, "Loews Regency San Francisco", hotel.Name)
	assert.Equal(t, "https://pointshound.com/hotel/results/108742/?checkin%3D2017-11-21T00%3A00%3A00%2B00%3A00%26checkout%3D2017-11-22T00%3A00%3A00%2B00%3A00%26adults%3D1%26flow%3Dearn%26locale%3Den-US%26currency%3DUSD%26destination_id%3Da_6000581%26program_id%3D16", hotel.Link)

	hotel, err = Search(37.79096485999155, -122.40051141823338, 148.4856719970703, HotelCache)
	assert.Nil(t, err)
	assert.Equal(t, "Palace Hotel, a Luxury Collection Hotel, San Francisco", hotel.Name)
	assert.Equal(t, "https://pointshound.com/hotel/results/141179/?checkin%3D2017-11-21T00%3A00%3A00%2B00%3A00%26checkout%3D2017-11-22T00%3A00%3A00%2B00%3A00%26adults%3D1%26flow%3Dearn%26locale%3Den-US%26currency%3DUSD%26destination_id%3Da_6000581%26program_id%3D16", hotel.Link)

	hotel, err = Search(37.79286925938579, -122.4003320700139, 43.22932052612305, HotelCache)
	assert.Nil(t, err)
	assert.Equal(t, "Le Meridien San Francisco", hotel.Name)
	assert.Equal(t, "https://pointshound.com/hotel/results/110021/?checkin%3D2017-11-21T00%3A00%3A00%2B00%3A00%26checkout%3D2017-11-22T00%3A00%3A00%2B00%3A00%26adults%3D1%26flow%3Dearn%26locale%3Den-US%26currency%3DUSD%26destination_id%3Da_6000581%26program_id%3D16", hotel.Link)
}
