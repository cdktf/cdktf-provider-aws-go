package gamelift


type GameliftFleetCertificateConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/gamelift_fleet#certificate_type GameliftFleet#certificate_type}.
	CertificateType *string `field:"optional" json:"certificateType" yaml:"certificateType"`
}

