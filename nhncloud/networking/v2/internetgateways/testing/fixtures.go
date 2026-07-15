package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

// Response fixtures below are the documented public-api.md examples. The Create
// request golden (CreateRequest) is synthesized from the required CreateOpts
// fields so the request-body assertion matches exactly what the client sends
// (guides often omit request examples); it is not a copy of a doc example.

const CreateRequest = `{"internetgateway":{"external_network_id":"sample","name":"sample"}}`

const CreateResponse = `{
  "internetgateway": {
    "create_time": "2025-02-11 00:44:18",
    "external_network_id": "50687905-7b9d-423a-b929-ab8b296a7f35",
    "id": "71f8e19a-8af3-4d3f-9b0e-b45a547fd7a7",
    "migrate_status": "none",
    "name": "ig-7ef985c1-8568",
    "state": "unavailable",
    "tenant_id": "130f20670ac34949b64b10ad8a5989c8"
  }
}`

// HandleCreateInternetGatewaySuccessfully sets up an HTTP handler that responds to
// a Create request with the documented response body.
func HandleCreateInternetGatewaySuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/internetgateways", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, CreateResponse)
	})
}

const GetResponse = `{
  "internetgateway": {
    "create_time": "2025-02-11 00:44:18",
    "external_network_id": "50687905-7b9d-423a-b929-ab8b296a7f35",
    "id": "71f8e19a-8af3-4d3f-9b0e-b45a547fd7a7",
    "migrate_error": null,
    "migrate_status": "none",
    "name": "ig-7ef985c1-8568",
    "routingtable_id": "7ef985c1-8568-4faa-a1ce-6a7116ba0e4d",
    "state": "available",
    "tenant_id": "130f20670ac34949b64b10ad8a5989c8"
  }
}`

// HandleGetInternetGatewaySuccessfully sets up an HTTP handler for a Get request.
func HandleGetInternetGatewaySuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/internetgateways/71f8e19a-8af3-4d3f-9b0e-b45a547fd7a7", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, GetResponse)
	})
}

const ListResponse = `{
  "internetgateways": [
    {
      "create_time": "2025-02-11 00:44:18",
      "external_network_id": "50687905-7b9d-423a-b929-ab8b296a7f35",
      "id": "71f8e19a-8af3-4d3f-9b0e-b45a547fd7a7",
      "migrate_error": null,
      "migrate_status": "none",
      "name": "ig-7ef985c1-8568",
      "routingtable_id": "7ef985c1-8568-4faa-a1ce-6a7116ba0e4d",
      "state": "available",
      "tenant_id": "130f20670ac34949b64b10ad8a5989c8"
    }
  ]
}`

// HandleListInternetGatewaySuccessfully sets up an HTTP handler for a List request.
func HandleListInternetGatewaySuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/internetgateways", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, ListResponse)
	})
}

// HandleDeleteInternetGatewaySuccessfully sets up an HTTP handler for a Delete request.
func HandleDeleteInternetGatewaySuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/internetgateways/71f8e19a-8af3-4d3f-9b0e-b45a547fd7a7", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})
}
