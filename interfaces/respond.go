package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// PrintDebugf behaves like log.Printf only in the debug env
func PrintDebugf(format string, args ...interface{}) {
	if env := os.Getenv("GO_SERVER_DEBUG"); len(env) != 0 {
		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}

// ErrorResponse is Error response template
type ErrorResponse struct {
	Message string `json:"reason"`
	Error   error  `json:"-"`
}

type EmptyResponse struct {

}

type SuccessResponse struct {
	Message string
}

// func (e *ErrorResponse) String() string {
// 	return fmt.Sprintf("reason: %s, error: %s", e.Message, e.Error.Error())
// }

// // Respond is response write to ResponseWriter
// func Respond(w http.ResponseWriter, code int, src interface{}) {
// 	var body []byte
// 	var err error

// 	switch s := src.(type) {
// 	case []byte:
// 		if !json.Valid(s) {
// 			Error(w, http.StatusInternalServerError, err, "invalid json")
// 			return
// 		}
// 		body = s
// 	case string:
// 		body = []byte(s)
// 		s := &SuccessResponse{
// 			Message: src
// 		}
// 	case *ErrorResponse, ErrorResponse:
// 		// avoid infinite loop
// 		if body, err = json.Marshal(src); err != nil {
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte("{\"reason\":\"failed to parse json\"}"))
// 			return
// 		}
// 	default:
// 		if body, err = json.Marshal(src); err != nil {
// 			Error(w, http.StatusInternalServerError, err, "failed to parse json")
// 			return
// 		}
// 	}
// 	w.WriteHeader(code)
// 	w.Write(body)
// }
// func Respond(w http.ResponseWriter, code int, response interface{}) {
// 	switch response.(type) {
// 		case *SuccessResponse, SuccessResponse:
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(code)
// 			json.NewEncoder(w).Encode(response)
// 		case *ErrorResponse, ErrorResponse:
// 			// PrintDebugf("%s", e.String())
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(code)
// 			json.NewEncoder(w).Encode(response)
// 		default:
// 			Error(w, http.StatusInternalServerError, nil, "failed to parse json")
// 	}
// }

// Respond is response write to ResponseWriter
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
		//var e interface
		str := fmt.Sprintf("%v", src)
		e := &SuccessResponse{Message: str}
		// if code == http.StatusOK
		// else e = &ErrorResponse{Message: src}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(e)
		return
	case *ErrorResponse, ErrorResponse:
		// avoid infinite loop
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

// Error is wrapped Respond when error response
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

	// // printDebugf("%s", e.String())
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    // w.WriteHeader(code)
    // json.NewEncoder(w).Encode(e)
}

// JSON is wrapped Respond when success response
func JSON(w http.ResponseWriter, code int, src interface{}) {
	Respond(w, code, src)

	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    // w.WriteHeader(code)
    // json.NewEncoder(w).Encode(e)
}