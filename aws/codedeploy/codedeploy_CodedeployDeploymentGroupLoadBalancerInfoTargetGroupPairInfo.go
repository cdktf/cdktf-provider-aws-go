package codedeploy


type CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo struct {
	// prod_traffic_route block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group#prod_traffic_route CodedeployDeploymentGroup#prod_traffic_route}
	ProdTrafficRoute *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute `field:"required" json:"prodTrafficRoute" yaml:"prodTrafficRoute"`
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group#target_group CodedeployDeploymentGroup#target_group}
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// test_traffic_route block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group#test_traffic_route CodedeployDeploymentGroup#test_traffic_route}
	TestTrafficRoute *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute `field:"optional" json:"testTrafficRoute" yaml:"testTrafficRoute"`
}

