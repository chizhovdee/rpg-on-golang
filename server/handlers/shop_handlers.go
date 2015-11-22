package handlers
import (
	"net/http"
	"github.com/chizhovdee/rpg/server/game_data/items"
	"encoding/json"
)

func SetupShopHandlers(mux *http.ServeMux){
	mux.HandleFunc("/shop", shopIndex)
}

func shopIndex(w http.ResponseWriter, r *http.Request){
	//items.SetupItems()

	js, err := json.Marshal(items.All().Select("item_2"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}