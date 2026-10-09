package main

import (
	"encoding/json"
	"log"
	"net/http"
	grpcclients "ride-sharing/services/api-gateway/grpc_clients"
	"ride-sharing/shared/contracts"
)

func handleTripStart(w http.ResponseWriter, r *http.Request) {
	var reqBody startTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse json data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if reqBody.UserID == "" {
		http.Error(w, "user id is required", http.StatusBadRequest)
	}

	tripService, err := grpcclients.NewTripServiceClient()
	if err != nil {
		log.Print(err)
		return
	}
	defer tripService.Close()

	trip, err := tripService.Client.CreateTrip(r.Context(), reqBody.ToProto())

	if err != nil {
		log.Printf("failed to start a trip: %v", err)
		http.Error(w, "failed to start trip", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{
		Data: trip,
	}

	writeJSON(w, http.StatusCreated, response)

}

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse json data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if reqBody.UserID == "" {
		http.Error(w, "user id is required", http.StatusBadRequest)
	}

	tripService, err := grpcclients.NewTripServiceClient()
	if err != nil {
		log.Print(err)
		return
	}
	defer tripService.Close()

	tripPreview, err := tripService.Client.PreviewTrip(r.Context(), reqBody.ToProto())

	if err != nil {
		log.Printf("failed to preview a trip: %v", err)
		http.Error(w, "failed to preview trip", http.StatusInternalServerError)
		return
	}
	response := contracts.APIResponse{Data: tripPreview}

	writeJSON(w, http.StatusCreated, response)
}
