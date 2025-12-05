// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreaminfluxdbdbcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TimestreaminfluxdbDbClusterConfig struct {
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
	// The Timestream for InfluxDB DB instance type to run InfluxDB on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#db_instance_type TimestreaminfluxdbDbCluster#db_instance_type}
	DbInstanceType *string `field:"required" json:"dbInstanceType" yaml:"dbInstanceType"`
	// The name that uniquely identifies the DB cluster when interacting with the  					Amazon Timestream for InfluxDB API and CLI commands.
	//
	// This name will also be a
	// 					prefix included in the endpoint. DB cluster names must be unique per customer
	// 					and per region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#name TimestreaminfluxdbDbCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of VPC security group IDs to associate with the Timestream for InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#vpc_security_group_ids TimestreaminfluxdbDbCluster#vpc_security_group_ids}
	VpcSecurityGroupIds *[]*string `field:"required" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// A list of VPC subnet IDs to associate with the DB cluster.
	//
	// Provide at least
	// 					two VPC subnet IDs in different availability zones when deploying with a Multi-AZ standby.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#vpc_subnet_ids TimestreaminfluxdbDbCluster#vpc_subnet_ids}
	VpcSubnetIds *[]*string `field:"required" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
	// The amount of storage to allocate for your DB storage type in GiB (gibibytes).
	//
	// This field is forbidden for InfluxDB V3 clusters (when using an InfluxDB V3 db parameter group).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#allocated_storage TimestreaminfluxdbDbCluster#allocated_storage}
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// Name of the initial InfluxDB bucket.
	//
	// All InfluxDB data is stored in a bucket.
	// 					A bucket combines the concept of a database and a retention period (the duration of time
	// 					that each data point persists). A bucket belongs to an organization. Along with organization,
	// 					username, and password, this argument will be stored in the secret referred to by the
	// 					influx_auth_parameters_secret_arn attribute. This field is forbidden for InfluxDB V3 clusters
	// 					(when using an InfluxDB V3 db parameter group).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#bucket TimestreaminfluxdbDbCluster#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The ID of the DB parameter group to assign to your DB cluster.
	//
	// DB parameter groups specify how the database is configured. For example, DB parameter groups
	// 					can specify the limit for query concurrency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#db_parameter_group_identifier TimestreaminfluxdbDbCluster#db_parameter_group_identifier}
	DbParameterGroupIdentifier *string `field:"optional" json:"dbParameterGroupIdentifier" yaml:"dbParameterGroupIdentifier"`
	// The Timestream for InfluxDB DB storage type to read and write InfluxDB data.
	//
	// You can choose between 3 different types of provisioned Influx IOPS included storage according
	// 					to your workloads requirements: Influx IO Included 3000 IOPS, Influx IO Included 12000 IOPS,
	// 					Influx IO Included 16000 IOPS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#db_storage_type TimestreaminfluxdbDbCluster#db_storage_type}
	DbStorageType *string `field:"optional" json:"dbStorageType" yaml:"dbStorageType"`
	// Specifies the type of cluster to create.
	//
	// This field is forbidden for InfluxDB V3 clusters
	// 					(when using an InfluxDB V3 db parameter group).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#deployment_type TimestreaminfluxdbDbCluster#deployment_type}
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// Specifies the behavior of failure recovery when the primary node of the cluster 					fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#failover_mode TimestreaminfluxdbDbCluster#failover_mode}
	FailoverMode *string `field:"optional" json:"failoverMode" yaml:"failoverMode"`
	// log_delivery_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#log_delivery_configuration TimestreaminfluxdbDbCluster#log_delivery_configuration}
	LogDeliveryConfiguration interface{} `field:"optional" json:"logDeliveryConfiguration" yaml:"logDeliveryConfiguration"`
	// Specifies whether the networkType of the Timestream for InfluxDB cluster is  					IPV4, which can communicate over IPv4 protocol only, or DUAL, which can communicate  					over both IPv4 and IPv6 protocols.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#network_type TimestreaminfluxdbDbCluster#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Name of the initial organization for the initial admin user in InfluxDB.
	//
	// An
	// 					InfluxDB organization is a workspace for a group of users. Along with bucket, username,
	// 					and password, this argument will be stored in the secret referred to by the
	// 					influx_auth_parameters_secret_arn attribute. This field is forbidden for InfluxDB V3 clusters
	// 					(when using an InfluxDB V3 db parameter group).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#organization TimestreaminfluxdbDbCluster#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Password of the initial admin user created in InfluxDB.
	//
	// This password will
	// 					allow you to access the InfluxDB UI to perform various administrative tasks and
	// 					also use the InfluxDB CLI to create an operator token. Along with bucket, username,
	// 					and organization, this argument will be stored in the secret referred to by the
	// 					influx_auth_parameters_secret_arn attribute. This field is forbidden for InfluxDB V3 clusters
	// 					(when using an InfluxDB V3 db parameter group) as the AWS API rejects it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#password TimestreaminfluxdbDbCluster#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port number on which InfluxDB accepts connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#port TimestreaminfluxdbDbCluster#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Configures the Timestream for InfluxDB cluster with a public IP to facilitate access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#publicly_accessible TimestreaminfluxdbDbCluster#publicly_accessible}
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#region TimestreaminfluxdbDbCluster#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#tags TimestreaminfluxdbDbCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#timeouts TimestreaminfluxdbDbCluster#timeouts}
	Timeouts *TimestreaminfluxdbDbClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Username of the initial admin user created in InfluxDB.
	//
	// Must start with a letter
	// 					and can't end with a hyphen or contain two consecutive hyphens. This username will allow
	// 					you to access the InfluxDB UI to perform various administrative tasks and also use the
	// 					InfluxDB CLI to create an operator token. Along with bucket, organization, and password,
	// 					this argument will be stored in the secret referred to by the influx_auth_parameters_secret_arn
	// 					attribute. This field is forbidden for InfluxDB V3 clusters (when using an InfluxDB V3 db parameter group).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreaminfluxdb_db_cluster#username TimestreaminfluxdbDbCluster#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

