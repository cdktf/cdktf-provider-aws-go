// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lightsaildatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LightsailDatabaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#blueprint_id LightsailDatabase#blueprint_id}.
	BlueprintId *string `field:"required" json:"blueprintId" yaml:"blueprintId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#bundle_id LightsailDatabase#bundle_id}.
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#master_database_name LightsailDatabase#master_database_name}.
	MasterDatabaseName *string `field:"required" json:"masterDatabaseName" yaml:"masterDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#master_password LightsailDatabase#master_password}.
	MasterPassword *string `field:"required" json:"masterPassword" yaml:"masterPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#master_username LightsailDatabase#master_username}.
	MasterUsername *string `field:"required" json:"masterUsername" yaml:"masterUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#relational_database_name LightsailDatabase#relational_database_name}.
	RelationalDatabaseName *string `field:"required" json:"relationalDatabaseName" yaml:"relationalDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#apply_immediately LightsailDatabase#apply_immediately}.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#availability_zone LightsailDatabase#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#backup_retention_enabled LightsailDatabase#backup_retention_enabled}.
	BackupRetentionEnabled interface{} `field:"optional" json:"backupRetentionEnabled" yaml:"backupRetentionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#final_snapshot_name LightsailDatabase#final_snapshot_name}.
	FinalSnapshotName *string `field:"optional" json:"finalSnapshotName" yaml:"finalSnapshotName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#id LightsailDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#preferred_backup_window LightsailDatabase#preferred_backup_window}.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#preferred_maintenance_window LightsailDatabase#preferred_maintenance_window}.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#publicly_accessible LightsailDatabase#publicly_accessible}.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#skip_final_snapshot LightsailDatabase#skip_final_snapshot}.
	SkipFinalSnapshot interface{} `field:"optional" json:"skipFinalSnapshot" yaml:"skipFinalSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#tags LightsailDatabase#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/lightsail_database#tags_all LightsailDatabase#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

