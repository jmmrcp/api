package store

type Product struct {
	ID     int
	Title  string
	Image  string
	Price  uint64
	Rating uint8
}

type Products []Product
