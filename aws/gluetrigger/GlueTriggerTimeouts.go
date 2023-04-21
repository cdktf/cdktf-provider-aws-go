package gluetrigger


type GlueTriggerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/glue_trigger#create GlueTrigger#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/glue_trigger#delete GlueTrigger#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

