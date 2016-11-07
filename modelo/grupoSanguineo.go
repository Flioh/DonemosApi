package modelo

type GrupoSanguineo struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}

func NewGrupoSanguineo(id int, nombre string) *GrupoSanguineo {
	return &GrupoSanguineo{id, nombre}
}
