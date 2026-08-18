package httpapi

import (
	"net/http"
	"testing"
	"time"
)

func TestCompletedServiceReturnsVehicleToDispatchPool(t *testing.T) {
	h := newAPIHarness(t)
	created := h.createVehicle(t, "602")
	id := requireString(t, created, "id")
	opened := h.request(t, http.MethodPost, "/api/v1/maintenance", map[string]any{"vehicle_id": id, "kind": "scheduled", "notes": "service", "due_at": h.now.Add(48 * time.Hour)}, http.StatusCreated)
	maintenanceID := requireString(t, opened, "id")
	h.request(t, http.MethodPost, "/api/v1/maintenance/"+maintenanceID+"/start", nil, http.StatusOK)
	h.request(t, http.MethodPost, "/api/v1/maintenance/"+maintenanceID+"/complete", nil, http.StatusOK)
	page := h.request(t, http.MethodGet, "/api/v1/vehicles?limit=10&q=602", nil, http.StatusOK)
	item := page["items"].([]any)[0].(map[string]any)
	if item["status"] != "available" {
		t.Fatalf("vehicle status=%v", item["status"])
	}
}
