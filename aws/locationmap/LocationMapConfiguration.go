// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package locationmap


type LocationMapConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/location_map#style LocationMap#style}.
	Style *string `field:"required" json:"style" yaml:"style"`
}

