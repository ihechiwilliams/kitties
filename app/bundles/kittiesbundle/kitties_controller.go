package kittiesbundle

import (
	"net/http"

	"github.com/charleyvibez/kitties/app/common"
)

// KittiesController struct
type KittiesController struct {
	common.Controller
}

// Index func return all kitties in database
func (c *KittiesController) Index(w http.ResponseWriter, r *http.Request) {
	c.SendJSON(
		w,
		r,
		[]*Kitty{NewKitty("Gaspart", "British", "2016-07-05")},
		http.StatusOK,
	)
}
