// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package prometheusresourcepolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrometheusResourcePolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/prometheus_resource_policy#policy_document PrometheusResourcePolicy#policy_document}.
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/prometheus_resource_policy#workspace_id PrometheusResourcePolicy#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/prometheus_resource_policy#region PrometheusResourcePolicy#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/prometheus_resource_policy#revision_id PrometheusResourcePolicy#revision_id}.
	RevisionId *string `field:"optional" json:"revisionId" yaml:"revisionId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/prometheus_resource_policy#timeouts PrometheusResourcePolicy#timeouts}
	Timeouts *PrometheusResourcePolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

