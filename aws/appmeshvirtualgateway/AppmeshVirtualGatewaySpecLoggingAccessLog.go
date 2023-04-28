package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecLoggingAccessLog struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appmesh_virtual_gateway#file AppmeshVirtualGateway#file}
	File *AppmeshVirtualGatewaySpecLoggingAccessLogFile `field:"optional" json:"file" yaml:"file"`
}

