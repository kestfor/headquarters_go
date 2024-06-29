package geo

import (
	nominatim "github.com/doppiogancio/go-nominatim"
	"strings"
)

type Address struct {
	HouseNumber string
	Road        string
	City        string
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

func AddressFromString(s string) *Address {
	if len(s) == 0 {
		return &Address{}
	}
	split := strings.Split(s, ", ")
	return addressFromStringArray(split)
}

func NewAddress(houseNumber string, road string, city string) *Address {
	return &Address{HouseNumber: houseNumber, Road: road, City: city}
}

func AddressFromLocation(latitude float64, longitude float64) *Address {
	addr, err := nominatim.ReverseGeocode(latitude, longitude, "ru")
	if err != nil {
		return nil
	}
	return AddressFromString(addr.DisplayName)
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
