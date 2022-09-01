package sagemaker


type SagemakerModelVpcConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_model#security_group_ids SagemakerModel#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_model#subnets SagemakerModel#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

