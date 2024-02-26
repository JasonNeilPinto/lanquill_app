package resp

import (
	"context"
	"net/http"

	"github.com/Lanquill/go-logger"
	"github.com/bitly/go-simplejson"
	"go.uber.org/zap"
)

func SendResponse(w http.ResponseWriter, httpCode int, jsonData *simplejson.Json, ctx *context.Context) {
	log := logger.FromCtx(*ctx)

	payload, err := jsonData.MarshalJSON()
	if err != nil {
		log.Error("error marshalling response: ", zap.Error(err), logger.LogUserId(*ctx))
		http.Error(w, InternalErrorMsg, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(payload)
}
