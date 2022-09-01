package appmesh


type AppmeshVirtualNodeSpecListenerConnectionPoolTcp struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#max_connections AppmeshVirtualNode#max_connections}.
	MaxConnections *float64 `field:"required" json:"maxConnections" yaml:"maxConnections"`
}

