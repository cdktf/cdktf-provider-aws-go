package appmeshvirtualnode


type AppmeshVirtualNodeSpecLoggingAccessLogFile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#path AppmeshVirtualNode#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// format block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#format AppmeshVirtualNode#format}
	Format *AppmeshVirtualNodeSpecLoggingAccessLogFileFormat `field:"optional" json:"format" yaml:"format"`
}

