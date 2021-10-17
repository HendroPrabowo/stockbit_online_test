package response

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func HttpResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	byteRespBody, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}
	w.Write(byteRespBody)
}
