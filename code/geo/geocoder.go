package geo

import (
	nominatim "github.com/doppiogancio/go-nominatim"
	"math"
	"strings"
)

type MapPoint interface {
	GetLatitude() float64
	GetLongitude() float64
}

type TelegramMapPoint struct {
	Longitude float64
	Latitude  float64
}

func (t *TelegramMapPoint) GetLatitude() float64 {
	return t.Latitude
}

func (t *TelegramMapPoint) GetLongitude() float64 {
	return t.Longitude
}

func NewTelegramMapPoint(longitude, latitude float64) *TelegramMapPoint {
	return &TelegramMapPoint{longitude, latitude}
}

type Address struct {
	HouseNumber string
	Road        string
	City        string
	Latitude    float64
	Longitude   float64
}

func (a *Address) GetLatitude() float64 {
	return a.Latitude
}

func (a *Address) GetLongitude() float64 {
	return a.Longitude
}

type AddressInterface interface {
	Equivalent(address *Address) bool
	ToString() string
}

func addressFromStringArray(arr []string) *Address {
	if len(arr) < 2 {
		return &Address{HouseNumber: arr[0]}
	} else if len(arr) < 4 {
		return &Address{HouseNumber: arr[0], Road: arr[1]}
	} else {
		return &Address{HouseNumber: arr[0], Road: arr[1], City: arr[3]}
	}
}

func addressFromString(s string) *Address {
	if len(s) == 0 {
		return &Address{}
	}
	split := strings.Split(s, ", ")
	return addressFromStringArray(split)
}

func NewAddress(houseNumber string, road string, city string, latitude float64, longitude float64) *Address {
	return &Address{HouseNumber: houseNumber, Road: road, City: city, Latitude: latitude, Longitude: longitude}
}

func AddressFromLocation(latitude float64, longitude float64) *Address {
	addr, err := nominatim.ReverseGeocode(latitude, longitude, "ru")
	if err != nil {
		return nil
	}
	res := addressFromString(addr.DisplayName)
	res.Latitude = latitude
	res.Longitude = longitude
	return res
}

func (addr *Address) Equivalent(address *Address) bool {
	if address == nil {
		return false
	}
	return addr.HouseNumber == address.HouseNumber && addr.Road == address.Road
}

func (addr *Address) ToString() string {
	if addr == nil {
		return ""
	}
	return addr.HouseNumber + ", " + addr.Road + ", " + addr.City
}

// Distance func returns distance between two points in meters
func Distance(first, second MapPoint) float64 {
	deltaFirst := degreesToRadians(first.GetLatitude()-second.GetLatitude()) / 2
	deltaSecond := degreesToRadians(first.GetLongitude()-second.GetLongitude()) / 2
	sinFirstDelta := math.Sin(deltaFirst)
	sinSecondDelta := math.Sin(deltaSecond)
	a := sinFirstDelta*sinFirstDelta + math.Cos(degreesToRadians(first.GetLatitude()))*math.Cos(degreesToRadians(second.GetLatitude()))*sinSecondDelta*sinSecondDelta
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
