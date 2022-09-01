package securityhub


type SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange `field:"optional" json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#end SecurityhubInsight#end}.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#start SecurityhubInsight#start}.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

