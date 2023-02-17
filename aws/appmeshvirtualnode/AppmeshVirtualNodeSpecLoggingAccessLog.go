package appmeshvirtualnode


type AppmeshVirtualNodeSpecLoggingAccessLog struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#file AppmeshVirtualNode#file}
	File *AppmeshVirtualNodeSpecLoggingAccessLogFile `field:"optional" json:"file" yaml:"file"`
}

