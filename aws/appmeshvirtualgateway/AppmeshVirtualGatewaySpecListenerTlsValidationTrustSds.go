package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListenerTlsValidationTrustSds struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/appmesh_virtual_gateway#secret_name AppmeshVirtualGateway#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

