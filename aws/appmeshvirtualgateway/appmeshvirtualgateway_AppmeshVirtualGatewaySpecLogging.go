package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecLogging struct {
	// access_log block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#access_log AppmeshVirtualGateway#access_log}
	AccessLog *AppmeshVirtualGatewaySpecLoggingAccessLog `field:"optional" json:"accessLog" yaml:"accessLog"`
}

