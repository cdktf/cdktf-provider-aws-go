package securityhubinsight


type SecurityhubInsightFiltersConfidence struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#eq SecurityhubInsight#eq}.
	Eq *string `field:"optional" json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#gte SecurityhubInsight#gte}.
	Gte *string `field:"optional" json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#lte SecurityhubInsight#lte}.
	Lte *string `field:"optional" json:"lte" yaml:"lte"`
}

