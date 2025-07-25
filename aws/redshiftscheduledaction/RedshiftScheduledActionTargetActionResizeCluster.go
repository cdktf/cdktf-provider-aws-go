// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftscheduledaction


type RedshiftScheduledActionTargetActionResizeCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/redshift_scheduled_action#cluster_identifier RedshiftScheduledAction#cluster_identifier}.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/redshift_scheduled_action#classic RedshiftScheduledAction#classic}.
	Classic interface{} `field:"optional" json:"classic" yaml:"classic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/redshift_scheduled_action#cluster_type RedshiftScheduledAction#cluster_type}.
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/redshift_scheduled_action#node_type RedshiftScheduledAction#node_type}.
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/redshift_scheduled_action#number_of_nodes RedshiftScheduledAction#number_of_nodes}.
	NumberOfNodes *float64 `field:"optional" json:"numberOfNodes" yaml:"numberOfNodes"`
}

