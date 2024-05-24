// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftdatashareauthorization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedshiftDataShareAuthorizationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/resources/redshift_data_share_authorization#consumer_identifier RedshiftDataShareAuthorization#consumer_identifier}.
	ConsumerIdentifier *string `field:"required" json:"consumerIdentifier" yaml:"consumerIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/resources/redshift_data_share_authorization#data_share_arn RedshiftDataShareAuthorization#data_share_arn}.
	DataShareArn *string `field:"required" json:"dataShareArn" yaml:"dataShareArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/resources/redshift_data_share_authorization#allow_writes RedshiftDataShareAuthorization#allow_writes}.
	AllowWrites interface{} `field:"optional" json:"allowWrites" yaml:"allowWrites"`
}

