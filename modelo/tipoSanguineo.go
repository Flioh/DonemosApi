package modelo

type GrupoSanguineo int

const (
	Grupo0 GrupoSanguineo = iota
	GrupoA
	GrupoB
	GrupoAB
)

type FactorSanguineo int

const (
	RhPos FactorSanguineo = iota
	RhNeg
)

type TipoSanguineo struct {
	GrupoSanguineo  GrupoSanguineo  `json:"grupoSanguineo" bson:"grupoSanguineo"`
	FactorSanguineo FactorSanguineo `json:"factorSanguineo" bson:"factorSanguineo"`
}
