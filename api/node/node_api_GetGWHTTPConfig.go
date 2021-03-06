package node

import (
	"net/http"

	"bytes"
	"fmt"

	"github.com/gorilla/mux"
	client "github.com/zero-os/0-core/client/go-client"
	"github.com/zero-os/0-orchestrator/api/tools"
)

// GetGWHTTPConfig is the handler for GET /nodes/{nodeid}/gws/{gwname}/advanced/http
// Get current HTTP config
// Once used you can not use gw.httpproxxies any longer
func (api *NodeAPI) GetGWHTTPConfig(w http.ResponseWriter, r *http.Request) {
	var config bytes.Buffer

	vars := mux.Vars(r)
	gwname := vars["gwname"]

	node, err := tools.GetConnection(r, api)
	if err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err, "Failed to establish connection to node")
		return
	}
	containerID, err := tools.GetContainerId(r, api, node, gwname)
	if err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err, "Error getting NodeId")
		return
	}

	containerClient := client.Container(node).Client(containerID)
	err = client.Filesystem(containerClient).Download("/etc/caddy.conf", &config)
	if err != nil {
		errmsg := fmt.Sprintf("Error getting  file from container '%s' at path '%s'.\n", gwname, "/etc/caddy.conf")
		tools.WriteError(w, http.StatusInternalServerError, err, errmsg)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(config.Bytes())
}
