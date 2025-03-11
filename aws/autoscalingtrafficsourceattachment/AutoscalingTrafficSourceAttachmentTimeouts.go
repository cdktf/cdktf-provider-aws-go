// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingtrafficsourceattachment


type AutoscalingTrafficSourceAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/autoscaling_traffic_source_attachment#create AutoscalingTrafficSourceAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/autoscaling_traffic_source_attachment#delete AutoscalingTrafficSourceAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

