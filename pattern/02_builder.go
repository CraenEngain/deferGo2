package pattern

/*
	Реализовать паттерн «строитель».
	Преимущества:
		Позволяет изменять внутреннее представление продукта.
		Позволяет пошагово конструировать объект.
		Поддерживает создание различных представлений объекта с помощью одного и того же процесса конструирования.
	Недостатки:
		Усложняет код за счет введения дополнительных классов.
	Использование:
		Создание сложных документов (например, HTML или PDF).
		Настройка сложных объектов конфигурации.
		Конфигурация игровых персонажей и объектов.
		Создание сложных UI-компонентов.
		Создание сложных объектов в тестировании.

*/
import (
	"net/http"
	"strings"
)

// Request представляет HTTP-запрос.
type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
}

func (r *Request) Send() (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, r.URL, strings.NewReader(r.Body))
	if err != nil {
		return nil, err
	}
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}
	return client.Do(req)
}

// RequestBuilder предоставляет интерфейс для построения HTTP-запроса.
type RequestBuilder struct {
	request *Request
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		request: &Request{
			Headers: make(map[string]string),
		},
	}
}

func (b *RequestBuilder) SetMethod(method string) *RequestBuilder {
	b.request.Method = method
	return b
}

func (b *RequestBuilder) SetURL(url string) *RequestBuilder {
	b.request.URL = url
	return b
}

func (b *RequestBuilder) AddHeader(key, value string) *RequestBuilder {
	b.request.Headers[key] = value
	return b
}

func (b *RequestBuilder) SetBody(body string) *RequestBuilder {
	b.request.Body = body
	return b
}

func (b *RequestBuilder) Build() *Request {
	return b.request
}
