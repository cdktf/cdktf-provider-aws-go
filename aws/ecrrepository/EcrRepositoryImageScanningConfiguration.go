package ecrrepository


type EcrRepositoryImageScanningConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecr_repository#scan_on_push EcrRepository#scan_on_push}.
	ScanOnPush interface{} `field:"required" json:"scanOnPush" yaml:"scanOnPush"`
}

