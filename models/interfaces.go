package models

// Model defines the operations a model can have
type Model interface {
	Insert() error
	//  Get(db database.Database) error
	//  GetById(db database.Database) error
	// 	Update(db database.Database) error
	// 	Delete(db database.Database) error
}
