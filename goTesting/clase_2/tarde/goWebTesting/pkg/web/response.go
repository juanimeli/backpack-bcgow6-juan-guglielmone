package web // Este paquete es para el manejo de respuestas. Sustituye el gin.H que enviamos en
// el segundo argumento de ctx.JSON en el handler. La validacion de la request la seguimos haciendo
// igual que antes en el handler.

import (
	"strconv"
)

type Response struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) Response {

	if code < 300 {
		return Response{strconv.FormatInt(int64(code), 10), data, ""}
	}
	return Response{strconv.FormatInt(int64(code), 10), nil, err}
}
