package efsaccesspoint


type EfsAccessPointRootDirectory struct {
	// creation_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/efs_access_point#creation_info EfsAccessPoint#creation_info}
	CreationInfo *EfsAccessPointRootDirectoryCreationInfo `field:"optional" json:"creationInfo" yaml:"creationInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/efs_access_point#path EfsAccessPoint#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

