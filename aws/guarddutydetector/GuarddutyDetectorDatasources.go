// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guarddutydetector


type GuarddutyDetectorDatasources struct {
	// kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/guardduty_detector#kubernetes GuarddutyDetector#kubernetes}
	Kubernetes *GuarddutyDetectorDatasourcesKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// malware_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/guardduty_detector#malware_protection GuarddutyDetector#malware_protection}
	MalwareProtection *GuarddutyDetectorDatasourcesMalwareProtection `field:"optional" json:"malwareProtection" yaml:"malwareProtection"`
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/guardduty_detector#s3_logs GuarddutyDetector#s3_logs}
	S3Logs *GuarddutyDetectorDatasourcesS3Logs `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

