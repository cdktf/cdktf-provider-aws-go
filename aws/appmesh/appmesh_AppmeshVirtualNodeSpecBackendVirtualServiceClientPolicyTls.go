package appmesh


type AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTls struct {
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#validation AppmeshVirtualNode#validation}
	Validation *AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidation `field:"required" json:"validation" yaml:"validation"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#certificate AppmeshVirtualNode#certificate}
	Certificate *AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#enforce AppmeshVirtualNode#enforce}.
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#ports AppmeshVirtualNode#ports}.
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
}

