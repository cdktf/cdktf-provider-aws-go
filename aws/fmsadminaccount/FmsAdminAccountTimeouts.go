package fmsadminaccount


type FmsAdminAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/fms_admin_account#create FmsAdminAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/fms_admin_account#delete FmsAdminAccount#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

