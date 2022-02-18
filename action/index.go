package action

/**
* Package to render the index endpoint
* @package action
* @author Jordan Duarte
**/

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

type IndexHandler struct {}

func NewIndexHandler() *IndexHandler {
   return &IndexHandler{}
}

func (bc IndexHandler) IndexAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    type ReturnDynamic struct {
        Message string `json:"message"`
    }
    returnDynamic := &ReturnDynamic{Message: "Api is running!"}
    JSON(w, http.StatusOK, returnDynamic)
}
