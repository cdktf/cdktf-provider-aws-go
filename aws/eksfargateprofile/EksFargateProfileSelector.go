package eksfargateprofile


type EksFargateProfileSelector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/eks_fargate_profile#namespace EksFargateProfile#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/eks_fargate_profile#labels EksFargateProfile#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

