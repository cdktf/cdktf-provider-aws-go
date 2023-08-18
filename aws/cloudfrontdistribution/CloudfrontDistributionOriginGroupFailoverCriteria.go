package cloudfrontdistribution


type CloudfrontDistributionOriginGroupFailoverCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/cloudfront_distribution#status_codes CloudfrontDistribution#status_codes}.
	StatusCodes *[]*float64 `field:"required" json:"statusCodes" yaml:"statusCodes"`
}

