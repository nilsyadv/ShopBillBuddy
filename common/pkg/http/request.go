package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
)

func RequestParse(r *http.Request, target interface{}) *wraperror.WrappedError {
	if r.Body == nil {
		err := wraperror.Wrap(errors.New(""), "Request Body is Empty", "error",
			http.StatusBadRequest)
		return err
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		err := wraperror.Wrap(err, "Failed to Read Request Body", "error",
			http.StatusBadRequest)
		return err
	}

	if len(body) == 0 {
		err := wraperror.Wrap(errors.New(""), "Request Body is Empty", "error",
			http.StatusBadRequest)
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		err := wraperror.Wrap(err, "Failed to Unmarshal Request Body", "error",
			http.StatusInternalServerError)
		return err
	}
	return nil
}
