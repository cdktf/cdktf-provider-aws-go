package appmeshgatewayroute


type AppmeshGatewayRouteSpecGrpcRouteAction struct {
	// target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#target AppmeshGatewayRoute#target}
	Target *AppmeshGatewayRouteSpecGrpcRouteActionTarget `field:"required" json:"target" yaml:"target"`
}

