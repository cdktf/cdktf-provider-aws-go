// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/opsworksapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_application aws_opsworks_application}.
type OpsworksApplication interface {
	cdktf.TerraformResource
	AppSource() OpsworksApplicationAppSourceList
	AppSourceInput() interface{}
	AutoBundleOnDeploy() *string
	SetAutoBundleOnDeploy(val *string)
	AutoBundleOnDeployInput() *string
	AwsFlowRubySettings() *string
	SetAwsFlowRubySettings(val *string)
	AwsFlowRubySettingsInput() *string
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
	DataSourceArn() *string
	SetDataSourceArn(val *string)
	DataSourceArnInput() *string
	DataSourceDatabaseName() *string
	SetDataSourceDatabaseName(val *string)
	DataSourceDatabaseNameInput() *string
	DataSourceType() *string
	SetDataSourceType(val *string)
	DataSourceTypeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DocumentRoot() *string
	SetDocumentRoot(val *string)
	DocumentRootInput() *string
	Domains() *[]*string
	SetDomains(val *[]*string)
	DomainsInput() *[]*string
	EnableSsl() interface{}
	SetEnableSsl(val interface{})
	EnableSslInput() interface{}
	Environment() OpsworksApplicationEnvironmentList
	EnvironmentInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	RailsEnv() *string
	SetRailsEnv(val *string)
	RailsEnvInput() *string
	// Experimental.
	RawOverrides() interface{}
	ShortName() *string
	SetShortName(val *string)
	ShortNameInput() *string
	SslConfiguration() OpsworksApplicationSslConfigurationList
	SslConfigurationInput() interface{}
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAppSource(value interface{})
	PutEnvironment(value interface{})
	PutSslConfiguration(value interface{})
	ResetAppSource()
	ResetAutoBundleOnDeploy()
	ResetAwsFlowRubySettings()
	ResetDataSourceArn()
	ResetDataSourceDatabaseName()
	ResetDataSourceType()
	ResetDescription()
	ResetDocumentRoot()
	ResetDomains()
	ResetEnableSsl()
	ResetEnvironment()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRailsEnv()
	ResetShortName()
	ResetSslConfiguration()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksApplication
type jsiiProxy_OpsworksApplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksApplication) AppSource() OpsworksApplicationAppSourceList {
	var returns OpsworksApplicationAppSourceList
	_jsii_.Get(
		j,
		"appSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) AppSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) AutoBundleOnDeploy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoBundleOnDeploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) AutoBundleOnDeployInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoBundleOnDeployInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) AwsFlowRubySettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsFlowRubySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) AwsFlowRubySettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsFlowRubySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DataSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DataSourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DataSourceDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DataSourceDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DataSourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DataSourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DocumentRoot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DocumentRootInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Domains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) DomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) EnableSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) EnableSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Environment() OpsworksApplicationEnvironmentList {
	var returns OpsworksApplicationEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) RailsEnv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"railsEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) RailsEnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"railsEnvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) ShortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) ShortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) SslConfiguration() OpsworksApplicationSslConfigurationList {
	var returns OpsworksApplicationSslConfigurationList
	_jsii_.Get(
		j,
		"sslConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) SslConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksApplication) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_application aws_opsworks_application} Resource.
func NewOpsworksApplication(scope constructs.Construct, id *string, config *OpsworksApplicationConfig) OpsworksApplication {
	_init_.Initialize()

	if err := validateNewOpsworksApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsworksApplication{}

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksApplication.OpsworksApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_application aws_opsworks_application} Resource.
func NewOpsworksApplication_Override(o OpsworksApplication, scope constructs.Construct, id *string, config *OpsworksApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksApplication.OpsworksApplication",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetAutoBundleOnDeploy(val *string) {
	if err := j.validateSetAutoBundleOnDeployParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoBundleOnDeploy",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetAwsFlowRubySettings(val *string) {
	if err := j.validateSetAwsFlowRubySettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsFlowRubySettings",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetDataSourceArn(val *string) {
	if err := j.validateSetDataSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetDataSourceDatabaseName(val *string) {
	if err := j.validateSetDataSourceDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceDatabaseName",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetDataSourceType(val *string) {
	if err := j.validateSetDataSourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceType",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetDocumentRoot(val *string) {
	if err := j.validateSetDocumentRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentRoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetDomains(val *[]*string) {
	if err := j.validateSetDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domains",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetEnableSsl(val interface{}) {
	if err := j.validateSetEnableSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSsl",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetRailsEnv(val *string) {
	if err := j.validateSetRailsEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"railsEnv",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetShortName(val *string) {
	if err := j.validateSetShortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shortName",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetStackId(val *string) {
	if err := j.validateSetStackIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksApplication)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a OpsworksApplication resource upon running "cdktf plan <stack-name>".
func OpsworksApplication_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOpsworksApplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksApplication.OpsworksApplication",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func OpsworksApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksApplication.OpsworksApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpsworksApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksApplication.OpsworksApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpsworksApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksApplication.OpsworksApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.opsworksApplication.OpsworksApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OpsworksApplication) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OpsworksApplication) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OpsworksApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OpsworksApplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OpsworksApplication) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OpsworksApplication) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OpsworksApplication) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksApplication) PutAppSource(value interface{}) {
	if err := o.validatePutAppSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAppSource",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksApplication) PutEnvironment(value interface{}) {
	if err := o.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksApplication) PutSslConfiguration(value interface{}) {
	if err := o.validatePutSslConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSslConfiguration",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetAppSource() {
	_jsii_.InvokeVoid(
		o,
		"resetAppSource",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetAutoBundleOnDeploy() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoBundleOnDeploy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetAwsFlowRubySettings() {
	_jsii_.InvokeVoid(
		o,
		"resetAwsFlowRubySettings",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetDataSourceArn() {
	_jsii_.InvokeVoid(
		o,
		"resetDataSourceArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetDataSourceDatabaseName() {
	_jsii_.InvokeVoid(
		o,
		"resetDataSourceDatabaseName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetDataSourceType() {
	_jsii_.InvokeVoid(
		o,
		"resetDataSourceType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetDescription() {
	_jsii_.InvokeVoid(
		o,
		"resetDescription",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetDocumentRoot() {
	_jsii_.InvokeVoid(
		o,
		"resetDocumentRoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetDomains() {
	_jsii_.InvokeVoid(
		o,
		"resetDomains",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetEnableSsl() {
	_jsii_.InvokeVoid(
		o,
		"resetEnableSsl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetEnvironment() {
	_jsii_.InvokeVoid(
		o,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetRailsEnv() {
	_jsii_.InvokeVoid(
		o,
		"resetRailsEnv",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetShortName() {
	_jsii_.InvokeVoid(
		o,
		"resetShortName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) ResetSslConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetSslConfiguration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

