package appmesh


type AppmeshVirtualNodeSpecServiceDiscoveryAwsCloudMap struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#namespace_name AppmeshVirtualNode#namespace_name}.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#service_name AppmeshVirtualNode#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#attributes AppmeshVirtualNode#attributes}.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
}

