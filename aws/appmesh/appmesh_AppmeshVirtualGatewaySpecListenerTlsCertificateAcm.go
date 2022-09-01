package appmesh


type AppmeshVirtualGatewaySpecListenerTlsCertificateAcm struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway#certificate_arn AppmeshVirtualGateway#certificate_arn}.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
}

