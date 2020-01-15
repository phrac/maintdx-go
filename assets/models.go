package assets

import (
	"time"
)

type Asset struct {
	Id           int
	Name         string
	Type         AssetType
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
