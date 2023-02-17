package cloudfrontdistribution


type CloudfrontDistributionOriginS3OriginConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution#origin_access_identity CloudfrontDistribution#origin_access_identity}.
	OriginAccessIdentity *string `field:"required" json:"originAccessIdentity" yaml:"originAccessIdentity"`
}

