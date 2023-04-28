package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerConnectionPoolTcp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appmesh_virtual_node#max_connections AppmeshVirtualNode#max_connections}.
	MaxConnections *float64 `field:"required" json:"maxConnections" yaml:"maxConnections"`
}

