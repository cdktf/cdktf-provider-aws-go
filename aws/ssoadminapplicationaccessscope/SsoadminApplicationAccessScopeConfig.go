// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssoadminapplicationaccessscope

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsoadminApplicationAccessScopeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/ssoadmin_application_access_scope#application_arn SsoadminApplicationAccessScope#application_arn}.
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/ssoadmin_application_access_scope#scope SsoadminApplicationAccessScope#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/ssoadmin_application_access_scope#authorized_targets SsoadminApplicationAccessScope#authorized_targets}.
	AuthorizedTargets *[]*string `field:"optional" json:"authorizedTargets" yaml:"authorizedTargets"`
}

