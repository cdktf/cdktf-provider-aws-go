package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecLoggingAccessLogFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/appmesh_virtual_gateway#path AppmeshVirtualGateway#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/appmesh_virtual_gateway#format AppmeshVirtualGateway#format}
	Format *AppmeshVirtualGatewaySpecLoggingAccessLogFileFormat `field:"optional" json:"format" yaml:"format"`
}

