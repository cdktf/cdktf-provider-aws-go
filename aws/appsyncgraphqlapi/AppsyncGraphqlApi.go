// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncgraphqlapi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/appsyncgraphqlapi/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/appsync_graphql_api aws_appsync_graphql_api}.
type AppsyncGraphqlApi interface {
	cdktf.TerraformResource
	AdditionalAuthenticationProvider() AppsyncGraphqlApiAdditionalAuthenticationProviderList
	AdditionalAuthenticationProviderInput() interface{}
	Arn() *string
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
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
	IntrospectionConfig() *string
	SetIntrospectionConfig(val *string)
	IntrospectionConfigInput() *string
	LambdaAuthorizerConfig() AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference
	LambdaAuthorizerConfigInput() *AppsyncGraphqlApiLambdaAuthorizerConfig
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogConfig() AppsyncGraphqlApiLogConfigOutputReference
	LogConfigInput() *AppsyncGraphqlApiLogConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OpenidConnectConfig() AppsyncGraphqlApiOpenidConnectConfigOutputReference
	OpenidConnectConfigInput() *AppsyncGraphqlApiOpenidConnectConfig
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QueryDepthLimit() *float64
	SetQueryDepthLimit(val *float64)
	QueryDepthLimitInput() *float64
	// Experimental.
	RawOverrides() interface{}
	ResolverCountLimit() *float64
	SetResolverCountLimit(val *float64)
	ResolverCountLimitInput() *float64
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Uris() cdktf.StringMap
	UserPoolConfig() AppsyncGraphqlApiUserPoolConfigOutputReference
	UserPoolConfigInput() *AppsyncGraphqlApiUserPoolConfig
	Visibility() *string
	SetVisibility(val *string)
	VisibilityInput() *string
	XrayEnabled() interface{}
	SetXrayEnabled(val interface{})
	XrayEnabledInput() interface{}
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
	PutAdditionalAuthenticationProvider(value interface{})
	PutLambdaAuthorizerConfig(value *AppsyncGraphqlApiLambdaAuthorizerConfig)
	PutLogConfig(value *AppsyncGraphqlApiLogConfig)
	PutOpenidConnectConfig(value *AppsyncGraphqlApiOpenidConnectConfig)
	PutUserPoolConfig(value *AppsyncGraphqlApiUserPoolConfig)
	ResetAdditionalAuthenticationProvider()
	ResetId()
	ResetIntrospectionConfig()
	ResetLambdaAuthorizerConfig()
	ResetLogConfig()
	ResetOpenidConnectConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQueryDepthLimit()
	ResetResolverCountLimit()
	ResetSchema()
	ResetTags()
	ResetTagsAll()
	ResetUserPoolConfig()
	ResetVisibility()
	ResetXrayEnabled()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AppsyncGraphqlApi
type jsiiProxy_AppsyncGraphqlApi struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsyncGraphqlApi) AdditionalAuthenticationProvider() AppsyncGraphqlApiAdditionalAuthenticationProviderList {
	var returns AppsyncGraphqlApiAdditionalAuthenticationProviderList
	_jsii_.Get(
		j,
		"additionalAuthenticationProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) AdditionalAuthenticationProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalAuthenticationProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) IntrospectionConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"introspectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) IntrospectionConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"introspectionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) LambdaAuthorizerConfig() AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference {
	var returns AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) LambdaAuthorizerConfigInput() *AppsyncGraphqlApiLambdaAuthorizerConfig {
	var returns *AppsyncGraphqlApiLambdaAuthorizerConfig
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) LogConfig() AppsyncGraphqlApiLogConfigOutputReference {
	var returns AppsyncGraphqlApiLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) LogConfigInput() *AppsyncGraphqlApiLogConfig {
	var returns *AppsyncGraphqlApiLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) OpenidConnectConfig() AppsyncGraphqlApiOpenidConnectConfigOutputReference {
	var returns AppsyncGraphqlApiOpenidConnectConfigOutputReference
	_jsii_.Get(
		j,
		"openidConnectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) OpenidConnectConfigInput() *AppsyncGraphqlApiOpenidConnectConfig {
	var returns *AppsyncGraphqlApiOpenidConnectConfig
	_jsii_.Get(
		j,
		"openidConnectConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) QueryDepthLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryDepthLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) QueryDepthLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryDepthLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) ResolverCountLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resolverCountLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) ResolverCountLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resolverCountLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Uris() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"uris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) UserPoolConfig() AppsyncGraphqlApiUserPoolConfigOutputReference {
	var returns AppsyncGraphqlApiUserPoolConfigOutputReference
	_jsii_.Get(
		j,
		"userPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) UserPoolConfigInput() *AppsyncGraphqlApiUserPoolConfig {
	var returns *AppsyncGraphqlApiUserPoolConfig
	_jsii_.Get(
		j,
		"userPoolConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) VisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) XrayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xrayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) XrayEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xrayEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/appsync_graphql_api aws_appsync_graphql_api} Resource.
func NewAppsyncGraphqlApi(scope constructs.Construct, id *string, config *AppsyncGraphqlApiConfig) AppsyncGraphqlApi {
	_init_.Initialize()

	if err := validateNewAppsyncGraphqlApiParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncGraphqlApi{}

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncGraphqlApi.AppsyncGraphqlApi",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/appsync_graphql_api aws_appsync_graphql_api} Resource.
func NewAppsyncGraphqlApi_Override(a AppsyncGraphqlApi, scope constructs.Construct, id *string, config *AppsyncGraphqlApiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncGraphqlApi.AppsyncGraphqlApi",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetIntrospectionConfig(val *string) {
	if err := j.validateSetIntrospectionConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"introspectionConfig",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetQueryDepthLimit(val *float64) {
	if err := j.validateSetQueryDepthLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryDepthLimit",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetResolverCountLimit(val *float64) {
	if err := j.validateSetResolverCountLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolverCountLimit",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetVisibility(val *string) {
	if err := j.validateSetVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi)SetXrayEnabled(val interface{}) {
	if err := j.validateSetXrayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xrayEnabled",
		val,
	)
}

// Generates CDKTF code for importing a AppsyncGraphqlApi resource upon running "cdktf plan <stack-name>".
func AppsyncGraphqlApi_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAppsyncGraphqlApi_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appsyncGraphqlApi.AppsyncGraphqlApi",
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
func AppsyncGraphqlApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncGraphqlApi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appsyncGraphqlApi.AppsyncGraphqlApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncGraphqlApi_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncGraphqlApi_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appsyncGraphqlApi.AppsyncGraphqlApi",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncGraphqlApi_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncGraphqlApi_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appsyncGraphqlApi.AppsyncGraphqlApi",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncGraphqlApi_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.appsyncGraphqlApi.AppsyncGraphqlApi",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApi) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncGraphqlApi) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncGraphqlApi) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncGraphqlApi) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncGraphqlApi) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncGraphqlApi) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncGraphqlApi) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncGraphqlApi) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncGraphqlApi) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncGraphqlApi) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApi) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncGraphqlApi) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) PutAdditionalAuthenticationProvider(value interface{}) {
	if err := a.validatePutAdditionalAuthenticationProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdditionalAuthenticationProvider",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) PutLambdaAuthorizerConfig(value *AppsyncGraphqlApiLambdaAuthorizerConfig) {
	if err := a.validatePutLambdaAuthorizerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaAuthorizerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) PutLogConfig(value *AppsyncGraphqlApiLogConfig) {
	if err := a.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) PutOpenidConnectConfig(value *AppsyncGraphqlApiOpenidConnectConfig) {
	if err := a.validatePutOpenidConnectConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpenidConnectConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) PutUserPoolConfig(value *AppsyncGraphqlApiUserPoolConfig) {
	if err := a.validatePutUserPoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserPoolConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetAdditionalAuthenticationProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalAuthenticationProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetIntrospectionConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetIntrospectionConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetLambdaAuthorizerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaAuthorizerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetLogConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetOpenidConnectConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOpenidConnectConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetQueryDepthLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryDepthLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetResolverCountLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetResolverCountLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetSchema() {
	_jsii_.InvokeVoid(
		a,
		"resetSchema",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetUserPoolConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetUserPoolConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetVisibility() {
	_jsii_.InvokeVoid(
		a,
		"resetVisibility",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetXrayEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetXrayEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApi) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApi) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApi) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApi) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

