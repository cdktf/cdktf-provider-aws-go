package ec2


type LaunchTemplateElasticGpuSpecifications struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#type LaunchTemplate#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

