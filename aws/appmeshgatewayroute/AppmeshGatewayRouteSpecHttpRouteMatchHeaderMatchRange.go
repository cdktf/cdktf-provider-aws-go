package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttpRouteMatchHeaderMatchRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#end AppmeshGatewayRoute#end}.
	End *float64 `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#start AppmeshGatewayRoute#start}.
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

