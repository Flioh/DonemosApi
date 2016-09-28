package modelo

type FactorSanguineo struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}

func NewFactorSanguineo(id int, nombre string) *FactorSanguineo {
	return &FactorSanguineo{id, nombre}
}
