package appmeshgatewayroute


type AppmeshGatewayRouteSpecGrpcRoute struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/appmesh_gateway_route#action AppmeshGatewayRoute#action}
	Action *AppmeshGatewayRouteSpecGrpcRouteAction `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/appmesh_gateway_route#match AppmeshGatewayRoute#match}
	Match *AppmeshGatewayRouteSpecGrpcRouteMatch `field:"required" json:"match" yaml:"match"`
}

