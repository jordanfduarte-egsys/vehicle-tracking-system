package action

/**
* Package used to prepare the return data for the front
* @package action
* @author Jordan Duarte
**/

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
)

func PrintDebugf(format string, args ...interface{}) {
    if env := os.Getenv("GO_SERVER_DEBUG"); len(env) != 0 {
        log.Printf("[DEBUG] "+format+"\n", args...)
    }
}

type ReturnDynamic struct {
    Id int `json:"id"`
}

type ErrorResponse struct {
    Message string `json:"reason"`
    Error   error  `json:"-"`
}

type EmptyResponse struct {

}

type SuccessResponse struct {
    Message string
}

func Respond(w http.ResponseWriter, code int, src interface{}) {
    var body []byte
    var err error

    switch s := src.(type) {
    case []byte:
        if !json.Valid(s) {
            Error(w, http.StatusInternalServerError, err, "invalid json")
            return
        }
        body = s
    case string:
        str := fmt.Sprintf("%v", src)
        e := &SuccessResponse{Message: str}

        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(code)
        json.NewEncoder(w).Encode(e)
        return
    case *ErrorResponse, ErrorResponse:
        if body, err = json.Marshal(src); err != nil {
            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("{\"reason\":\"failed to parse json\"}"))
            return
        }
    case *EmptyResponse, EmptyResponse:
        body = []byte("")
    default:
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        if body, err = json.Marshal(src); err != nil {
            Error(w, http.StatusInternalServerError, err, "failed to parse json")
            return
        }
    }
    w.WriteHeader(code)
    w.Write(body)
}

func Error(w http.ResponseWriter, code int, err error, msg string) {
    if err != nil {
        e := &ErrorResponse{
            Message: msg,
            Error:   err,
        }
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        Respond(w, code, e)
    } else {
        Respond(w, code, &EmptyResponse{})
    }
}

func JSON(w http.ResponseWriter, code int, src interface{}) {
    Respond(w, code, src)
}