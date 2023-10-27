// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lightsailcontainerservice


type LightsailContainerServicePublicDomainNames struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/lightsail_container_service#certificate LightsailContainerService#certificate}
	Certificate interface{} `field:"required" json:"certificate" yaml:"certificate"`
}

