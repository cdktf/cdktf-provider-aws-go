package appmesh


type AppmeshGatewayRouteSpecHttp2Route struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#action AppmeshGatewayRoute#action}
	Action *AppmeshGatewayRouteSpecHttp2RouteAction `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#match AppmeshGatewayRoute#match}
	Match *AppmeshGatewayRouteSpecHttp2RouteMatch `field:"required" json:"match" yaml:"match"`
}

