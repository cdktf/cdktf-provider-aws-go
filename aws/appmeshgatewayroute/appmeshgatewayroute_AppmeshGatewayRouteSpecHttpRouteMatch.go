package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttpRouteMatch struct {
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#hostname AppmeshGatewayRoute#hostname}
	Hostname *AppmeshGatewayRouteSpecHttpRouteMatchHostname `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#prefix AppmeshGatewayRoute#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}
