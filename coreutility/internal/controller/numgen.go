package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/https"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/coreutility/internal/model"
	"github.com/nilsyadv/ShopBillBuddy/coreutility/internal/service"
)

type NumGenController struct {
	Log    logger.InterfaceLogger
	NumSer service.NumGenService
}

func NewNumGen(numgen service.NumGenService, log logger.InterfaceLogger) *NumGenController {
	return &NumGenController{
		NumSer: numgen,
		Log:    log,
	}
}

func (numgen NumGenController) SetKey(w http.ResponseWriter, req *http.Request) {
	urlquery := mux.Vars(req)

	initialval, _ := strconv.ParseInt(urlquery["value"], 10, 64)
	length, _ := strconv.ParseInt(urlquery["length"], 10, 64)

	numreq := model.NumberGeneratorRequest{
		Prefix:     urlquery["key"],
		InitialVal: int(initialval),
		Length:     int(length),
	}

	werr := numgen.NumSer.SetKey(numreq)
	if werr != nil {
		numgen.Log.Error(werr.Context, werr.Err)
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	https.RespondJSON(w, http.StatusOK, map[string]string{"key": numreq.Prefix, "msg": "success"})
}

func (numgen NumGenController) GetKey(w http.ResponseWriter, req *http.Request) {
	urlquery := mux.Vars(req)

	numreq := model.NumberGeneratorRequest{
		Prefix: urlquery["key"],
	}

	nxtval, werr := numgen.NumSer.NextKey(numreq)
	if werr != nil {
		numgen.Log.Error(werr.Context, werr.Err)
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	resp := model.NumberGeneratorResponse{
		Key:   numreq.Prefix,
		Value: nxtval,
	}

	https.RespondJSON(w, http.StatusOK, resp)
}
