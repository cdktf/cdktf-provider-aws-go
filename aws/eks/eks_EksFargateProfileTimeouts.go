package eks


type EksFargateProfileTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eks_fargate_profile#create EksFargateProfile#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eks_fargate_profile#delete EksFargateProfile#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

