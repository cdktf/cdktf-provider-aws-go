package ecrrepository


type EcrRepositoryImageScanningConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/ecr_repository#scan_on_push EcrRepository#scan_on_push}.
	ScanOnPush interface{} `field:"required" json:"scanOnPush" yaml:"scanOnPush"`
}

