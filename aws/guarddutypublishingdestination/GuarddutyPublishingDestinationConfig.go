// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guarddutypublishingdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GuarddutyPublishingDestinationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_publishing_destination#destination_arn GuarddutyPublishingDestination#destination_arn}.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_publishing_destination#detector_id GuarddutyPublishingDestination#detector_id}.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_publishing_destination#kms_key_arn GuarddutyPublishingDestination#kms_key_arn}.
	KmsKeyArn *string `field:"required" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_publishing_destination#destination_type GuarddutyPublishingDestination#destination_type}.
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_publishing_destination#id GuarddutyPublishingDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/guardduty_publishing_destination#region GuarddutyPublishingDestination#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

