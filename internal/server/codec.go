package server

import (
	"github.com/go-kratos/kratos/v2/errors"
	kratoshttp "github.com/go-kratos/kratos/v2/transport/http"
	"net/http"
)

type HttpErrors map[string]Errors
type Errors map[string][]string

// HttpErrorEncoder encodes the error to the HTTP response.
func HttpErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	codec, _ := kratoshttp.CodecForRequest(r, "Accept")

	// Custom Response Error format
	errors := make(Errors)
	errors[se.Reason] = append(errors[se.Reason], se.Message)
	httpErrors := make(HttpErrors)
	httpErrors["errors"] = errors

	body, err := codec.Marshal(httpErrors)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(int(se.Code))
	_, _ = w.Write(body)
}
