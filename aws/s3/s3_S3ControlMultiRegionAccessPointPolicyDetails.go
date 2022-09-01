package s3


type S3ControlMultiRegionAccessPointPolicyDetails struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_multi_region_access_point_policy#name S3ControlMultiRegionAccessPointPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_multi_region_access_point_policy#policy S3ControlMultiRegionAccessPointPolicy#policy}.
	Policy *string `field:"required" json:"policy" yaml:"policy"`
}

