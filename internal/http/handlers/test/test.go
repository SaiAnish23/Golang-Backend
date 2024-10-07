package test

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/SaiAnish23/Golang-Backend/internal/types"
	"github.com/SaiAnish23/Golang-Backend/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return

		}

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		// slog.Info("student data", student)

		err1 := validator.New().Struct(student)
		slog.Error("error", err1)

		if err1 != nil {
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err1.(validator.ValidationErrors)))
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]string{"message": "student data received"})

		// w.Write([]byte("Welecome to the test api"))
	}
}
