/*
 * Control plane - API Gatway
 * Responsible for handling auth, upload signing and access tokens
 */

package gateway

import (
	"net/http"

	v1 "omegaui.io/edgeflow/gateway/v1"
)

func Listen() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/auth", v1.HandleAuth)
	mux.HandleFunc("/v1/media/status", v1.HandleMediaStatus)
	mux.HandleFunc("/v1/media/upload", v1.HandleMediaUpload)
	return mux
}
