// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace


type SagemakerSpaceSpaceSharingSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/sagemaker_space#sharing_type SagemakerSpace#sharing_type}.
	SharingType *string `field:"required" json:"sharingType" yaml:"sharingType"`
}

