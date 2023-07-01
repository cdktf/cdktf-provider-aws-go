package securityhubinsight


type SecurityhubInsightFiltersThreatIntelIndicatorType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

