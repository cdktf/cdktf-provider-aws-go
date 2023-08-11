package dataawskeypair


type DataAwsKeyPairFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/data-sources/key_pair#name DataAwsKeyPair#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/data-sources/key_pair#values DataAwsKeyPair#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

