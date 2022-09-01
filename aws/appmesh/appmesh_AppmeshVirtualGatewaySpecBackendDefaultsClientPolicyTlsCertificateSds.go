package appmesh


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateSds struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#secret_name AppmeshVirtualGateway#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

