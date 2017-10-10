package storagecluster

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/zero-os/0-orchestrator/api/tools"
)

// KillCluster is the handler for DELETE /storageclusters/{label}
// Kill cluster
func (api *StorageclustersAPI) KillCluster(w http.ResponseWriter, r *http.Request) {
	aysClient, err := tools.GetAysConnection(api)
	if err != nil {
		tools.WriteError(w, http.StatusUnauthorized, err, "")
		return
	}
	vars := mux.Vars(r)
	storageCluster := vars["label"]

	// Prevent deletion of nonempty clusters
	query := map[string]interface{}{
		"consume": fmt.Sprintf("storage_cluster!%s", storageCluster),
	}
	services, res, err := aysClient.Ays.ListServicesByRole("vdisk", api.AysRepo, nil, query)
	if !tools.HandleAYSResponse(err, res, w, "listing vdisks") {
		return
	}

	if len(services) > 0 {
		err := fmt.Errorf("Can't delete storage clusters with attached vdisks")
		tools.WriteError(w, http.StatusBadRequest, err, "")
		return
	}

	// execute the delete action
	blueprint := map[string]interface{}{
		"actions": []tools.ActionBlock{{
			Action:  "delete",
			Actor:   "storage_cluster",
			Service: storageCluster,
		}},
	}

	run, err := aysClient.ExecuteBlueprint(api.AysRepo, "storage_cluster", storageCluster, "delete", blueprint)
	if err != nil {
		httpErr := err.(tools.HTTPError)
		tools.WriteError(w, httpErr.Resp.StatusCode, httpErr, "Error executing blueprint for storage_cluster deletion")
		return
	}

	// Wait for the delete job to be finshed before we delete the service
	if _, err = aysClient.WaitRunDone(run.Key, api.AysRepo); err != nil {
		httpErr, ok := err.(tools.HTTPError)
		if ok {
			tools.WriteError(w, httpErr.Resp.StatusCode, httpErr, "Error running blueprint for storage_cluster deletion")
		} else {
			tools.WriteError(w, http.StatusInternalServerError, err, "Error running blueprint for storage_cluster deletion")
		}
		return
	}

	res, err = aysClient.Ays.DeleteServiceByName(storageCluster, "storage_cluster", api.AysRepo, nil, nil)
	if !tools.HandleAYSDeleteResponse(err, res, w, "deleting storage_cluster") {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
