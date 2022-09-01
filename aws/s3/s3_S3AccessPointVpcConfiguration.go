package s3


type S3AccessPointVpcConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_access_point#vpc_id S3AccessPoint#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

