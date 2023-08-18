package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerTlsCertificateAcm struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/appmesh_virtual_node#certificate_arn AppmeshVirtualNode#certificate_arn}.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
}

