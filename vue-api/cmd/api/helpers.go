package main

import (
	"encoding/json"
	"errors"
	"io"
	"maps"
	"net/http"
)

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // one megabyte
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{}) // make sure there is only one json file
	if err != io.EOF {
		/*
			The value io.EOF is just a constant, that says you've reached the end of the file. This is from the documentation for that constant:
			In this check, we just make sure that we have read the entire JSON that we've received. If we get an error that is io.EOF,
			then all is good -- there is nothing else to read. However, any other error means that something went wrong --
			we have more data to read, and we shouldn't have (and don't want) any more.*/
		return errors.New("body must have only a single json value")
	}

	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		maps.Copy(w.Header(), headers[0]) // Go 1.21+ feature that allows you to copy a map efficiently
		// for key, value := range headers[0] {	w.Header()[key] = value	}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	app.writeJSON(w, statusCode, payload)
}
