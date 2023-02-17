package securityhubinsight


type SecurityhubInsightFiltersNetworkSourceIpv4 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/securityhub_insight#cidr SecurityhubInsight#cidr}.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
}

