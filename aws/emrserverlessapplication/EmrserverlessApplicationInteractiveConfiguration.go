// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication


type EmrserverlessApplicationInteractiveConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/emrserverless_application#livy_endpoint_enabled EmrserverlessApplication#livy_endpoint_enabled}.
	LivyEndpointEnabled interface{} `field:"optional" json:"livyEndpointEnabled" yaml:"livyEndpointEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/emrserverless_application#studio_enabled EmrserverlessApplication#studio_enabled}.
	StudioEnabled interface{} `field:"optional" json:"studioEnabled" yaml:"studioEnabled"`
}

