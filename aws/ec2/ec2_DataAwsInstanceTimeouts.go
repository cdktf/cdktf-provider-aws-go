package ec2


type DataAwsInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/instance#read DataAwsInstance#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

