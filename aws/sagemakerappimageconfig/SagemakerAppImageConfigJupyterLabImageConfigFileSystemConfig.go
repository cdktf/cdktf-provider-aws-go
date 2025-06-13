// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerappimageconfig


type SagemakerAppImageConfigJupyterLabImageConfigFileSystemConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/sagemaker_app_image_config#default_gid SagemakerAppImageConfig#default_gid}.
	DefaultGid *float64 `field:"optional" json:"defaultGid" yaml:"defaultGid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/sagemaker_app_image_config#default_uid SagemakerAppImageConfig#default_uid}.
	DefaultUid *float64 `field:"optional" json:"defaultUid" yaml:"defaultUid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/sagemaker_app_image_config#mount_path SagemakerAppImageConfig#mount_path}.
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

