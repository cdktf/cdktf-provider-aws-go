// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daxcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DaxClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#cluster_name DaxCluster#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#iam_role_arn DaxCluster#iam_role_arn}.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#node_type DaxCluster#node_type}.
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#replication_factor DaxCluster#replication_factor}.
	ReplicationFactor *float64 `field:"required" json:"replicationFactor" yaml:"replicationFactor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#availability_zones DaxCluster#availability_zones}.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#cluster_endpoint_encryption_type DaxCluster#cluster_endpoint_encryption_type}.
	ClusterEndpointEncryptionType *string `field:"optional" json:"clusterEndpointEncryptionType" yaml:"clusterEndpointEncryptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#description DaxCluster#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#id DaxCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#maintenance_window DaxCluster#maintenance_window}.
	MaintenanceWindow *string `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#notification_topic_arn DaxCluster#notification_topic_arn}.
	NotificationTopicArn *string `field:"optional" json:"notificationTopicArn" yaml:"notificationTopicArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#parameter_group_name DaxCluster#parameter_group_name}.
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#region DaxCluster#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#security_group_ids DaxCluster#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// server_side_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#server_side_encryption DaxCluster#server_side_encryption}
	ServerSideEncryption *DaxClusterServerSideEncryption `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#subnet_group_name DaxCluster#subnet_group_name}.
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#tags DaxCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#tags_all DaxCluster#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/dax_cluster#timeouts DaxCluster#timeouts}
	Timeouts *DaxClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

