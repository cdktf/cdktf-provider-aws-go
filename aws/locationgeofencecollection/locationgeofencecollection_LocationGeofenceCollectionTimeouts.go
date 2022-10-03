package locationgeofencecollection


type LocationGeofenceCollectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/location_geofence_collection#create LocationGeofenceCollection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/location_geofence_collection#delete LocationGeofenceCollection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/location_geofence_collection#update LocationGeofenceCollection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

