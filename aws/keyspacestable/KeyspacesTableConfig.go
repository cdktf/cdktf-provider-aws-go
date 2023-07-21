package keyspacestable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyspacesTableConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#keyspace_name KeyspacesTable#keyspace_name}.
	KeyspaceName *string `field:"required" json:"keyspaceName" yaml:"keyspaceName"`
	// schema_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#schema_definition KeyspacesTable#schema_definition}
	SchemaDefinition *KeyspacesTableSchemaDefinition `field:"required" json:"schemaDefinition" yaml:"schemaDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#table_name KeyspacesTable#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// capacity_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#capacity_specification KeyspacesTable#capacity_specification}
	CapacitySpecification *KeyspacesTableCapacitySpecification `field:"optional" json:"capacitySpecification" yaml:"capacitySpecification"`
	// client_side_timestamps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#client_side_timestamps KeyspacesTable#client_side_timestamps}
	ClientSideTimestamps *KeyspacesTableClientSideTimestamps `field:"optional" json:"clientSideTimestamps" yaml:"clientSideTimestamps"`
	// comment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#comment KeyspacesTable#comment}
	Comment *KeyspacesTableComment `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#default_time_to_live KeyspacesTable#default_time_to_live}.
	DefaultTimeToLive *float64 `field:"optional" json:"defaultTimeToLive" yaml:"defaultTimeToLive"`
	// encryption_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#encryption_specification KeyspacesTable#encryption_specification}
	EncryptionSpecification *KeyspacesTableEncryptionSpecification `field:"optional" json:"encryptionSpecification" yaml:"encryptionSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#id KeyspacesTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// point_in_time_recovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#point_in_time_recovery KeyspacesTable#point_in_time_recovery}
	PointInTimeRecovery *KeyspacesTablePointInTimeRecovery `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#tags KeyspacesTable#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#tags_all KeyspacesTable#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#timeouts KeyspacesTable#timeouts}
	Timeouts *KeyspacesTableTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/keyspaces_table#ttl KeyspacesTable#ttl}
	Ttl *KeyspacesTableTtl `field:"optional" json:"ttl" yaml:"ttl"`
}

