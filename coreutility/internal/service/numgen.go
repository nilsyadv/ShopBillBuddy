package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/coreutility/internal/db"
	"github.com/nilsyadv/ShopBillBuddy/coreutility/internal/model"
)

type NumGenService struct {
	rds *db.RedisDB
	log logger.InterfaceLogger
}

func NewNumGenService(rds *db.RedisDB, log logger.InterfaceLogger) *NumGenService {
	return &NumGenService{
		rds: rds,
		log: log,
	}
}

// SetKey sets up a new key in the Number Generator service with the provided configuration.
// It takes a NumberGeneratorRequest and creates a NumberGenerator with the specified parameters.
// The NumberGenerator is then stored in the Redis database using the key from the request.
// If successful, it returns nil. If an error occurs during the setup process, it returns a WrappedError.
func (numGenSer *NumGenService) SetKey(req model.NumberGeneratorRequest) *wraperror.WrappedError {
	// Create a NumberGenerator with the parameters from the request
	numReq := model.NumberGenerator{
		Prefix:   req.Prefix,
		NxtValue: req.InitialVal,
		Length:   req.Length,
	}

	// Store the NumberGenerator in the Redis database with the specified key
	return numGenSer.rds.Set(req.Prefix, numReq)
}

// NextKey generates the next key in the Number Generator service based on the provided configuration.
// It takes a NumberGeneratorRequest, retrieves the current value from Redis, increments it,
// and updates the Redis database with the incremented value.
// The generated key is then formatted with leading zeros based on the specified length.
// If successful, it returns the formatted next key. If an error occurs during the process,
// it returns a WrappedError.
func (numgenser *NumGenService) NextKey(req model.NumberGeneratorRequest) (string, *wraperror.WrappedError) {
	// Retrieve the current value from Redis based on the provided prefix
	val, werr := numgenser.rds.Get(req.Prefix)
	if werr != nil {
		return "", werr
	}

	// Unmarshal the Redis response into a NumberGenerator struct
	var dbResp model.NumberGenerator
	err := json.Unmarshal([]byte(val), &dbResp)
	if err != nil {
		return "", wraperror.Wrap(err, "error during unmarshaling Redis response", "error", http.StatusInternalServerError)
	}

	// Retrieve the next value and increment the counter in the database
	nextValue := dbResp.NxtValue
	dbResp.NxtValue++

	// Update the database with the incremented value
	werr = numgenser.rds.Update(req.Prefix, dbResp)
	if werr != nil {
		return "", werr
	}

	// Format the next value with leading zeros based on the specified length
	formattedNextValue := fmt.Sprintf(fmt.Sprintf("%%0%dd", req.Length), nextValue)

	return formattedNextValue, nil
}
