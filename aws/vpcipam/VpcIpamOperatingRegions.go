package vpcipam


type VpcIpamOperatingRegions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/vpc_ipam#region_name VpcIpam#region_name}.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
}

