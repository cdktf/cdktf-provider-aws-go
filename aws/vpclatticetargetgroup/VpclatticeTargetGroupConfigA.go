// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpclatticetargetgroup


type VpclatticeTargetGroupConfigA struct {
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/vpclattice_target_group#health_check VpclatticeTargetGroup#health_check}
	HealthCheck *VpclatticeTargetGroupConfigHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/vpclattice_target_group#ip_address_type VpclatticeTargetGroup#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/vpclattice_target_group#lambda_event_structure_version VpclatticeTargetGroup#lambda_event_structure_version}.
	LambdaEventStructureVersion *string `field:"optional" json:"lambdaEventStructureVersion" yaml:"lambdaEventStructureVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/vpclattice_target_group#port VpclatticeTargetGroup#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/vpclattice_target_group#protocol VpclatticeTargetGroup#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/vpclattice_target_group#protocol_version VpclatticeTargetGroup#protocol_version}.
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/vpclattice_target_group#vpc_identifier VpclatticeTargetGroup#vpc_identifier}.
	VpcIdentifier *string `field:"optional" json:"vpcIdentifier" yaml:"vpcIdentifier"`
}

