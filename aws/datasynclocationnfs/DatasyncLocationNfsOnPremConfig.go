package datasynclocationnfs


type DatasyncLocationNfsOnPremConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/datasync_location_nfs#agent_arns DatasyncLocationNfs#agent_arns}.
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
}

