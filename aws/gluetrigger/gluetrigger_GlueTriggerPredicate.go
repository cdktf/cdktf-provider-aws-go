package gluetrigger


type GlueTriggerPredicate struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger#conditions GlueTrigger#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger#logical GlueTrigger#logical}.
	Logical *string `field:"optional" json:"logical" yaml:"logical"`
}

