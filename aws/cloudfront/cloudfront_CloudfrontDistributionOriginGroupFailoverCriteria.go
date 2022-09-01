package cloudfront


type CloudfrontDistributionOriginGroupFailoverCriteria struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution#status_codes CloudfrontDistribution#status_codes}.
	StatusCodes *[]*float64 `field:"required" json:"statusCodes" yaml:"statusCodes"`
}

