// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/appsyncdatasource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appsync_datasource aws_appsync_datasource}.
type AppsyncDatasource interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	Arn() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DynamodbConfig() AppsyncDatasourceDynamodbConfigOutputReference
	DynamodbConfigInput() *AppsyncDatasourceDynamodbConfig
	ElasticsearchConfig() AppsyncDatasourceElasticsearchConfigOutputReference
	ElasticsearchConfigInput() *AppsyncDatasourceElasticsearchConfig
	EventBridgeConfig() AppsyncDatasourceEventBridgeConfigOutputReference
	EventBridgeConfigInput() *AppsyncDatasourceEventBridgeConfig
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpConfig() AppsyncDatasourceHttpConfigOutputReference
	HttpConfigInput() *AppsyncDatasourceHttpConfig
	Id() *string
	SetId(val *string)
	IdInput() *string
	LambdaConfig() AppsyncDatasourceLambdaConfigOutputReference
	LambdaConfigInput() *AppsyncDatasourceLambdaConfig
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OpensearchserviceConfig() AppsyncDatasourceOpensearchserviceConfigOutputReference
	OpensearchserviceConfigInput() *AppsyncDatasourceOpensearchserviceConfig
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
	RelationalDatabaseConfig() AppsyncDatasourceRelationalDatabaseConfigOutputReference
	RelationalDatabaseConfigInput() *AppsyncDatasourceRelationalDatabaseConfig
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	ServiceRoleArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutDynamodbConfig(value *AppsyncDatasourceDynamodbConfig)
	PutElasticsearchConfig(value *AppsyncDatasourceElasticsearchConfig)
	PutEventBridgeConfig(value *AppsyncDatasourceEventBridgeConfig)
	PutHttpConfig(value *AppsyncDatasourceHttpConfig)
	PutLambdaConfig(value *AppsyncDatasourceLambdaConfig)
	PutOpensearchserviceConfig(value *AppsyncDatasourceOpensearchserviceConfig)
	PutRelationalDatabaseConfig(value *AppsyncDatasourceRelationalDatabaseConfig)
	ResetDescription()
	ResetDynamodbConfig()
	ResetElasticsearchConfig()
	ResetEventBridgeConfig()
	ResetHttpConfig()
	ResetId()
	ResetLambdaConfig()
	ResetOpensearchserviceConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRelationalDatabaseConfig()
	ResetServiceRoleArn()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AppsyncDatasource
type jsiiProxy_AppsyncDatasource struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsyncDatasource) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) DynamodbConfig() AppsyncDatasourceDynamodbConfigOutputReference {
	var returns AppsyncDatasourceDynamodbConfigOutputReference
	_jsii_.Get(
		j,
		"dynamodbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) DynamodbConfigInput() *AppsyncDatasourceDynamodbConfig {
	var returns *AppsyncDatasourceDynamodbConfig
	_jsii_.Get(
		j,
		"dynamodbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) ElasticsearchConfig() AppsyncDatasourceElasticsearchConfigOutputReference {
	var returns AppsyncDatasourceElasticsearchConfigOutputReference
	_jsii_.Get(
		j,
		"elasticsearchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) ElasticsearchConfigInput() *AppsyncDatasourceElasticsearchConfig {
	var returns *AppsyncDatasourceElasticsearchConfig
	_jsii_.Get(
		j,
		"elasticsearchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) EventBridgeConfig() AppsyncDatasourceEventBridgeConfigOutputReference {
	var returns AppsyncDatasourceEventBridgeConfigOutputReference
	_jsii_.Get(
		j,
		"eventBridgeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) EventBridgeConfigInput() *AppsyncDatasourceEventBridgeConfig {
	var returns *AppsyncDatasourceEventBridgeConfig
	_jsii_.Get(
		j,
		"eventBridgeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) HttpConfig() AppsyncDatasourceHttpConfigOutputReference {
	var returns AppsyncDatasourceHttpConfigOutputReference
	_jsii_.Get(
		j,
		"httpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) HttpConfigInput() *AppsyncDatasourceHttpConfig {
	var returns *AppsyncDatasourceHttpConfig
	_jsii_.Get(
		j,
		"httpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) LambdaConfig() AppsyncDatasourceLambdaConfigOutputReference {
	var returns AppsyncDatasourceLambdaConfigOutputReference
	_jsii_.Get(
		j,
		"lambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) LambdaConfigInput() *AppsyncDatasourceLambdaConfig {
	var returns *AppsyncDatasourceLambdaConfig
	_jsii_.Get(
		j,
		"lambdaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) OpensearchserviceConfig() AppsyncDatasourceOpensearchserviceConfigOutputReference {
	var returns AppsyncDatasourceOpensearchserviceConfigOutputReference
	_jsii_.Get(
		j,
		"opensearchserviceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) OpensearchserviceConfigInput() *AppsyncDatasourceOpensearchserviceConfig {
	var returns *AppsyncDatasourceOpensearchserviceConfig
	_jsii_.Get(
		j,
		"opensearchserviceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) RelationalDatabaseConfig() AppsyncDatasourceRelationalDatabaseConfigOutputReference {
	var returns AppsyncDatasourceRelationalDatabaseConfigOutputReference
	_jsii_.Get(
		j,
		"relationalDatabaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) RelationalDatabaseConfigInput() *AppsyncDatasourceRelationalDatabaseConfig {
	var returns *AppsyncDatasourceRelationalDatabaseConfig
	_jsii_.Get(
		j,
		"relationalDatabaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appsync_datasource aws_appsync_datasource} Resource.
func NewAppsyncDatasource(scope constructs.Construct, id *string, config *AppsyncDatasourceConfig) AppsyncDatasource {
	_init_.Initialize()

	if err := validateNewAppsyncDatasourceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncDatasource{}

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncDatasource.AppsyncDatasource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appsync_datasource aws_appsync_datasource} Resource.
func NewAppsyncDatasource_Override(a AppsyncDatasource, scope constructs.Construct, id *string, config *AppsyncDatasourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncDatasource.AppsyncDatasource",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasource)SetApiId(val *string) {
	if err := j.validateSetApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource)SetServiceRoleArn(val *string) {
	if err := j.validateSetServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
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
func AppsyncDatasource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncDatasource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appsyncDatasource.AppsyncDatasource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncDatasource_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncDatasource_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appsyncDatasource.AppsyncDatasource",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncDatasource_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncDatasource_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appsyncDatasource.AppsyncDatasource",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncDatasource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.appsyncDatasource.AppsyncDatasource",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppsyncDatasource) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppsyncDatasource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasource) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasource) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasource) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasource) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasource) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncDatasource) PutDynamodbConfig(value *AppsyncDatasourceDynamodbConfig) {
	if err := a.validatePutDynamodbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynamodbConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasource) PutElasticsearchConfig(value *AppsyncDatasourceElasticsearchConfig) {
	if err := a.validatePutElasticsearchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElasticsearchConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasource) PutEventBridgeConfig(value *AppsyncDatasourceEventBridgeConfig) {
	if err := a.validatePutEventBridgeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventBridgeConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasource) PutHttpConfig(value *AppsyncDatasourceHttpConfig) {
	if err := a.validatePutHttpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasource) PutLambdaConfig(value *AppsyncDatasourceLambdaConfig) {
	if err := a.validatePutLambdaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasource) PutOpensearchserviceConfig(value *AppsyncDatasourceOpensearchserviceConfig) {
	if err := a.validatePutOpensearchserviceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpensearchserviceConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasource) PutRelationalDatabaseConfig(value *AppsyncDatasourceRelationalDatabaseConfig) {
	if err := a.validatePutRelationalDatabaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelationalDatabaseConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetDynamodbConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamodbConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetElasticsearchConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetEventBridgeConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEventBridgeConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetHttpConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetLambdaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetOpensearchserviceConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOpensearchserviceConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetRelationalDatabaseConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRelationalDatabaseConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetServiceRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

