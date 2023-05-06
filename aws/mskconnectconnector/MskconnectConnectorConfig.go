package mskconnectconnector

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MskconnectConnectorConfig struct {
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
	// capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#capacity MskconnectConnector#capacity}
	Capacity *MskconnectConnectorCapacity `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#connector_configuration MskconnectConnector#connector_configuration}.
	ConnectorConfiguration *map[string]*string `field:"required" json:"connectorConfiguration" yaml:"connectorConfiguration"`
	// kafka_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#kafka_cluster MskconnectConnector#kafka_cluster}
	KafkaCluster *MskconnectConnectorKafkaCluster `field:"required" json:"kafkaCluster" yaml:"kafkaCluster"`
	// kafka_cluster_client_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#kafka_cluster_client_authentication MskconnectConnector#kafka_cluster_client_authentication}
	KafkaClusterClientAuthentication *MskconnectConnectorKafkaClusterClientAuthentication `field:"required" json:"kafkaClusterClientAuthentication" yaml:"kafkaClusterClientAuthentication"`
	// kafka_cluster_encryption_in_transit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#kafka_cluster_encryption_in_transit MskconnectConnector#kafka_cluster_encryption_in_transit}
	KafkaClusterEncryptionInTransit *MskconnectConnectorKafkaClusterEncryptionInTransit `field:"required" json:"kafkaClusterEncryptionInTransit" yaml:"kafkaClusterEncryptionInTransit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#kafkaconnect_version MskconnectConnector#kafkaconnect_version}.
	KafkaconnectVersion *string `field:"required" json:"kafkaconnectVersion" yaml:"kafkaconnectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#name MskconnectConnector#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// plugin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#plugin MskconnectConnector#plugin}
	Plugin interface{} `field:"required" json:"plugin" yaml:"plugin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#service_execution_role_arn MskconnectConnector#service_execution_role_arn}.
	ServiceExecutionRoleArn *string `field:"required" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#description MskconnectConnector#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#id MskconnectConnector#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// log_delivery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#log_delivery MskconnectConnector#log_delivery}
	LogDelivery *MskconnectConnectorLogDelivery `field:"optional" json:"logDelivery" yaml:"logDelivery"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#timeouts MskconnectConnector#timeouts}
	Timeouts *MskconnectConnectorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// worker_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/mskconnect_connector#worker_configuration MskconnectConnector#worker_configuration}
	WorkerConfiguration *MskconnectConnectorWorkerConfiguration `field:"optional" json:"workerConfiguration" yaml:"workerConfiguration"`
}

