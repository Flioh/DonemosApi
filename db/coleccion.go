package db

type IColección interface {
	Todos() ([]interface{}, error)
	//CRUD
	Create(interface{}) error
	Read(idHex string) (interface{}, error)
	Update(idHex string, object interface{}) error
	Delete(idHex string) error
}
