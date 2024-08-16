package pattern

/*
	Реализовать паттерн «стратегия».
	Преимущества:
		Позволяет легко добавлять новые алгоритмы.
		Снижает связанность между клиентским кодом и конкретными реализациями алгоритмов.
		Облегчает тестирование различных алгоритмов.
	Недостатки:
		Увеличивает число классов.
		Усложняет код из-за необходимости создания дополнительных объектов для реализации различных алгоритмов.
	Использование:
		Алгоритмы сортировки и поиска.
		Маршрутизация и навигация.
		Форматирование данных.
		Игровые стратегии.
*/
import "fmt"

// SortStrategy интерфейс стратегии сортировки.
type SortStrategy interface {
	Sort([]int)
}

// BubbleSortStrategy стратегия сортировки пузырьком.
type BubbleSortStrategy struct{}

func (s *BubbleSortStrategy) Sort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("Использование сортировки пузырьком:", arr)
}

// QuickSortStrategy стратегия быстрой сортировки.
type QuickSortStrategy struct{}

func (s *QuickSortStrategy) Sort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Использование быстрой сортировки:", arr)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Context контекст, использующий стратегию сортировки.
type Context struct {
	strategy SortStrategy
}

func NewContext(strategy SortStrategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) SetStrategy(strategy SortStrategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(arr []int) {
	c.strategy.Sort(arr)
}
