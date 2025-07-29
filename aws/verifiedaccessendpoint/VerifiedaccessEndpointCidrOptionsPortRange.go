// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccessendpoint


type VerifiedaccessEndpointCidrOptionsPortRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/verifiedaccess_endpoint#from_port VerifiedaccessEndpoint#from_port}.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/verifiedaccess_endpoint#to_port VerifiedaccessEndpoint#to_port}.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

