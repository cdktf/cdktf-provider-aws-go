package ssmcontactsplan


type SsmcontactsPlanStage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/ssmcontacts_plan#duration_in_minutes SsmcontactsPlan#duration_in_minutes}.
	DurationInMinutes *float64 `field:"required" json:"durationInMinutes" yaml:"durationInMinutes"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/ssmcontacts_plan#target SsmcontactsPlan#target}
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}

