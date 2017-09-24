package healthcheck

//This file is auto-generated by go-raml
//Do not edit this file by hand since it will be overwritten during the next generation

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Healthcheck Interface is interface for /health root endpoint
type HealthCheckInterface interface {
	// ListNodesHealth is the handler for GET /health/nodes
	// List NodesHealth
	ListNodesHealth(http.ResponseWriter, *http.Request)
	// ListNodeHealth is the handler for GET /health/nodes/{nodeid}
	// List NodeHealth
	ListNodeHealth(http.ResponseWriter, *http.Request)
	// ListStorageClustersHealth is the handler for GET /health/storageclusters
	// List StorageClustersHealthCheck
	ListStorageClustersHealth(http.ResponseWriter, *http.Request)
	// ListStorageClusterHealth is the handler for GET /health/storageclusters/{storageclusterid}
	// List StorageClusterHealthCheck
	ListStorageClusterHealth(http.ResponseWriter, *http.Request)
}

// HealthcheckInterfaceRoutes is routing for /health root endpoint
func HealthChechInterfaceRoutes(r *mux.Router, i HealthCheckInterface, org string) {
	r.HandleFunc("/health/nodes", i.ListNodesHealth).Methods("GET")
	r.HandleFunc("/health/nodes/{nodeid}", i.ListNodeHealth).Methods("GET")
	r.HandleFunc("/health/storageclusters", i.ListStorageClustersHealth).Methods("GET")
	r.HandleFunc("/health/storageclusters/{storageclusterid}", i.ListStorageClusterHealth).Methods("GET")
}