package pattern

/*
	Реализовать паттерн «фасад».
	Преимущества:
		Упрощение использования сложных систем: Предоставляя простой интерфейс, фасад облегчает работу с подсистемой.
		Снижение зависимости: Фасад уменьшает количество зависимостей между клиентом и подсистемой.
		Легкость изменения подсистемы: Изменения в подсистеме не влияют на клиентский код, так как взаимодействие идет через фасад.
	Недостатки:
		Возможное ухудшение производительности: Дополнительный уровень абстракции может повлиять на производительность.
		Может скрыть слишком много функциональности: Иногда фасад может ограничить доступ к полезным функциям подсистемы.
	Использование:
	Системы управления базами данных (DBMS).
	Системы управления мультимедийными устройствами.
	Web Frameworks.
	Системы логирования.
*/

import "fmt"

// Подсистема 1: Телевизор
type TV struct{}

func (tv *TV) On() {
	fmt.Println("Включение ТВ.")
}

func (tv *TV) SetInputChannel(channel int) {
	fmt.Printf("Канал входа %d.\n", channel)
}

// Подсистема 2: Аудиосистема
type AudioSystem struct{}

func (audio *AudioSystem) On() {
	fmt.Println("Включение аудиосистемы.")
}

func (audio *AudioSystem) SetVolume(volume int) {
	fmt.Printf("Установка громкости %d.\n", volume)
}

// Подсистема 3: DVD-плеер
type DVDPlayer struct{}

func (dvd *DVDPlayer) On() {
	fmt.Println("Включение DVD-player.")
}

func (dvd *DVDPlayer) Play(movie string) {
	fmt.Printf("Воспроизведение фильма: %s.\n", movie)
}

// Фасад: Домашний кинотеатр
type HomeTheaterFacade struct {
	tv    *TV
	audio *AudioSystem
	dvd   *DVDPlayer
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		tv:    &TV{},
		audio: &AudioSystem{},
		dvd:   &DVDPlayer{},
	}
}

func (ht *HomeTheaterFacade) WatchMovie(movie string) {
	ht.tv.On()
	ht.tv.SetInputChannel(3)
	ht.audio.On()
	ht.audio.SetVolume(5)
	ht.dvd.On()
	ht.dvd.Play(movie)
}
