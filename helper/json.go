package helper

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/zanuardi/go-xyz-multifinance/logger"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	if err != nil {
		logger.Error(context.Background(), "ReadFromRequestBody", err)
	}

}

func WriteToResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		logger.Error(context.Background(), "WriteToResponseBody", err)
	}
}
