package appmesh


type AppmeshVirtualNodeSpecListenerConnectionPool struct {
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#grpc AppmeshVirtualNode#grpc}
	Grpc *AppmeshVirtualNodeSpecListenerConnectionPoolGrpc `field:"optional" json:"grpc" yaml:"grpc"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#http AppmeshVirtualNode#http}
	Http *AppmeshVirtualNodeSpecListenerConnectionPoolHttp `field:"optional" json:"http" yaml:"http"`
	// http2 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#http2 AppmeshVirtualNode#http2}
	Http2 *AppmeshVirtualNodeSpecListenerConnectionPoolHttp2 `field:"optional" json:"http2" yaml:"http2"`
	// tcp block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#tcp AppmeshVirtualNode#tcp}
	Tcp *AppmeshVirtualNodeSpecListenerConnectionPoolTcp `field:"optional" json:"tcp" yaml:"tcp"`
}

