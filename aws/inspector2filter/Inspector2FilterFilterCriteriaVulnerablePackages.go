// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector2filter


type Inspector2FilterFilterCriteriaVulnerablePackages struct {
	// architecture block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/inspector2_filter#architecture Inspector2Filter#architecture}
	Architecture interface{} `field:"optional" json:"architecture" yaml:"architecture"`
	// epoch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/inspector2_filter#epoch Inspector2Filter#epoch}
	Epoch interface{} `field:"optional" json:"epoch" yaml:"epoch"`
	// file_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/inspector2_filter#file_path Inspector2Filter#file_path}
	FilePath interface{} `field:"optional" json:"filePath" yaml:"filePath"`
	// name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/inspector2_filter#name Inspector2Filter#name}
	Name interface{} `field:"optional" json:"name" yaml:"name"`
	// release block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/inspector2_filter#release Inspector2Filter#release}
	Release interface{} `field:"optional" json:"release" yaml:"release"`
	// source_lambda_layer_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/inspector2_filter#source_lambda_layer_arn Inspector2Filter#source_lambda_layer_arn}
	SourceLambdaLayerArn interface{} `field:"optional" json:"sourceLambdaLayerArn" yaml:"sourceLambdaLayerArn"`
	// source_layer_hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/inspector2_filter#source_layer_hash Inspector2Filter#source_layer_hash}
	SourceLayerHash interface{} `field:"optional" json:"sourceLayerHash" yaml:"sourceLayerHash"`
	// version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/inspector2_filter#version Inspector2Filter#version}
	Version interface{} `field:"optional" json:"version" yaml:"version"`
}

