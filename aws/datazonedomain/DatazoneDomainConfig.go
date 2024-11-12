// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datazonedomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatazoneDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/datazone_domain#domain_execution_role DatazoneDomain#domain_execution_role}.
	DomainExecutionRole *string `field:"required" json:"domainExecutionRole" yaml:"domainExecutionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/datazone_domain#name DatazoneDomain#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/datazone_domain#description DatazoneDomain#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/datazone_domain#kms_key_identifier DatazoneDomain#kms_key_identifier}.
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// single_sign_on block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/datazone_domain#single_sign_on DatazoneDomain#single_sign_on}
	SingleSignOn interface{} `field:"optional" json:"singleSignOn" yaml:"singleSignOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/datazone_domain#skip_deletion_check DatazoneDomain#skip_deletion_check}.
	SkipDeletionCheck interface{} `field:"optional" json:"skipDeletionCheck" yaml:"skipDeletionCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/datazone_domain#tags DatazoneDomain#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/datazone_domain#timeouts DatazoneDomain#timeouts}
	Timeouts *DatazoneDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

