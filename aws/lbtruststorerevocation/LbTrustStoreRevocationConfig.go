// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbtruststorerevocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbTrustStoreRevocationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/lb_trust_store_revocation#revocations_s3_bucket LbTrustStoreRevocation#revocations_s3_bucket}.
	RevocationsS3Bucket *string `field:"required" json:"revocationsS3Bucket" yaml:"revocationsS3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/lb_trust_store_revocation#revocations_s3_key LbTrustStoreRevocation#revocations_s3_key}.
	RevocationsS3Key *string `field:"required" json:"revocationsS3Key" yaml:"revocationsS3Key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/lb_trust_store_revocation#trust_store_arn LbTrustStoreRevocation#trust_store_arn}.
	TrustStoreArn *string `field:"required" json:"trustStoreArn" yaml:"trustStoreArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/lb_trust_store_revocation#id LbTrustStoreRevocation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/lb_trust_store_revocation#region LbTrustStoreRevocation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/lb_trust_store_revocation#revocations_s3_object_version LbTrustStoreRevocation#revocations_s3_object_version}.
	RevocationsS3ObjectVersion *string `field:"optional" json:"revocationsS3ObjectVersion" yaml:"revocationsS3ObjectVersion"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/lb_trust_store_revocation#timeouts LbTrustStoreRevocation#timeouts}
	Timeouts *LbTrustStoreRevocationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

