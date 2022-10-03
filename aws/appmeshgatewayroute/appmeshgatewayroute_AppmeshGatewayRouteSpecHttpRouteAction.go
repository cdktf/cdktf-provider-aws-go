package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttpRouteAction struct {
	// target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#target AppmeshGatewayRoute#target}
	Target *AppmeshGatewayRouteSpecHttpRouteActionTarget `field:"required" json:"target" yaml:"target"`
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#rewrite AppmeshGatewayRoute#rewrite}
	Rewrite *AppmeshGatewayRouteSpecHttpRouteActionRewrite `field:"optional" json:"rewrite" yaml:"rewrite"`
}

