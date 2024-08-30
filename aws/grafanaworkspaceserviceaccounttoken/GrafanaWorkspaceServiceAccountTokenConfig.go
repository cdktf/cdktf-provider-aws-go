// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grafanaworkspaceserviceaccounttoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrafanaWorkspaceServiceAccountTokenConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/grafana_workspace_service_account_token#name GrafanaWorkspaceServiceAccountToken#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/grafana_workspace_service_account_token#seconds_to_live GrafanaWorkspaceServiceAccountToken#seconds_to_live}.
	SecondsToLive *float64 `field:"required" json:"secondsToLive" yaml:"secondsToLive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/grafana_workspace_service_account_token#service_account_id GrafanaWorkspaceServiceAccountToken#service_account_id}.
	ServiceAccountId *string `field:"required" json:"serviceAccountId" yaml:"serviceAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/grafana_workspace_service_account_token#workspace_id GrafanaWorkspaceServiceAccountToken#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

