package tests

import (
	"fmt"
	"headquarters/geo"
	"testing"
)

func TestReverseGeo(t *testing.T) {
	addr := geo.AddressFromLocation(54.973927, 82.890809)
	if !geo.MainHome.Equivalent(addr) {
		t.Fatalf("Main home is incorrect")
	}
}

func TestEqGeo(t *testing.T) {
	addr1 := geo.Houses[geo.HomeOfAlena]
	addr2 := geo.Houses[geo.HomeOfIlya]
	if addr1.Equivalent(addr2) {
		t.Fatalf("Houses are not equal")
	}
}

func TestToString(t *testing.T) {
	fmt.Println(geo.Houses[geo.HomeOfAlena].ToString())
}
