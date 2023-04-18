package kendraindex


type KendraIndexUserGroupResolutionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/kendra_index#user_group_resolution_mode KendraIndex#user_group_resolution_mode}.
	UserGroupResolutionMode *string `field:"required" json:"userGroupResolutionMode" yaml:"userGroupResolutionMode"`
}

