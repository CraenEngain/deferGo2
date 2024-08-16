package pattern

/*
	Реализовать паттерн «фабричный метод».
	Преимущества:
		Позволяет подклассам изменять тип создаваемых объектов.
		Уменьшает зависимость между клиентским кодом и конкретными классами объектов.
		Обеспечивает гибкость и расширяемость системы.
	Недостатки:
		Может усложнить систему за счет увеличения числа классов.
	Использование:
		Создание пользовательских интерфейсов.
		Подключение к базам данных.
		Игровые объекты.
		Обработка документов.
*/

import "fmt"

// Transport интерфейс для транспорта.
type Transport interface {
	Drive()
}

// Car конкретный тип транспорта - автомобиль.
type Car struct {
	Brand string
}

func (c *Car) Drive() {
	fmt.Printf("Driving a %s car\n", c.Brand)
}

// Bike конкретный тип транспорта - велосипед.
type Bike struct {
	Brand string
}

func (b *Bike) Drive() {
	fmt.Printf("Riding a %s bike\n", b.Brand)
}

// TransportFactory интерфейс фабрики транспорта.
type TransportFactory interface {
	CreateTransport() Transport
}

// CarFactory фабрика для создания автомобилей.
type CarFactory struct {
	Brand string
}

func (f *CarFactory) CreateTransport() Transport {
	return &Car{Brand: f.Brand}
}

// BikeFactory фабрика для создания велосипедов.
type BikeFactory struct {
	Brand string
}

func (f *BikeFactory) CreateTransport() Transport {
	return &Bike{Brand: f.Brand}
}
