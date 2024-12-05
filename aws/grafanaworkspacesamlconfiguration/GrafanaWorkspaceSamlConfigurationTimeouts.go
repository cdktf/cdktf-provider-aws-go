// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grafanaworkspacesamlconfiguration


type GrafanaWorkspaceSamlConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/grafana_workspace_saml_configuration#create GrafanaWorkspaceSamlConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/grafana_workspace_saml_configuration#delete GrafanaWorkspaceSamlConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

