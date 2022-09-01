package vpc


type VpcIpamOperatingRegions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_ipam#region_name VpcIpam#region_name}.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
}

