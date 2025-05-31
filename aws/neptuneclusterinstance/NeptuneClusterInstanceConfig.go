// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package neptuneclusterinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NeptuneClusterInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#cluster_identifier NeptuneClusterInstance#cluster_identifier}.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#instance_class NeptuneClusterInstance#instance_class}.
	InstanceClass *string `field:"required" json:"instanceClass" yaml:"instanceClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#apply_immediately NeptuneClusterInstance#apply_immediately}.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#auto_minor_version_upgrade NeptuneClusterInstance#auto_minor_version_upgrade}.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#availability_zone NeptuneClusterInstance#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#engine NeptuneClusterInstance#engine}.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#engine_version NeptuneClusterInstance#engine_version}.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#id NeptuneClusterInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#identifier NeptuneClusterInstance#identifier}.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#identifier_prefix NeptuneClusterInstance#identifier_prefix}.
	IdentifierPrefix *string `field:"optional" json:"identifierPrefix" yaml:"identifierPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#neptune_parameter_group_name NeptuneClusterInstance#neptune_parameter_group_name}.
	NeptuneParameterGroupName *string `field:"optional" json:"neptuneParameterGroupName" yaml:"neptuneParameterGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#neptune_subnet_group_name NeptuneClusterInstance#neptune_subnet_group_name}.
	NeptuneSubnetGroupName *string `field:"optional" json:"neptuneSubnetGroupName" yaml:"neptuneSubnetGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#port NeptuneClusterInstance#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#preferred_backup_window NeptuneClusterInstance#preferred_backup_window}.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#preferred_maintenance_window NeptuneClusterInstance#preferred_maintenance_window}.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#promotion_tier NeptuneClusterInstance#promotion_tier}.
	PromotionTier *float64 `field:"optional" json:"promotionTier" yaml:"promotionTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#publicly_accessible NeptuneClusterInstance#publicly_accessible}.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#skip_final_snapshot NeptuneClusterInstance#skip_final_snapshot}.
	SkipFinalSnapshot interface{} `field:"optional" json:"skipFinalSnapshot" yaml:"skipFinalSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#tags NeptuneClusterInstance#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#tags_all NeptuneClusterInstance#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/neptune_cluster_instance#timeouts NeptuneClusterInstance#timeouts}
	Timeouts *NeptuneClusterInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

