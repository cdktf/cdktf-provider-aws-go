package s3


type S3ControlMultiRegionAccessPointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_multi_region_access_point#create S3ControlMultiRegionAccessPoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_multi_region_access_point#delete S3ControlMultiRegionAccessPoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

