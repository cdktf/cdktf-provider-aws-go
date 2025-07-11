// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsCanvasAppSettingsWorkspaceSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/sagemaker_domain#s3_artifact_path SagemakerDomain#s3_artifact_path}.
	S3ArtifactPath *string `field:"optional" json:"s3ArtifactPath" yaml:"s3ArtifactPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/sagemaker_domain#s3_kms_key_id SagemakerDomain#s3_kms_key_id}.
	S3KmsKeyId *string `field:"optional" json:"s3KmsKeyId" yaml:"s3KmsKeyId"`
}

