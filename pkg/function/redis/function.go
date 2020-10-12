package food

import (
	"flaxxed/pkg/util"
	"flaxxed/pkg/food"
	"github.com/google/wire"
	"net/http"
)

var srv food.FoodService

func init() {

	srv = initService()
}
func initService() food.FoodService {
	wire.Build(food.FoodServiceProvidersSet)
	return srv
}

func SearchFood(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.Respond(w, nil, nil, http.StatusInternalServerError)
		return
	}

	qry := r.URL.Query().Get("query")
	if qry == "" {
		util.Respond(w, nil, nil, http.StatusBadRequest)
		return
	}

	menu, err := srv.GetByString("Name", qry)
	if err != nil {
		util.Respond(w, err, nil, http.StatusInternalServerError)
		return
	}
	util.Respond(w, nil, menu, http.StatusOK)

}
