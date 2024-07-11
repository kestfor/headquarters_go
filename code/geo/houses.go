package geo

const (
	HomeOfAlena = "Дом Алены"
	HomeOfIlya  = "Дом Ильи"
	HomeOfDima  = "Дом Димы"
	HomeOfAnton = "Академ"
)

type Home struct {
	Owner   string
	Address *Address
}

var Houses map[string]Home = map[string]Home{
	HomeOfIlya:  {"kestfor", NewAddress("15", "Телевизионная улица", "Новосибирск", 54.973853, 82.890825)},
	HomeOfAlena: {"alenochka_a_a", NewAddress("36/1", "Степная улица", "Новосибирск", 54.979519, 82.873449)},
	HomeOfDima:  {"Dadimka", NewAddress("49", "Улица Немировича-Данченко", "Новосибирск", 54.971897, 82.874717)},
	HomeOfAnton: {"adon_antonin", NewAddress("3", "Рубиновая улица", "Новосибирск", 54.867684, 83.082252)}}

var MainHome = Houses[HomeOfIlya]
