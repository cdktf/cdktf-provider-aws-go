package opsworks


type OpsworksStackTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_stack#create OpsworksStack#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

