package appmeshvirtualnode


type AppmeshVirtualNodeSpecLogging struct {
	// access_log block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#access_log AppmeshVirtualNode#access_log}
	AccessLog *AppmeshVirtualNodeSpecLoggingAccessLog `field:"optional" json:"accessLog" yaml:"accessLog"`
}

