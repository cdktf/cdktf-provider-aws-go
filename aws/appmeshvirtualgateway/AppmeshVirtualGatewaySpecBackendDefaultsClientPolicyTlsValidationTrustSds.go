package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSds struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/appmesh_virtual_gateway#secret_name AppmeshVirtualGateway#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

