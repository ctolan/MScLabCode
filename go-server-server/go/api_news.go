/*
 * The EADes News
 *
 * This is the first TUDTCDC MSc in DevOps EADes microservices example
 *
 * API version: 1.0.0
 * Contact: jvkitt@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"net/http"
)

func GetAllNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}