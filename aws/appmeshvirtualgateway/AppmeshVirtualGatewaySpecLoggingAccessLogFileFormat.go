package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecLoggingAccessLogFileFormat struct {
	// json block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#json AppmeshVirtualGateway#json}
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#text AppmeshVirtualGateway#text}.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

