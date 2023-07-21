package grafanaworkspace


type GrafanaWorkspaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/grafana_workspace#create GrafanaWorkspace#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/grafana_workspace#update GrafanaWorkspace#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

