// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package locationgeofencecollection


type LocationGeofenceCollectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/location_geofence_collection#create LocationGeofenceCollection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/location_geofence_collection#delete LocationGeofenceCollection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/location_geofence_collection#update LocationGeofenceCollection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

