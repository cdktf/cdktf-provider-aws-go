package ec2


type DataAwsKeyPairFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/key_pair#name DataAwsKeyPair#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/key_pair#values DataAwsKeyPair#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

