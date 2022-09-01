// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type DataAwsCeTagsTimePeriod struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ce_tags#end DataAwsCeTags#end}.
	End *string `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ce_tags#start DataAwsCeTags#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

