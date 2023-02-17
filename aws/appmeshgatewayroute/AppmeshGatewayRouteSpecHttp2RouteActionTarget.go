package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttp2RouteActionTarget struct {
	// virtual_service block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_gateway_route#virtual_service AppmeshGatewayRoute#virtual_service}
	VirtualService *AppmeshGatewayRouteSpecHttp2RouteActionTargetVirtualService `field:"required" json:"virtualService" yaml:"virtualService"`
}

