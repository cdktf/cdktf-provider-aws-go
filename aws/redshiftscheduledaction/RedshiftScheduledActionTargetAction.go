// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftscheduledaction


type RedshiftScheduledActionTargetAction struct {
	// pause_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/redshift_scheduled_action#pause_cluster RedshiftScheduledAction#pause_cluster}
	PauseCluster *RedshiftScheduledActionTargetActionPauseCluster `field:"optional" json:"pauseCluster" yaml:"pauseCluster"`
	// resize_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/redshift_scheduled_action#resize_cluster RedshiftScheduledAction#resize_cluster}
	ResizeCluster *RedshiftScheduledActionTargetActionResizeCluster `field:"optional" json:"resizeCluster" yaml:"resizeCluster"`
	// resume_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/redshift_scheduled_action#resume_cluster RedshiftScheduledAction#resume_cluster}
	ResumeCluster *RedshiftScheduledActionTargetActionResumeCluster `field:"optional" json:"resumeCluster" yaml:"resumeCluster"`
}

