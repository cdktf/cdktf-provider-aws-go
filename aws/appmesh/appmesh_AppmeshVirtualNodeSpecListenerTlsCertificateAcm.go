package appmesh


type AppmeshVirtualNodeSpecListenerTlsCertificateAcm struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_node#certificate_arn AppmeshVirtualNode#certificate_arn}.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
}

