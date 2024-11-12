/* =========================================================================
*  File Name: utils/general.go
*  Description: A file containing common functions for all modules.
*  Author: MagnusChase03
*  =======================================================================*/
package utils

import (
	"net/http"
	"os"
	"strings"
)

/*
*  Retuns a map of environment variables
*
*  Arguments:
*      - N/A
*
*  Returns:
*      - map[string]string: The mapping of environment variables
 */
func GetEnvironment() map[string]string {
	result := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		result[pair[0]] = pair[1]
	}
	return result
}

/*
*  Retuns whether the requested websocket should be established
*
*  Arguments:
*      - r (*http.Request): The request that was made.
*
*  Returns:
*      - bool: Whether to allow or deny.
 */
func WebsocketOriginCheck(r *http.Request) bool {
	return true
}
