package models

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func NewProduct(name string, price float32) *Product {
	return &Product{Id: 1, Name: name, Price: price}
}

func (p *Product) GetName() string {
	return p.Name
}

func (p Product) GetId() int {
	return p.Id
}

func (p *Product) SetName(name string) {
	p.Name = name
}
