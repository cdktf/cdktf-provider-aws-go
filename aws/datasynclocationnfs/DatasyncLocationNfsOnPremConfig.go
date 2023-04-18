package datasynclocationnfs


type DatasyncLocationNfsOnPremConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/datasync_location_nfs#agent_arns DatasyncLocationNfs#agent_arns}.
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
}

