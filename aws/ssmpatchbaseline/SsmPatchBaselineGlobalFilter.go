package ssmpatchbaseline


type SsmPatchBaselineGlobalFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/ssm_patch_baseline#key SsmPatchBaseline#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/ssm_patch_baseline#values SsmPatchBaseline#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

