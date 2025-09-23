// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticachereplicationgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticacheReplicationGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#description ElasticacheReplicationGroup#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#replication_group_id ElasticacheReplicationGroup#replication_group_id}.
	ReplicationGroupId *string `field:"required" json:"replicationGroupId" yaml:"replicationGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#apply_immediately ElasticacheReplicationGroup#apply_immediately}.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#at_rest_encryption_enabled ElasticacheReplicationGroup#at_rest_encryption_enabled}.
	AtRestEncryptionEnabled *string `field:"optional" json:"atRestEncryptionEnabled" yaml:"atRestEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#auth_token ElasticacheReplicationGroup#auth_token}.
	AuthToken *string `field:"optional" json:"authToken" yaml:"authToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#auth_token_update_strategy ElasticacheReplicationGroup#auth_token_update_strategy}.
	AuthTokenUpdateStrategy *string `field:"optional" json:"authTokenUpdateStrategy" yaml:"authTokenUpdateStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#automatic_failover_enabled ElasticacheReplicationGroup#automatic_failover_enabled}.
	AutomaticFailoverEnabled interface{} `field:"optional" json:"automaticFailoverEnabled" yaml:"automaticFailoverEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#auto_minor_version_upgrade ElasticacheReplicationGroup#auto_minor_version_upgrade}.
	AutoMinorVersionUpgrade *string `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#cluster_mode ElasticacheReplicationGroup#cluster_mode}.
	ClusterMode *string `field:"optional" json:"clusterMode" yaml:"clusterMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#data_tiering_enabled ElasticacheReplicationGroup#data_tiering_enabled}.
	DataTieringEnabled interface{} `field:"optional" json:"dataTieringEnabled" yaml:"dataTieringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#engine ElasticacheReplicationGroup#engine}.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#engine_version ElasticacheReplicationGroup#engine_version}.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#final_snapshot_identifier ElasticacheReplicationGroup#final_snapshot_identifier}.
	FinalSnapshotIdentifier *string `field:"optional" json:"finalSnapshotIdentifier" yaml:"finalSnapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#global_replication_group_id ElasticacheReplicationGroup#global_replication_group_id}.
	GlobalReplicationGroupId *string `field:"optional" json:"globalReplicationGroupId" yaml:"globalReplicationGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#id ElasticacheReplicationGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#ip_discovery ElasticacheReplicationGroup#ip_discovery}.
	IpDiscovery *string `field:"optional" json:"ipDiscovery" yaml:"ipDiscovery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#kms_key_id ElasticacheReplicationGroup#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// log_delivery_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#log_delivery_configuration ElasticacheReplicationGroup#log_delivery_configuration}
	LogDeliveryConfiguration interface{} `field:"optional" json:"logDeliveryConfiguration" yaml:"logDeliveryConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#maintenance_window ElasticacheReplicationGroup#maintenance_window}.
	MaintenanceWindow *string `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#multi_az_enabled ElasticacheReplicationGroup#multi_az_enabled}.
	MultiAzEnabled interface{} `field:"optional" json:"multiAzEnabled" yaml:"multiAzEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#network_type ElasticacheReplicationGroup#network_type}.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#node_type ElasticacheReplicationGroup#node_type}.
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#notification_topic_arn ElasticacheReplicationGroup#notification_topic_arn}.
	NotificationTopicArn *string `field:"optional" json:"notificationTopicArn" yaml:"notificationTopicArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#num_cache_clusters ElasticacheReplicationGroup#num_cache_clusters}.
	NumCacheClusters *float64 `field:"optional" json:"numCacheClusters" yaml:"numCacheClusters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#num_node_groups ElasticacheReplicationGroup#num_node_groups}.
	NumNodeGroups *float64 `field:"optional" json:"numNodeGroups" yaml:"numNodeGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#parameter_group_name ElasticacheReplicationGroup#parameter_group_name}.
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#port ElasticacheReplicationGroup#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#preferred_cache_cluster_azs ElasticacheReplicationGroup#preferred_cache_cluster_azs}.
	PreferredCacheClusterAzs *[]*string `field:"optional" json:"preferredCacheClusterAzs" yaml:"preferredCacheClusterAzs"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#region ElasticacheReplicationGroup#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#replicas_per_node_group ElasticacheReplicationGroup#replicas_per_node_group}.
	ReplicasPerNodeGroup *float64 `field:"optional" json:"replicasPerNodeGroup" yaml:"replicasPerNodeGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#security_group_ids ElasticacheReplicationGroup#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#security_group_names ElasticacheReplicationGroup#security_group_names}.
	SecurityGroupNames *[]*string `field:"optional" json:"securityGroupNames" yaml:"securityGroupNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#snapshot_arns ElasticacheReplicationGroup#snapshot_arns}.
	SnapshotArns *[]*string `field:"optional" json:"snapshotArns" yaml:"snapshotArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#snapshot_name ElasticacheReplicationGroup#snapshot_name}.
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#snapshot_retention_limit ElasticacheReplicationGroup#snapshot_retention_limit}.
	SnapshotRetentionLimit *float64 `field:"optional" json:"snapshotRetentionLimit" yaml:"snapshotRetentionLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#snapshot_window ElasticacheReplicationGroup#snapshot_window}.
	SnapshotWindow *string `field:"optional" json:"snapshotWindow" yaml:"snapshotWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#subnet_group_name ElasticacheReplicationGroup#subnet_group_name}.
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#tags ElasticacheReplicationGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#tags_all ElasticacheReplicationGroup#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#timeouts ElasticacheReplicationGroup#timeouts}
	Timeouts *ElasticacheReplicationGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#transit_encryption_enabled ElasticacheReplicationGroup#transit_encryption_enabled}.
	TransitEncryptionEnabled interface{} `field:"optional" json:"transitEncryptionEnabled" yaml:"transitEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#transit_encryption_mode ElasticacheReplicationGroup#transit_encryption_mode}.
	TransitEncryptionMode *string `field:"optional" json:"transitEncryptionMode" yaml:"transitEncryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticache_replication_group#user_group_ids ElasticacheReplicationGroup#user_group_ids}.
	UserGroupIds *[]*string `field:"optional" json:"userGroupIds" yaml:"userGroupIds"`
}

