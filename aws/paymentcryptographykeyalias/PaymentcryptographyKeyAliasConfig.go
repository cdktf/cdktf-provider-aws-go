// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paymentcryptographykeyalias

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaymentcryptographyKeyAliasConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/paymentcryptography_key_alias#alias_name PaymentcryptographyKeyAlias#alias_name}.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/paymentcryptography_key_alias#key_arn PaymentcryptographyKeyAlias#key_arn}.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/paymentcryptography_key_alias#region PaymentcryptographyKeyAlias#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

