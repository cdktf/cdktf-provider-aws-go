package codedeploy


type CodedeployDeploymentGroupEc2TagSet struct {
	// ec2_tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group#ec2_tag_filter CodedeployDeploymentGroup#ec2_tag_filter}
	Ec2TagFilter interface{} `field:"optional" json:"ec2TagFilter" yaml:"ec2TagFilter"`
}

