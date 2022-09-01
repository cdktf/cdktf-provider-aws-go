// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraIndexUserGroupResolutionConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#user_group_resolution_mode KendraIndex#user_group_resolution_mode}.
	UserGroupResolutionMode *string `field:"required" json:"userGroupResolutionMode" yaml:"userGroupResolutionMode"`
}

