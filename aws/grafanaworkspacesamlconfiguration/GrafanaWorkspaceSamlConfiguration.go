// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grafanaworkspacesamlconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/grafanaworkspacesamlconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.17.0/docs/resources/grafana_workspace_saml_configuration aws_grafana_workspace_saml_configuration}.
type GrafanaWorkspaceSamlConfiguration interface {
	cdktf.TerraformResource
	AdminRoleValues() *[]*string
	SetAdminRoleValues(val *[]*string)
	AdminRoleValuesInput() *[]*string
	AllowedOrganizations() *[]*string
	SetAllowedOrganizations(val *[]*string)
	AllowedOrganizationsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EditorRoleValues() *[]*string
	SetEditorRoleValues(val *[]*string)
	EditorRoleValuesInput() *[]*string
	EmailAssertion() *string
	SetEmailAssertion(val *string)
	EmailAssertionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupsAssertion() *string
	SetGroupsAssertion(val *string)
	GroupsAssertionInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IdpMetadataUrl() *string
	SetIdpMetadataUrl(val *string)
	IdpMetadataUrlInput() *string
	IdpMetadataXml() *string
	SetIdpMetadataXml(val *string)
	IdpMetadataXmlInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoginAssertion() *string
	SetLoginAssertion(val *string)
	LoginAssertionInput() *string
	LoginValidityDuration() *float64
	SetLoginValidityDuration(val *float64)
	LoginValidityDurationInput() *float64
	NameAssertion() *string
	SetNameAssertion(val *string)
	NameAssertionInput() *string
	// The tree node.
	Node() constructs.Node
	OrgAssertion() *string
	SetOrgAssertion(val *string)
	OrgAssertionInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RoleAssertion() *string
	SetRoleAssertion(val *string)
	RoleAssertionInput() *string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GrafanaWorkspaceSamlConfigurationTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *GrafanaWorkspaceSamlConfigurationTimeouts)
	ResetAdminRoleValues()
	ResetAllowedOrganizations()
	ResetEmailAssertion()
	ResetGroupsAssertion()
	ResetId()
	ResetIdpMetadataUrl()
	ResetIdpMetadataXml()
	ResetLoginAssertion()
	ResetLoginValidityDuration()
	ResetNameAssertion()
	ResetOrgAssertion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRoleAssertion()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GrafanaWorkspaceSamlConfiguration
type jsiiProxy_GrafanaWorkspaceSamlConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) AdminRoleValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminRoleValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) AdminRoleValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminRoleValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) AllowedOrganizations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrganizations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) AllowedOrganizationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrganizationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) EditorRoleValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"editorRoleValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) EditorRoleValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"editorRoleValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) EmailAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) EmailAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) GroupsAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) GroupsAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) IdpMetadataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) IdpMetadataUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) IdpMetadataXml() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataXml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) IdpMetadataXmlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataXmlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) LoginAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) LoginAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) LoginValidityDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loginValidityDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) LoginValidityDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loginValidityDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) NameAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) NameAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) OrgAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) OrgAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) RoleAssertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) RoleAssertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) Timeouts() GrafanaWorkspaceSamlConfigurationTimeoutsOutputReference {
	var returns GrafanaWorkspaceSamlConfigurationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.17.0/docs/resources/grafana_workspace_saml_configuration aws_grafana_workspace_saml_configuration} Resource.
func NewGrafanaWorkspaceSamlConfiguration(scope constructs.Construct, id *string, config *GrafanaWorkspaceSamlConfigurationConfig) GrafanaWorkspaceSamlConfiguration {
	_init_.Initialize()

	if err := validateNewGrafanaWorkspaceSamlConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GrafanaWorkspaceSamlConfiguration{}

	_jsii_.Create(
		"@cdktf/provider-aws.grafanaWorkspaceSamlConfiguration.GrafanaWorkspaceSamlConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.17.0/docs/resources/grafana_workspace_saml_configuration aws_grafana_workspace_saml_configuration} Resource.
func NewGrafanaWorkspaceSamlConfiguration_Override(g GrafanaWorkspaceSamlConfiguration, scope constructs.Construct, id *string, config *GrafanaWorkspaceSamlConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.grafanaWorkspaceSamlConfiguration.GrafanaWorkspaceSamlConfiguration",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetAdminRoleValues(val *[]*string) {
	if err := j.validateSetAdminRoleValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminRoleValues",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetAllowedOrganizations(val *[]*string) {
	if err := j.validateSetAllowedOrganizationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOrganizations",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetEditorRoleValues(val *[]*string) {
	if err := j.validateSetEditorRoleValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"editorRoleValues",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetEmailAssertion(val *string) {
	if err := j.validateSetEmailAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAssertion",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetGroupsAssertion(val *string) {
	if err := j.validateSetGroupsAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsAssertion",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetIdpMetadataUrl(val *string) {
	if err := j.validateSetIdpMetadataUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpMetadataUrl",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetIdpMetadataXml(val *string) {
	if err := j.validateSetIdpMetadataXmlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpMetadataXml",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetLoginAssertion(val *string) {
	if err := j.validateSetLoginAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginAssertion",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetLoginValidityDuration(val *float64) {
	if err := j.validateSetLoginValidityDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginValidityDuration",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetNameAssertion(val *string) {
	if err := j.validateSetNameAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameAssertion",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetOrgAssertion(val *string) {
	if err := j.validateSetOrgAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgAssertion",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetRoleAssertion(val *string) {
	if err := j.validateSetRoleAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleAssertion",
		val,
	)
}

func (j *jsiiProxy_GrafanaWorkspaceSamlConfiguration)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func GrafanaWorkspaceSamlConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrafanaWorkspaceSamlConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.grafanaWorkspaceSamlConfiguration.GrafanaWorkspaceSamlConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GrafanaWorkspaceSamlConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrafanaWorkspaceSamlConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.grafanaWorkspaceSamlConfiguration.GrafanaWorkspaceSamlConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GrafanaWorkspaceSamlConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrafanaWorkspaceSamlConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.grafanaWorkspaceSamlConfiguration.GrafanaWorkspaceSamlConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GrafanaWorkspaceSamlConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.grafanaWorkspaceSamlConfiguration.GrafanaWorkspaceSamlConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) PutTimeouts(value *GrafanaWorkspaceSamlConfigurationTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetAdminRoleValues() {
	_jsii_.InvokeVoid(
		g,
		"resetAdminRoleValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetAllowedOrganizations() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedOrganizations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetEmailAssertion() {
	_jsii_.InvokeVoid(
		g,
		"resetEmailAssertion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetGroupsAssertion() {
	_jsii_.InvokeVoid(
		g,
		"resetGroupsAssertion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetIdpMetadataUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetIdpMetadataUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetIdpMetadataXml() {
	_jsii_.InvokeVoid(
		g,
		"resetIdpMetadataXml",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetLoginAssertion() {
	_jsii_.InvokeVoid(
		g,
		"resetLoginAssertion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetLoginValidityDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetLoginValidityDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetNameAssertion() {
	_jsii_.InvokeVoid(
		g,
		"resetNameAssertion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetOrgAssertion() {
	_jsii_.InvokeVoid(
		g,
		"resetOrgAssertion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetRoleAssertion() {
	_jsii_.InvokeVoid(
		g,
		"resetRoleAssertion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrafanaWorkspaceSamlConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

