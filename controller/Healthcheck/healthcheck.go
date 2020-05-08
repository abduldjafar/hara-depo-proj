package Healthcheck

import (
	"hara-depo-proj/util/customResponse"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	customResponse.RespondOk(w, 200, "Healthcheck")
}
