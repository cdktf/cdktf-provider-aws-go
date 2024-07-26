// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsreplicationtask

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsReplicationTaskConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#migration_type DmsReplicationTask#migration_type}.
	MigrationType *string `field:"required" json:"migrationType" yaml:"migrationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#replication_instance_arn DmsReplicationTask#replication_instance_arn}.
	ReplicationInstanceArn *string `field:"required" json:"replicationInstanceArn" yaml:"replicationInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#replication_task_id DmsReplicationTask#replication_task_id}.
	ReplicationTaskId *string `field:"required" json:"replicationTaskId" yaml:"replicationTaskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#source_endpoint_arn DmsReplicationTask#source_endpoint_arn}.
	SourceEndpointArn *string `field:"required" json:"sourceEndpointArn" yaml:"sourceEndpointArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#table_mappings DmsReplicationTask#table_mappings}.
	TableMappings *string `field:"required" json:"tableMappings" yaml:"tableMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#target_endpoint_arn DmsReplicationTask#target_endpoint_arn}.
	TargetEndpointArn *string `field:"required" json:"targetEndpointArn" yaml:"targetEndpointArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#cdc_start_position DmsReplicationTask#cdc_start_position}.
	CdcStartPosition *string `field:"optional" json:"cdcStartPosition" yaml:"cdcStartPosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#cdc_start_time DmsReplicationTask#cdc_start_time}.
	CdcStartTime *string `field:"optional" json:"cdcStartTime" yaml:"cdcStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#id DmsReplicationTask#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#replication_task_settings DmsReplicationTask#replication_task_settings}.
	ReplicationTaskSettings *string `field:"optional" json:"replicationTaskSettings" yaml:"replicationTaskSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#resource_identifier DmsReplicationTask#resource_identifier}.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#start_replication_task DmsReplicationTask#start_replication_task}.
	StartReplicationTask interface{} `field:"optional" json:"startReplicationTask" yaml:"startReplicationTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#tags DmsReplicationTask#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dms_replication_task#tags_all DmsReplicationTask#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

