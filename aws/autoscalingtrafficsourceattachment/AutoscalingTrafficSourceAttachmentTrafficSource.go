// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingtrafficsourceattachment


type AutoscalingTrafficSourceAttachmentTrafficSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/autoscaling_traffic_source_attachment#identifier AutoscalingTrafficSourceAttachment#identifier}.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/autoscaling_traffic_source_attachment#type AutoscalingTrafficSourceAttachment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

