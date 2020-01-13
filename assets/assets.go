package assets

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"maintdx/common"
	"net/http"
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
}

type AssetType struct {
	gorm.Model
	Name       string
	Properties postgres.Hstore
}

func GetAsset(w http.ResponseWriter, r *http.Request) {
	var a Asset
	db := common.GetDB()
	id := chi.URLParam(r, "id")
	db.First(&a, id)
	render.JSON(w, r, a)
}

func GetAllAssets(w http.ResponseWriter, r *http.Request) {
	var assets []Asset
	db := common.GetDB()
	db.Limit(10).Find(&assets)
	render.JSON(w, r, assets)
}
