package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecLoggingAccessLog struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#file AppmeshVirtualGateway#file}
	File *AppmeshVirtualGatewaySpecLoggingAccessLogFile `field:"optional" json:"file" yaml:"file"`
}

