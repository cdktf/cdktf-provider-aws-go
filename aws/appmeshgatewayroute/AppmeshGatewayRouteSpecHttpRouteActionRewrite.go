package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttpRouteActionRewrite struct {
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#hostname AppmeshGatewayRoute#hostname}
	Hostname *AppmeshGatewayRouteSpecHttpRouteActionRewriteHostname `field:"optional" json:"hostname" yaml:"hostname"`
	// prefix block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#prefix AppmeshGatewayRoute#prefix}
	Prefix *AppmeshGatewayRouteSpecHttpRouteActionRewritePrefix `field:"optional" json:"prefix" yaml:"prefix"`
}

