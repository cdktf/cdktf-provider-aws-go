package securityhubinsight


type SecurityhubInsightFiltersProcessLaunchedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#unit SecurityhubInsight#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#value SecurityhubInsight#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

