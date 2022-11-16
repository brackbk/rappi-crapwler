package domain

import (
	"time"
)

type TypeStore struct {
	Base
	Name       string        `json:"name" gorm:"type:varchar(255)"`
	TypeStore  string        `json:"typestore" gorm:"type:varchar(255)"`
	SubOptions []interface{} `gorm:"-"`
	Scan       string        `json:"list" gorm:"type:varchar(255)"`
}

func NewTypeStore(name string, typeStore string, scan string, SubOptions []interface{}) (*TypeStore, error) {

	typestore := &TypeStore{
		Name:       name,
		TypeStore:  typeStore,
		SubOptions: SubOptions,
		Scan:       scan,
	}

	//typestore.ID = uuid.NewV4().String()
	typestore.CreatedAt = time.Now()

	return typestore, nil
}
