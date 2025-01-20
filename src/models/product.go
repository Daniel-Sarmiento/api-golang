package models

type Product struct {
	id    int32   `json:"id"`
	name  string  `json:"name"`
	price float32 `json:"price"`
}

func NewProduct(name string, price float32) *Product {
	return &Product{id: 1, name: name, price: price}
}

func (p *Product) GetName() string {
	return p.name
}

func (p Product) GetId() int32 {
	return p.id
}

func (p *Product) SetName(name string) {
	p.name = name
}
