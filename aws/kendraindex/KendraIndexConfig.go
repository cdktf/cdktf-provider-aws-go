package kendraindex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraIndexConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#name KendraIndex#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#role_arn KendraIndex#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// capacity_units block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#capacity_units KendraIndex#capacity_units}
	CapacityUnits *KendraIndexCapacityUnits `field:"optional" json:"capacityUnits" yaml:"capacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#description KendraIndex#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// document_metadata_configuration_updates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#document_metadata_configuration_updates KendraIndex#document_metadata_configuration_updates}
	DocumentMetadataConfigurationUpdates interface{} `field:"optional" json:"documentMetadataConfigurationUpdates" yaml:"documentMetadataConfigurationUpdates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#edition KendraIndex#edition}.
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#id KendraIndex#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// server_side_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#server_side_encryption_configuration KendraIndex#server_side_encryption_configuration}
	ServerSideEncryptionConfiguration *KendraIndexServerSideEncryptionConfiguration `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#tags KendraIndex#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#tags_all KendraIndex#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#timeouts KendraIndex#timeouts}
	Timeouts *KendraIndexTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#user_context_policy KendraIndex#user_context_policy}.
	UserContextPolicy *string `field:"optional" json:"userContextPolicy" yaml:"userContextPolicy"`
	// user_group_resolution_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#user_group_resolution_configuration KendraIndex#user_group_resolution_configuration}
	UserGroupResolutionConfiguration *KendraIndexUserGroupResolutionConfiguration `field:"optional" json:"userGroupResolutionConfiguration" yaml:"userGroupResolutionConfiguration"`
	// user_token_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/kendra_index#user_token_configurations KendraIndex#user_token_configurations}
	UserTokenConfigurations *KendraIndexUserTokenConfigurations `field:"optional" json:"userTokenConfigurations" yaml:"userTokenConfigurations"`
}

