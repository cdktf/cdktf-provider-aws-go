// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appstreamimagebuilder


type AppstreamImageBuilderAccessEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appstream_image_builder#endpoint_type AppstreamImageBuilder#endpoint_type}.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appstream_image_builder#vpce_id AppstreamImageBuilder#vpce_id}.
	VpceId *string `field:"optional" json:"vpceId" yaml:"vpceId"`
}

