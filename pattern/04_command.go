package pattern

/*
	Реализовать паттерн «комманда».
	Преимущества:
		Упрощает добавление новых команд.
		Позволяет реализовать отмену и повтор операций.
		Позволяет легко создавать сложные макрокоманды (команды, состоящие из нескольких команд).
	Недостатки:
		Может усложнить код за счет введения большого количества классов команд.
	Использование:
		GUI-фреймворки.
		Игровые движки.
		Транзакционные системы.
*/
import "fmt"

// Command интерфейс для команд.
type Command interface {
	Execute()
	Undo()
}

// LightReceiver получатель команды для управления светом.
type LightReceiver struct {
	IsOn bool
}

func (l *LightReceiver) On() {
	l.IsOn = true
	fmt.Println("Свет включен")
}

func (l *LightReceiver) Off() {
	l.IsOn = false
	fmt.Println("Свет выключен")
}

// TVReceiver получатель команды для управления телевизором.
type TVReceiver struct {
	IsOn bool
}

func (tv *TVReceiver) On() {
	tv.IsOn = true
	fmt.Println("ТВ включен")
}

func (tv *TVReceiver) Off() {
	tv.IsOn = false
	fmt.Println("ТВ выключен")
}

// LightOnCommand команда для включения света.
type LightOnCommand struct {
	Light *LightReceiver
}

func (c *LightOnCommand) Execute() {
	c.Light.On()
}

func (c *LightOnCommand) Undo() {
	c.Light.Off()
}

// LightOffCommand команда для выключения света.
type LightOffCommand struct {
	Light *LightReceiver
}

func (c *LightOffCommand) Execute() {
	c.Light.Off()
}

func (c *LightOffCommand) Undo() {
	c.Light.On()
}

// TVOnCommand команда для включения телевизора.
type TVOnCommand struct {
	TV *TVReceiver
}

func (c *TVOnCommand) Execute() {
	c.TV.On()
}

func (c *TVOnCommand) Undo() {
	c.TV.Off()
}

// TVOffCommand команда для выключения телевизора.
type TVOffCommand struct {
	TV *TVReceiver
}

func (c *TVOffCommand) Execute() {
	c.TV.Off()
}

func (c *TVOffCommand) Undo() {
	c.TV.On()
}

// RemoteControlInvoker инициатор команд.
type RemoteControlInvoker struct {
	OnCommands  []Command
	OffCommands []Command
	UndoCommand Command
}

func NewRemoteControlInvoker() *RemoteControlInvoker {
	return &RemoteControlInvoker{
		OnCommands:  make([]Command, 2),
		OffCommands: make([]Command, 2),
	}
}

func (r *RemoteControlInvoker) SetCommand(slot int, onCommand, offCommand Command) {
	r.OnCommands[slot] = onCommand
	r.OffCommands[slot] = offCommand
}

func (r *RemoteControlInvoker) PressOnButton(slot int) {
	r.OnCommands[slot].Execute()
	r.UndoCommand = r.OnCommands[slot]
}

func (r *RemoteControlInvoker) PressOffButton(slot int) {
	r.OffCommands[slot].Execute()
	r.UndoCommand = r.OffCommands[slot]
}

func (r *RemoteControlInvoker) PressUndoButton() {
	r.UndoCommand.Undo()
}
