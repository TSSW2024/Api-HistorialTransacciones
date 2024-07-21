package models

type Monedero struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	UsuarioID string   `json:"usuarioID"` // Clave foránea a Usuario
	Monedas   []Moneda `json:"monedas"`   // Relación uno a muchos con Moneda
}

func (Monedero) TableName() string {
	return "monedero"
}
