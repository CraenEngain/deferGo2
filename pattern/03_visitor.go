package pattern

/*
	Реализовать паттерн «посетитель».
	Преимущества:
		Упрощает добавление новых операций без изменения существующих классов.
		Позволяет объединять родственные операции, помещая их в один класс.
	Недостатки:
		Может затруднить добавление новых классов элементов, так как нужно будет обновлять всех посетителей.
		Часто требует раскрытия внутренней структуры классов.
	Использование:
		Компиляторы.
		Графические редакторы.
		Обработка документов.
		Инспекторы и анализаторы кода.
*/
import "fmt"

// Element интерфейс для элементов, которые будут принимать посетителей.
type Element interface {
	Accept(visitor Visitor)
}

// ConcreteElementA конкретный элемент A.
type ConcreteElementA struct {
	Name string
}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

func (e *ConcreteElementA) OperationA() string {
	return fmt.Sprintf("ElementA: %s", e.Name)
}

// ConcreteElementB конкретный элемент B.
type ConcreteElementB struct {
	ID int
}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

func (e *ConcreteElementB) OperationB() string {
	return fmt.Sprintf("Элемент: %d", e.ID)
}

// Visitor интерфейс для посетителей.
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// ConcreteVisitor1 конкретный посетитель 1.
type ConcreteVisitor1 struct{}

func (v *ConcreteVisitor1) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor1: Visiting", element.OperationA())
}

func (v *ConcreteVisitor1) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor1: Visiting", element.OperationB())
}

// ConcreteVisitor2 конкретный посетитель 2.
type ConcreteVisitor2 struct{}

func (v *ConcreteVisitor2) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor2: Visiting", element.OperationA())
}

func (v *ConcreteVisitor2) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor2: Visiting", element.OperationB())
}
