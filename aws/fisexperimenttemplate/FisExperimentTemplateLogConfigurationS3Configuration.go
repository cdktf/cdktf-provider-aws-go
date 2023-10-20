// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate


type FisExperimentTemplateLogConfigurationS3Configuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/fis_experiment_template#bucket_name FisExperimentTemplate#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/fis_experiment_template#prefix FisExperimentTemplate#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

