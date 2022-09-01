package appmesh


type AppmeshVirtualGatewaySpec struct {
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#listener AppmeshVirtualGateway#listener}
	Listener *AppmeshVirtualGatewaySpecListener `field:"required" json:"listener" yaml:"listener"`
	// backend_defaults block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#backend_defaults AppmeshVirtualGateway#backend_defaults}
	BackendDefaults *AppmeshVirtualGatewaySpecBackendDefaults `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#logging AppmeshVirtualGateway#logging}
	Logging *AppmeshVirtualGatewaySpecLogging `field:"optional" json:"logging" yaml:"logging"`
}

