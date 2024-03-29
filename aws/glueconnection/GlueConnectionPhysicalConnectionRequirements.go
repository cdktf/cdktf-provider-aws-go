// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package glueconnection


type GlueConnectionPhysicalConnectionRequirements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/glue_connection#availability_zone GlueConnection#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/glue_connection#security_group_id_list GlueConnection#security_group_id_list}.
	SecurityGroupIdList *[]*string `field:"optional" json:"securityGroupIdList" yaml:"securityGroupIdList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/glue_connection#subnet_id GlueConnection#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

