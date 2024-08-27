package model

type Product struct {
	ID        int
	Name      string
	CreatedAt string
	UpdatedAt string
	Variants  []Variant
}
