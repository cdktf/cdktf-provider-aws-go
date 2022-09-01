// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrafanaWorkspaceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#account_access_type GrafanaWorkspace#account_access_type}.
	AccountAccessType *string `field:"required" json:"accountAccessType" yaml:"accountAccessType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#authentication_providers GrafanaWorkspace#authentication_providers}.
	AuthenticationProviders *[]*string `field:"required" json:"authenticationProviders" yaml:"authenticationProviders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#permission_type GrafanaWorkspace#permission_type}.
	PermissionType *string `field:"required" json:"permissionType" yaml:"permissionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#data_sources GrafanaWorkspace#data_sources}.
	DataSources *[]*string `field:"optional" json:"dataSources" yaml:"dataSources"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#description GrafanaWorkspace#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#id GrafanaWorkspace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#name GrafanaWorkspace#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#notification_destinations GrafanaWorkspace#notification_destinations}.
	NotificationDestinations *[]*string `field:"optional" json:"notificationDestinations" yaml:"notificationDestinations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#organizational_units GrafanaWorkspace#organizational_units}.
	OrganizationalUnits *[]*string `field:"optional" json:"organizationalUnits" yaml:"organizationalUnits"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#organization_role_name GrafanaWorkspace#organization_role_name}.
	OrganizationRoleName *string `field:"optional" json:"organizationRoleName" yaml:"organizationRoleName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#role_arn GrafanaWorkspace#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#stack_set_name GrafanaWorkspace#stack_set_name}.
	StackSetName *string `field:"optional" json:"stackSetName" yaml:"stackSetName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#tags GrafanaWorkspace#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#tags_all GrafanaWorkspace#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/grafana_workspace#timeouts GrafanaWorkspace#timeouts}
	Timeouts *GrafanaWorkspaceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

