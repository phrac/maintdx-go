package assets

import (
	"time"
)

type Asset struct {
	Id           int
	CreateAt     time.Time `pg:"default:now()"`
	UpdatedAt    time.Time
	ParentId     int
	Parent       Parent `pg:"fk:parent_id"`
	DepartmentId int
	Department   Department `pg:"fk:department_id"`
	LocationId   int
	Location     Location `pg:"fk:location_id"`
	Parent       Asset    `pg:"fk:parent_id"`
	Name         string
	TypeId       int
	Type         AssetType `pg:"fk:type_id"`
	SerialNumber string
	Make         string
	ModelNum     string
	InstallDate  time.Time
}

type AssetType struct {
	Name       string
	Properties []AssetTypeProperties
}

type AssetTypeProperties struct {
	AssetType AssetType
	required  bool
	Name      string
}

type Department struct {
	Id   int
	Name string
}

type Location struct {
	Id   int
	Name string
}
