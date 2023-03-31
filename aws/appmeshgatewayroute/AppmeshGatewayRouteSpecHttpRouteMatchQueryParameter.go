package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttpRouteMatchQueryParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#name AppmeshGatewayRoute#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#match AppmeshGatewayRoute#match}
	Match *AppmeshGatewayRouteSpecHttpRouteMatchQueryParameterMatch `field:"optional" json:"match" yaml:"match"`
}

