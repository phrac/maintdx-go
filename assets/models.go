package assets

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Asset struct {
	gorm.Model
	Name         string
	Type         AssetType
	SerialNumber string
	Make         string
	ModelNum     string
	InstallDate  time.Time
	Properties   postgres.Hstore
}

type AssetType struct {
	gorm.Model
	Name       string
	Properties []AssetProperties
}

type AssetProperties struct {
	gorm.Model
	AssetType AssetType
}
