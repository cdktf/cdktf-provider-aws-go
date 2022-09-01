package codedeploy


type CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanary struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config#interval CodedeployDeploymentConfig#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_config#percentage CodedeployDeploymentConfig#percentage}.
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

