package ec2


type SpotInstanceRequestTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/spot_instance_request#create SpotInstanceRequest#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/spot_instance_request#delete SpotInstanceRequest#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

