package workspacesipgroup


type WorkspacesIpGroupRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/workspaces_ip_group#source WorkspacesIpGroup#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/workspaces_ip_group#description WorkspacesIpGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

