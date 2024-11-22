// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package m2environment


type M2EnvironmentStorageConfiguration struct {
	// efs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/m2_environment#efs M2Environment#efs}
	Efs interface{} `field:"optional" json:"efs" yaml:"efs"`
	// fsx block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/m2_environment#fsx M2Environment#fsx}
	Fsx interface{} `field:"optional" json:"fsx" yaml:"fsx"`
}

