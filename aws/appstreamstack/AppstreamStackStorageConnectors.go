// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appstreamstack


type AppstreamStackStorageConnectors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/appstream_stack#connector_type AppstreamStack#connector_type}.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/appstream_stack#domains AppstreamStack#domains}.
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/appstream_stack#resource_identifier AppstreamStack#resource_identifier}.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

