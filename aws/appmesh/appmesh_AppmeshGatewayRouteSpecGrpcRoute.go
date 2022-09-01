package appmesh


type AppmeshGatewayRouteSpecGrpcRoute struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#action AppmeshGatewayRoute#action}
	Action *AppmeshGatewayRouteSpecGrpcRouteAction `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#match AppmeshGatewayRoute#match}
	Match *AppmeshGatewayRouteSpecGrpcRouteMatch `field:"required" json:"match" yaml:"match"`
}

