package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListenerTlsCertificateAcm struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_virtual_gateway#certificate_arn AppmeshVirtualGateway#certificate_arn}.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
}

