// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsrdsorderabledbinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsRdsOrderableDbInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#engine DataAwsRdsOrderableDbInstance#engine}.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#availability_zone_group DataAwsRdsOrderableDbInstance#availability_zone_group}.
	AvailabilityZoneGroup *string `field:"optional" json:"availabilityZoneGroup" yaml:"availabilityZoneGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#engine_latest_version DataAwsRdsOrderableDbInstance#engine_latest_version}.
	EngineLatestVersion interface{} `field:"optional" json:"engineLatestVersion" yaml:"engineLatestVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#engine_version DataAwsRdsOrderableDbInstance#engine_version}.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#id DataAwsRdsOrderableDbInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#instance_class DataAwsRdsOrderableDbInstance#instance_class}.
	InstanceClass *string `field:"optional" json:"instanceClass" yaml:"instanceClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#license_model DataAwsRdsOrderableDbInstance#license_model}.
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#preferred_engine_versions DataAwsRdsOrderableDbInstance#preferred_engine_versions}.
	PreferredEngineVersions *[]*string `field:"optional" json:"preferredEngineVersions" yaml:"preferredEngineVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#preferred_instance_classes DataAwsRdsOrderableDbInstance#preferred_instance_classes}.
	PreferredInstanceClasses *[]*string `field:"optional" json:"preferredInstanceClasses" yaml:"preferredInstanceClasses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#read_replica_capable DataAwsRdsOrderableDbInstance#read_replica_capable}.
	ReadReplicaCapable interface{} `field:"optional" json:"readReplicaCapable" yaml:"readReplicaCapable"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#region DataAwsRdsOrderableDbInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#storage_type DataAwsRdsOrderableDbInstance#storage_type}.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#supported_engine_modes DataAwsRdsOrderableDbInstance#supported_engine_modes}.
	SupportedEngineModes *[]*string `field:"optional" json:"supportedEngineModes" yaml:"supportedEngineModes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#supported_network_types DataAwsRdsOrderableDbInstance#supported_network_types}.
	SupportedNetworkTypes *[]*string `field:"optional" json:"supportedNetworkTypes" yaml:"supportedNetworkTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#supports_clusters DataAwsRdsOrderableDbInstance#supports_clusters}.
	SupportsClusters interface{} `field:"optional" json:"supportsClusters" yaml:"supportsClusters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#supports_enhanced_monitoring DataAwsRdsOrderableDbInstance#supports_enhanced_monitoring}.
	SupportsEnhancedMonitoring interface{} `field:"optional" json:"supportsEnhancedMonitoring" yaml:"supportsEnhancedMonitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#supports_global_databases DataAwsRdsOrderableDbInstance#supports_global_databases}.
	SupportsGlobalDatabases interface{} `field:"optional" json:"supportsGlobalDatabases" yaml:"supportsGlobalDatabases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#supports_iam_database_authentication DataAwsRdsOrderableDbInstance#supports_iam_database_authentication}.
	SupportsIamDatabaseAuthentication interface{} `field:"optional" json:"supportsIamDatabaseAuthentication" yaml:"supportsIamDatabaseAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#supports_iops DataAwsRdsOrderableDbInstance#supports_iops}.
	SupportsIops interface{} `field:"optional" json:"supportsIops" yaml:"supportsIops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#supports_kerberos_authentication DataAwsRdsOrderableDbInstance#supports_kerberos_authentication}.
	SupportsKerberosAuthentication interface{} `field:"optional" json:"supportsKerberosAuthentication" yaml:"supportsKerberosAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#supports_multi_az DataAwsRdsOrderableDbInstance#supports_multi_az}.
	SupportsMultiAz interface{} `field:"optional" json:"supportsMultiAz" yaml:"supportsMultiAz"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#supports_performance_insights DataAwsRdsOrderableDbInstance#supports_performance_insights}.
	SupportsPerformanceInsights interface{} `field:"optional" json:"supportsPerformanceInsights" yaml:"supportsPerformanceInsights"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#supports_storage_autoscaling DataAwsRdsOrderableDbInstance#supports_storage_autoscaling}.
	SupportsStorageAutoscaling interface{} `field:"optional" json:"supportsStorageAutoscaling" yaml:"supportsStorageAutoscaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#supports_storage_encryption DataAwsRdsOrderableDbInstance#supports_storage_encryption}.
	SupportsStorageEncryption interface{} `field:"optional" json:"supportsStorageEncryption" yaml:"supportsStorageEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/rds_orderable_db_instance#vpc DataAwsRdsOrderableDbInstance#vpc}.
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

