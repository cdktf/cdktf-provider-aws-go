// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mqbroker

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MqBrokerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#broker_name MqBroker#broker_name}.
	BrokerName *string `field:"required" json:"brokerName" yaml:"brokerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#engine_type MqBroker#engine_type}.
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#engine_version MqBroker#engine_version}.
	EngineVersion *string `field:"required" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#host_instance_type MqBroker#host_instance_type}.
	HostInstanceType *string `field:"required" json:"hostInstanceType" yaml:"hostInstanceType"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#user MqBroker#user}
	User interface{} `field:"required" json:"user" yaml:"user"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#apply_immediately MqBroker#apply_immediately}.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#authentication_strategy MqBroker#authentication_strategy}.
	AuthenticationStrategy *string `field:"optional" json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#auto_minor_version_upgrade MqBroker#auto_minor_version_upgrade}.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#configuration MqBroker#configuration}
	Configuration *MqBrokerConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#data_replication_mode MqBroker#data_replication_mode}.
	DataReplicationMode *string `field:"optional" json:"dataReplicationMode" yaml:"dataReplicationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#data_replication_primary_broker_arn MqBroker#data_replication_primary_broker_arn}.
	DataReplicationPrimaryBrokerArn *string `field:"optional" json:"dataReplicationPrimaryBrokerArn" yaml:"dataReplicationPrimaryBrokerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#deployment_mode MqBroker#deployment_mode}.
	DeploymentMode *string `field:"optional" json:"deploymentMode" yaml:"deploymentMode"`
	// encryption_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#encryption_options MqBroker#encryption_options}
	EncryptionOptions *MqBrokerEncryptionOptions `field:"optional" json:"encryptionOptions" yaml:"encryptionOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#id MqBroker#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ldap_server_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#ldap_server_metadata MqBroker#ldap_server_metadata}
	LdapServerMetadata *MqBrokerLdapServerMetadata `field:"optional" json:"ldapServerMetadata" yaml:"ldapServerMetadata"`
	// logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#logs MqBroker#logs}
	Logs *MqBrokerLogs `field:"optional" json:"logs" yaml:"logs"`
	// maintenance_window_start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#maintenance_window_start_time MqBroker#maintenance_window_start_time}
	MaintenanceWindowStartTime *MqBrokerMaintenanceWindowStartTime `field:"optional" json:"maintenanceWindowStartTime" yaml:"maintenanceWindowStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#publicly_accessible MqBroker#publicly_accessible}.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#region MqBroker#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#security_groups MqBroker#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#storage_type MqBroker#storage_type}.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#subnet_ids MqBroker#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#tags MqBroker#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#tags_all MqBroker#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/mq_broker#timeouts MqBroker#timeouts}
	Timeouts *MqBrokerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

