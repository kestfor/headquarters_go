package geo

const (
	HomeOfAlena = "Дом Алены"
	HomeOfIlya  = "Дом Ильи"
	HomeOfDima  = "Дом Димы"
)

var Houses map[string]*Address = map[string]*Address{
	HomeOfIlya:  NewAddress("15", "Телевизионная улица", "Новосибирск"),
	HomeOfAlena: NewAddress("36/1", "Степная улица", "Новосибирск"),
	HomeOfDima:  NewAddress("49", "Улица Немировича-Данченко", "Новосибирск"),
}

var MainHome = Houses[HomeOfIlya]
