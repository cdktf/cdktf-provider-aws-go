// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paymentcryptographykey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaymentcryptographyKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/paymentcryptography_key#exportable PaymentcryptographyKey#exportable}.
	Exportable interface{} `field:"required" json:"exportable" yaml:"exportable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/paymentcryptography_key#deletion_window_in_days PaymentcryptographyKey#deletion_window_in_days}.
	DeletionWindowInDays *float64 `field:"optional" json:"deletionWindowInDays" yaml:"deletionWindowInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/paymentcryptography_key#enabled PaymentcryptographyKey#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// key_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/paymentcryptography_key#key_attributes PaymentcryptographyKey#key_attributes}
	KeyAttributes interface{} `field:"optional" json:"keyAttributes" yaml:"keyAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/paymentcryptography_key#key_check_value_algorithm PaymentcryptographyKey#key_check_value_algorithm}.
	KeyCheckValueAlgorithm *string `field:"optional" json:"keyCheckValueAlgorithm" yaml:"keyCheckValueAlgorithm"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/paymentcryptography_key#region PaymentcryptographyKey#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/paymentcryptography_key#tags PaymentcryptographyKey#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/paymentcryptography_key#timeouts PaymentcryptographyKey#timeouts}
	Timeouts *PaymentcryptographyKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

