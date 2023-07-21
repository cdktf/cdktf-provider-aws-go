package lakeformationdatalakesettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LakeformationDataLakeSettingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lakeformation_data_lake_settings#admins LakeformationDataLakeSettings#admins}.
	Admins *[]*string `field:"optional" json:"admins" yaml:"admins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lakeformation_data_lake_settings#allow_external_data_filtering LakeformationDataLakeSettings#allow_external_data_filtering}.
	AllowExternalDataFiltering interface{} `field:"optional" json:"allowExternalDataFiltering" yaml:"allowExternalDataFiltering"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lakeformation_data_lake_settings#authorized_session_tag_value_list LakeformationDataLakeSettings#authorized_session_tag_value_list}.
	AuthorizedSessionTagValueList *[]*string `field:"optional" json:"authorizedSessionTagValueList" yaml:"authorizedSessionTagValueList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lakeformation_data_lake_settings#catalog_id LakeformationDataLakeSettings#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// create_database_default_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lakeformation_data_lake_settings#create_database_default_permissions LakeformationDataLakeSettings#create_database_default_permissions}
	CreateDatabaseDefaultPermissions interface{} `field:"optional" json:"createDatabaseDefaultPermissions" yaml:"createDatabaseDefaultPermissions"`
	// create_table_default_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lakeformation_data_lake_settings#create_table_default_permissions LakeformationDataLakeSettings#create_table_default_permissions}
	CreateTableDefaultPermissions interface{} `field:"optional" json:"createTableDefaultPermissions" yaml:"createTableDefaultPermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lakeformation_data_lake_settings#external_data_filtering_allow_list LakeformationDataLakeSettings#external_data_filtering_allow_list}.
	ExternalDataFilteringAllowList *[]*string `field:"optional" json:"externalDataFilteringAllowList" yaml:"externalDataFilteringAllowList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lakeformation_data_lake_settings#id LakeformationDataLakeSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/lakeformation_data_lake_settings#trusted_resource_owners LakeformationDataLakeSettings#trusted_resource_owners}.
	TrustedResourceOwners *[]*string `field:"optional" json:"trustedResourceOwners" yaml:"trustedResourceOwners"`
}

