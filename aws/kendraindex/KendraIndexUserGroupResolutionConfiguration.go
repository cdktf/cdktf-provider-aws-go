package kendraindex


type KendraIndexUserGroupResolutionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/kendra_index#user_group_resolution_mode KendraIndex#user_group_resolution_mode}.
	UserGroupResolutionMode *string `field:"required" json:"userGroupResolutionMode" yaml:"userGroupResolutionMode"`
}

