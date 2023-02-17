package cloudfrontdistribution


type CloudfrontDistributionOriginGroupMember struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution#origin_id CloudfrontDistribution#origin_id}.
	OriginId *string `field:"required" json:"originId" yaml:"originId"`
}

