package cloudtrail


type CloudtrailAdvancedEventSelector struct {
	// field_selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail#field_selector Cloudtrail#field_selector}
	FieldSelector interface{} `field:"required" json:"fieldSelector" yaml:"fieldSelector"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail#name Cloudtrail#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

