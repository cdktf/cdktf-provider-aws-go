// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appstreamfleet


type AppstreamFleetComputeCapacity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/appstream_fleet#desired_instances AppstreamFleet#desired_instances}.
	DesiredInstances *float64 `field:"optional" json:"desiredInstances" yaml:"desiredInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/appstream_fleet#desired_sessions AppstreamFleet#desired_sessions}.
	DesiredSessions *float64 `field:"optional" json:"desiredSessions" yaml:"desiredSessions"`
}

