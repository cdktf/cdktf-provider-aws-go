package directoryservice


type DirectoryServiceRegionVpcSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/directory_service_region#subnet_ids DirectoryServiceRegion#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/directory_service_region#vpc_id DirectoryServiceRegion#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

