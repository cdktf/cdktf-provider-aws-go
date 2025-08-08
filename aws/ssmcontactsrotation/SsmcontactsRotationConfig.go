// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsrotation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsmcontactsRotationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ssmcontacts_rotation#contact_ids SsmcontactsRotation#contact_ids}.
	ContactIds *[]*string `field:"required" json:"contactIds" yaml:"contactIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ssmcontacts_rotation#name SsmcontactsRotation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ssmcontacts_rotation#time_zone_id SsmcontactsRotation#time_zone_id}.
	TimeZoneId *string `field:"required" json:"timeZoneId" yaml:"timeZoneId"`
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ssmcontacts_rotation#recurrence SsmcontactsRotation#recurrence}
	Recurrence interface{} `field:"optional" json:"recurrence" yaml:"recurrence"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ssmcontacts_rotation#region SsmcontactsRotation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ssmcontacts_rotation#start_time SsmcontactsRotation#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ssmcontacts_rotation#tags SsmcontactsRotation#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

