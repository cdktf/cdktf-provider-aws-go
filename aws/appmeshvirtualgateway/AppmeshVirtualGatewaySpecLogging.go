package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecLogging struct {
	// access_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/appmesh_virtual_gateway#access_log AppmeshVirtualGateway#access_log}
	AccessLog *AppmeshVirtualGatewaySpecLoggingAccessLog `field:"optional" json:"accessLog" yaml:"accessLog"`
}

