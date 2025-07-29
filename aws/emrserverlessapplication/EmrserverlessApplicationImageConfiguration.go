// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication


type EmrserverlessApplicationImageConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/emrserverless_application#image_uri EmrserverlessApplication#image_uri}.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
}

