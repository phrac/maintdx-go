package assets

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
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
	Name      string
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

func CreateAsset(w http.ResponseWriter, r *http.Request) {
	var a Asset
	db := common.GetDB()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	err = json.Unmarshal(body, &a)
	db.Create(&a)
	render.JSON(w, r, a)
}
