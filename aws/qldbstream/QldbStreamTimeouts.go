package qldbstream


type QldbStreamTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/qldb_stream#create QldbStream#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/qldb_stream#delete QldbStream#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

