package ec2


type DataAwsEipTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/eip#read DataAwsEip#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

