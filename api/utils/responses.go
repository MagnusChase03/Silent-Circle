/* =========================================================================
*  File Name: utils/responses.go
*  Description: A file containing common response functions for all modules.
*  Author: MagnusChase03
*  =======================================================================*/
package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Struct for consistant response format.
type JSONResponse struct {
	StatusCode int
	Data       interface{}
}

/*
*  Sends given data as response to a HTTP request.
*
*  Arguments:
*      - w (http.ResponseWriter): The object used to write a response to the client.
*      - data (interface{}): The data to send as a response.
*
*  Returns:
*      - error: The error if any occured.
 */
func SendResponse(w http.ResponseWriter, data JSONResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.StatusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("[ERROR] Failed to encode response into JSON. %w", err)
	}
	return nil
}

/*
*  Sends a response for an internal server error.
*
*  Arguments:
*      - w (http.ResponseWriter): The object used to write a response to the client.
*
*  Returns:
*      - N/A
 */
func SendInternalServerError(w http.ResponseWriter, e error) {
	fmt.Fprintf(os.Stderr, "[ERROR] %v\n", e)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(JSONResponse{
		StatusCode: 500,
		Data:       "Internal Server Error",
	})
}

/*
*  Sends a response for an unauthorized request.
*
*  Arguments:
*      - w (http.ResponseWriter): The object used to write a response to the client.
*
*  Returns:
*      - N/A
 */
func SendUnauthorizedRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(JSONResponse{
		StatusCode: 401,
		Data:       "Unauthorized",
	})
}

/*
*  Sends a response for a bad request.
*
*  Arguments:
*      - w (http.ResponseWriter): The object used to write a response to the client.
*
*  Returns:
*      - N/A
 */
func SendBadRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(JSONResponse{
		StatusCode: 400,
		Data:       "Bad Request",
	})
}
