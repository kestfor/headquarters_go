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
	HomeOfIlya:  {"kestfor", NewAddress("15", "Телевизионная улица", "Новосибирск")},
	HomeOfAlena: {"alenochka_a_a", NewAddress("36/1", "Степная улица", "Новосибирск")},
	HomeOfDima:  {"Dadimka", NewAddress("49", "Улица Немировича-Данченко", "Новосибирск")},
	HomeOfAnton: {"adon_antonin", NewAddress("3", "Рубиновая улица", "Новосибирск")}}

var MainHome = Houses[HomeOfIlya]
