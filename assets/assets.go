package assets

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io/ioutil"
	"log"
	"maintdx/common"
	"net/http"
)

func GetAsset(w http.ResponseWriter, r *http.Request) {
	db := common.GetDB()
	id := chi.URLParam(r, "id")
	a := new(Asset)
	err := db.Model(a).Where("id=?", id).Select()
	if err != nil {
		log.Printf("error getting asset")
	}
	render.JSON(w, r, a)
}

func GetAllAssets(w http.ResponseWriter, r *http.Request) {
	var assets []Asset
	db := common.GetDB()
	err := db.Model(&assets).Order("id ASC").Limit(20).Select()
	if err != nil {
		log.Printf("Error getting assets:", err)
	}
	render.JSON(w, r, assets)
}

func CreateAsset(w http.ResponseWriter, r *http.Request) {
	var a Asset
	db := common.GetDB()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}

	err = json.Unmarshal(body, &a)
	db.Insert(&a)
	render.JSON(w, r, a)
}

func CreateAssetType(w http.ResponseWriter, r *http.Request) {
	var at AssetType
	db := common.GetDB()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}

	err = json.Unmarshal(body, &at)
	db.Insert(&at)
	render.JSON(w, r, at)
}
