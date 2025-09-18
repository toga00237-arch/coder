package coderd

import (
	"net/http"

	"github.com/coder/coder/v2/coderd/httpapi"
	"github.com/coder/coder/v2/coderd/rbac"
)

// listExternalScopes returns the curated list of API key scopes (resource:action)
// requestable via the API.
//
// @Summary List API key scopes
// @ID list-api-key-scopes
// @Tags Authorization
// @Produce json
// @Success 200 {array} string
// @Router /auth/scopes [get]
func (*API) listExternalScopes(rw http.ResponseWriter, r *http.Request) {
	httpapi.Write(r.Context(), rw, http.StatusOK, rbac.ExternalScopeNames())
}
