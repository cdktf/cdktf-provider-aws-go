// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccessinstancetrustproviderattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VerifiedaccessInstanceTrustProviderAttachmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/verifiedaccess_instance_trust_provider_attachment#verifiedaccess_instance_id VerifiedaccessInstanceTrustProviderAttachment#verifiedaccess_instance_id}.
	VerifiedaccessInstanceId *string `field:"required" json:"verifiedaccessInstanceId" yaml:"verifiedaccessInstanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/verifiedaccess_instance_trust_provider_attachment#verifiedaccess_trust_provider_id VerifiedaccessInstanceTrustProviderAttachment#verifiedaccess_trust_provider_id}.
	VerifiedaccessTrustProviderId *string `field:"required" json:"verifiedaccessTrustProviderId" yaml:"verifiedaccessTrustProviderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/verifiedaccess_instance_trust_provider_attachment#id VerifiedaccessInstanceTrustProviderAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

