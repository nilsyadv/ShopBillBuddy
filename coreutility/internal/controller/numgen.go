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
	NumSer *service.NumGenService
}

func NewNumGen(numgen *service.NumGenService, log logger.InterfaceLogger) *NumGenController {
	return &NumGenController{
		NumSer: numgen,
		Log:    log,
	}
}

// SetKey is an HTTP handler that handles requests to set up a new key in the Number Generator service.
// It expects the key, initial value, and length parameters in the URL path.
// If successful, it responds with a JSON-encoded success message.
// If an error occurs during the setup process, it logs the error and responds with an error message.
func (numgen NumGenController) SetKey(w http.ResponseWriter, req *http.Request) {
	// Extract key, initial value, and length from the URL parameters
	urlParams := mux.Vars(req)
	initialVal, _ := strconv.ParseInt(urlParams["intialvalue"], 10, 64)
	length, _ := strconv.ParseInt(urlParams["length"], 10, 64)

	// Create a NumberGeneratorRequest with the extracted parameters
	numReq := model.NumberGeneratorRequest{
		Prefix:     urlParams["key"],
		InitialVal: int(initialVal),
		Length:     int(length),
	}

	// Call the SetKey method of the NumGenService to set up the new key
	werr := numgen.NumSer.SetKey(numReq)
	if werr != nil {
		// Log the error and respond with an error message
		numgen.Log.Error(werr.Context, werr.Err)
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	// Respond with a JSON-encoded success message
	https.RespondJSON(w, http.StatusOK, map[string]string{"key": numReq.Prefix, "msg": "success"})
}

// NextKey is an HTTP handler that handles requests to generate the next key based on the provided key prefix.
// It expects the key parameter in the URL path.
// If successful, it responds with a JSON-encoded NumberGeneratorResponse containing the key and the generated value.
// If an error occurs during the generation process, it logs the error and responds with an error message.
func (numgen NumGenController) NextKey(w http.ResponseWriter, req *http.Request) {
	// Extract the key from the URL parameters
	urlParams := mux.Vars(req)
	key := urlParams["key"]

	// Create a NumberGeneratorRequest using the extracted key
	numReq := model.NumberGeneratorRequest{
		Prefix: key,
	}

	// Call the NextKey method of the NumGenService to get the next key value
	nextValue, werr := numgen.NumSer.NextKey(numReq)
	if werr != nil {
		// Log the error and respond with an error message
		numgen.Log.Error(werr.Context, werr.Err)
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	// Prepare the response with the key and the generated value
	resp := model.NumberGeneratorResponse{
		Key:   numReq.Prefix,
		Value: nextValue,
	}

	// Respond with the JSON-encoded response and HTTP status OK
	https.RespondJSON(w, http.StatusOK, resp)
}
