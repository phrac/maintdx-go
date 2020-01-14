package assets

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io/ioutil"
	"maintdx/common"
	"net/http"
)

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
