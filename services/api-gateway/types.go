package main

import (
	pb "ride-sharing/shared/proto/trip"
	"ride-sharing/shared/types"
)

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (p *previewTripRequest) ToProto() *pb.PreviewTripRequest {
	return &pb.PreviewTripRequest{
		UserID: p.UserID,
		StartLocation: &pb.Coordinate{
			Longitude: p.Pickup.Longitude,
			Latitude:  p.Pickup.Latitude,
		},

		EndLocation: &pb.Coordinate{
			Longitude: p.Destination.Longitude,
			Latitude:  p.Destination.Latitude,
		},
	}
}
