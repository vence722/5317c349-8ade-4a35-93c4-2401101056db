package models

// The order model
type Order struct {
	ID                   string
	OriginLatitude       float64
	OriginLongitude      float64
	DestinationLatitude  float64
	DestinationLongitude float64
	Distance             float64
	Status               string // enum of UNASSIGN, TAKEN
}
