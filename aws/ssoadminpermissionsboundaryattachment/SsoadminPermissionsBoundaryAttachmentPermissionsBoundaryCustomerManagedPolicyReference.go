package ssoadminpermissionsboundaryattachment


type SsoadminPermissionsBoundaryAttachmentPermissionsBoundaryCustomerManagedPolicyReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/ssoadmin_permissions_boundary_attachment#name SsoadminPermissionsBoundaryAttachment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/ssoadmin_permissions_boundary_attachment#path SsoadminPermissionsBoundaryAttachment#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

