package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustSds struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#secret_name AppmeshVirtualNode#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

