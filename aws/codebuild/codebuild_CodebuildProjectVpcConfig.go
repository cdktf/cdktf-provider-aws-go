package codebuild


type CodebuildProjectVpcConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#security_group_ids CodebuildProject#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#subnets CodebuildProject#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#vpc_id CodebuildProject#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

