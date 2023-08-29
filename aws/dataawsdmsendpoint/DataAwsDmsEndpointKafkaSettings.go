// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsdmsendpoint


type DataAwsDmsEndpointKafkaSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/data-sources/dms_endpoint#broker DataAwsDmsEndpoint#broker}.
	Broker *string `field:"required" json:"broker" yaml:"broker"`
}

