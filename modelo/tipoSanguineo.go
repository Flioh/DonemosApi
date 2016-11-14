package modelo

type GrupoSanguineo struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}

func NewGrupoSanguineo(id int, nombre string) *GrupoSanguineo {
	return &GrupoSanguineo{id, nombre}
}

type FactorSanguineo struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}

func NewFactorSanguineo(id int, nombre string) *FactorSanguineo {
	return &FactorSanguineo{id, nombre}
}

type TipoSanguineo struct {
	GrupoSanguineo  GrupoSanguineo  `json:"grupoSanguineo" bson:"grupoSanguineo"`
	FactorSanguineo FactorSanguineo `json:"factorSanguineo" bson:"factorSanguineo"`
}
