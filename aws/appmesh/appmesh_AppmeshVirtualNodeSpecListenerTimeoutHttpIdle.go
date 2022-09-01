package appmesh


type AppmeshVirtualNodeSpecListenerTimeoutHttpIdle struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#unit AppmeshVirtualNode#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#value AppmeshVirtualNode#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

