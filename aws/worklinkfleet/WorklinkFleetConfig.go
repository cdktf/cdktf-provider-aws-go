// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package worklinkfleet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorklinkFleetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/worklink_fleet#name WorklinkFleet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/worklink_fleet#audit_stream_arn WorklinkFleet#audit_stream_arn}.
	AuditStreamArn *string `field:"optional" json:"auditStreamArn" yaml:"auditStreamArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/worklink_fleet#device_ca_certificate WorklinkFleet#device_ca_certificate}.
	DeviceCaCertificate *string `field:"optional" json:"deviceCaCertificate" yaml:"deviceCaCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/worklink_fleet#display_name WorklinkFleet#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/worklink_fleet#id WorklinkFleet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity_provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/worklink_fleet#identity_provider WorklinkFleet#identity_provider}
	IdentityProvider *WorklinkFleetIdentityProvider `field:"optional" json:"identityProvider" yaml:"identityProvider"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/worklink_fleet#network WorklinkFleet#network}
	Network *WorklinkFleetNetwork `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/worklink_fleet#optimize_for_end_user_location WorklinkFleet#optimize_for_end_user_location}.
	OptimizeForEndUserLocation interface{} `field:"optional" json:"optimizeForEndUserLocation" yaml:"optimizeForEndUserLocation"`
}

