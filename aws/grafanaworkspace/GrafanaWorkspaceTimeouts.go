// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grafanaworkspace


type GrafanaWorkspaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/grafana_workspace#create GrafanaWorkspace#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/grafana_workspace#update GrafanaWorkspace#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

