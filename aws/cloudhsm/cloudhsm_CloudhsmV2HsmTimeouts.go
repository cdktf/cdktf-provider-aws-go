package cloudhsm


type CloudhsmV2HsmTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudhsm_v2_hsm#create CloudhsmV2Hsm#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudhsm_v2_hsm#delete CloudhsmV2Hsm#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

