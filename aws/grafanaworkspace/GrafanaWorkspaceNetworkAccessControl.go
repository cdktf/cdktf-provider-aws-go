package grafanaworkspace


type GrafanaWorkspaceNetworkAccessControl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#prefix_list_ids GrafanaWorkspace#prefix_list_ids}.
	PrefixListIds *[]*string `field:"required" json:"prefixListIds" yaml:"prefixListIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#vpce_ids GrafanaWorkspace#vpce_ids}.
	VpceIds *[]*string `field:"required" json:"vpceIds" yaml:"vpceIds"`
}

