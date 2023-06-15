package gluetrigger


type GlueTriggerPredicate struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/glue_trigger#conditions GlueTrigger#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/glue_trigger#logical GlueTrigger#logical}.
	Logical *string `field:"optional" json:"logical" yaml:"logical"`
}

