package tests

import (
	"fmt"
	geo2 "headquarters/code/geo"
	"testing"
)

func TestReverseGeo(t *testing.T) {
	addr := geo2.AddressFromLocation(54.973927, 82.890809)
	if !geo2.MainHome.Equivalent(addr) {
		t.Fatalf("Main home is incorrect")
	}
}

func TestEqGeo(t *testing.T) {
	addr1 := geo2.Houses[geo2.HomeOfAlena]
	addr2 := geo2.Houses[geo2.HomeOfIlya]
	if addr1.Equivalent(addr2) {
		t.Fatalf("Houses are not equal")
	}
}

func TestToString(t *testing.T) {
	fmt.Println(geo2.Houses[geo2.HomeOfAlena].ToString())
}
