package pattern

/*
	Реализовать паттерн «состояние».
	Преимущества:
		Упрощает код, так как логика, связанная с состояниями, разделена на отдельные классы.
		Легко добавлять новые состояния и изменять существующие без изменения контекста.
		Уменьшает количество условных операторов в коде.
	Недостатки:
		Увеличивает количество классов.
		Сложнее понять и отладить, так как логика разбросана по разным классам.
	Использование:
		Автоматические банкоматы (ATM).
		Игровые персонажи.
		Процессоры платежей.
		Редакторы текста.
*/
import "fmt"

// VendingMachineState интерфейс состояния автомата.
type VendingMachineState interface {
	InsertCoin()
	PressButton()
	Dispense()
}

// VendingMachine контекст, управляющий состояниями.
type VendingMachine struct {
	state VendingMachineState
}

// NewVendingMachine создает новый автомат и устанавливает начальное состояние.
func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{}
	vm.setState(&NoCoinState{vendingMachine: vm})
	return vm
}

func (vm *VendingMachine) setState(state VendingMachineState) {
	vm.state = state
}

func (vm *VendingMachine) InsertCoin() {
	vm.state.InsertCoin()
}

func (vm *VendingMachine) PressButton() {
	vm.state.PressButton()
}

func (vm *VendingMachine) Dispense() {
	vm.state.Dispense()
}

// NoCoinState состояние, когда монета не вставлена.
type NoCoinState struct {
	vendingMachine *VendingMachine
}

func (s *NoCoinState) InsertCoin() {
	fmt.Println("Вставлена монета.")
	s.vendingMachine.setState(&HasCoinState{vendingMachine: s.vendingMachine})
}

func (s *NoCoinState) PressButton() {
	fmt.Println("Для начала вставьте монету.")
}

func (s *NoCoinState) Dispense() {
	fmt.Println("Для начала вставьте монету.")
}

// HasCoinState состояние, когда монета вставлена.
type HasCoinState struct {
	vendingMachine *VendingMachine
}

func (s *HasCoinState) InsertCoin() {
	fmt.Println("Монета уже вставлена.")
}

func (s *HasCoinState) PressButton() {
	fmt.Println("Кнопка нажата.")
	s.vendingMachine.setState(&DispensingState{vendingMachine: s.vendingMachine})
}

func (s *HasCoinState) Dispense() {
	fmt.Println("Нажмите кнопку для выдачи.")
}

// DispensingState состояние, когда происходит выдача товара.
type DispensingState struct {
	vendingMachine *VendingMachine
}

func (s *DispensingState) InsertCoin() {
	fmt.Println("Пожалуйста подождите, происходит выдача.")
}

func (s *DispensingState) PressButton() {
	fmt.Println("Уже выдана.")
}

func (s *DispensingState) Dispense() {
	fmt.Println("Выдача предмета.")
	s.vendingMachine.setState(&NoCoinState{vendingMachine: s.vendingMachine})
}
