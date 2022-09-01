package ec2


type DataAwsAmiTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ami#read DataAwsAmi#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

