package db

import "github.com/flioh/DonemosApi/modelo"

type IColección interface {
	Todos() (interface{}, error)
	//CRUD
	Create(modelo.IModelo) error
	Read(idHex string) (modelo.IModelo, error)
	Update(idHex string, object modelo.IModelo) error
	Delete(idHex string) error
}
