package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttp2RouteActionRewrite struct {
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appmesh_gateway_route#hostname AppmeshGatewayRoute#hostname}
	Hostname *AppmeshGatewayRouteSpecHttp2RouteActionRewriteHostname `field:"optional" json:"hostname" yaml:"hostname"`
	// prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appmesh_gateway_route#prefix AppmeshGatewayRoute#prefix}
	Prefix *AppmeshGatewayRouteSpecHttp2RouteActionRewritePrefix `field:"optional" json:"prefix" yaml:"prefix"`
}

