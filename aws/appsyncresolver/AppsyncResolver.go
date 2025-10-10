// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncresolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appsyncresolver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_resolver aws_appsync_resolver}.
type AppsyncResolver interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	Arn() *string
	CachingConfig() AppsyncResolverCachingConfigOutputReference
	CachingConfigInput() *AppsyncResolverCachingConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Code() *string
	SetCode(val *string)
	CodeInput() *string
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
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Field() *string
	SetField(val *string)
	FieldInput() *string
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
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxBatchSize() *float64
	SetMaxBatchSize(val *float64)
	MaxBatchSizeInput() *float64
	// The tree node.
	Node() constructs.Node
	PipelineConfig() AppsyncResolverPipelineConfigOutputReference
	PipelineConfigInput() *AppsyncResolverPipelineConfig
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RequestTemplate() *string
	SetRequestTemplate(val *string)
	RequestTemplateInput() *string
	ResponseTemplate() *string
	SetResponseTemplate(val *string)
	ResponseTemplateInput() *string
	Runtime() AppsyncResolverRuntimeOutputReference
	RuntimeInput() *AppsyncResolverRuntime
	SyncConfig() AppsyncResolverSyncConfigOutputReference
	SyncConfigInput() *AppsyncResolverSyncConfig
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
	PutCachingConfig(value *AppsyncResolverCachingConfig)
	PutPipelineConfig(value *AppsyncResolverPipelineConfig)
	PutRuntime(value *AppsyncResolverRuntime)
	PutSyncConfig(value *AppsyncResolverSyncConfig)
	ResetCachingConfig()
	ResetCode()
	ResetDataSource()
	ResetId()
	ResetKind()
	ResetMaxBatchSize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPipelineConfig()
	ResetRegion()
	ResetRequestTemplate()
	ResetResponseTemplate()
	ResetRuntime()
	ResetSyncConfig()
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

// The jsii proxy struct for AppsyncResolver
type jsiiProxy_AppsyncResolver struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsyncResolver) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) CachingConfig() AppsyncResolverCachingConfigOutputReference {
	var returns AppsyncResolverCachingConfigOutputReference
	_jsii_.Get(
		j,
		"cachingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) CachingConfigInput() *AppsyncResolverCachingConfig {
	var returns *AppsyncResolverCachingConfig
	_jsii_.Get(
		j,
		"cachingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Code() *string {
	var returns *string
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) CodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) MaxBatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) MaxBatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) PipelineConfig() AppsyncResolverPipelineConfigOutputReference {
	var returns AppsyncResolverPipelineConfigOutputReference
	_jsii_.Get(
		j,
		"pipelineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) PipelineConfigInput() *AppsyncResolverPipelineConfig {
	var returns *AppsyncResolverPipelineConfig
	_jsii_.Get(
		j,
		"pipelineConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) RequestTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) RequestTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) ResponseTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) ResponseTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Runtime() AppsyncResolverRuntimeOutputReference {
	var returns AppsyncResolverRuntimeOutputReference
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) RuntimeInput() *AppsyncResolverRuntime {
	var returns *AppsyncResolverRuntime
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) SyncConfig() AppsyncResolverSyncConfigOutputReference {
	var returns AppsyncResolverSyncConfigOutputReference
	_jsii_.Get(
		j,
		"syncConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) SyncConfigInput() *AppsyncResolverSyncConfig {
	var returns *AppsyncResolverSyncConfig
	_jsii_.Get(
		j,
		"syncConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_resolver aws_appsync_resolver} Resource.
func NewAppsyncResolver(scope constructs.Construct, id *string, config *AppsyncResolverConfig) AppsyncResolver {
	_init_.Initialize()

	if err := validateNewAppsyncResolverParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncResolver{}

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncResolver.AppsyncResolver",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_resolver aws_appsync_resolver} Resource.
func NewAppsyncResolver_Override(a AppsyncResolver, scope constructs.Construct, id *string, config *AppsyncResolverConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncResolver.AppsyncResolver",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetApiId(val *string) {
	if err := j.validateSetApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetCode(val *string) {
	if err := j.validateSetCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetMaxBatchSize(val *float64) {
	if err := j.validateSetMaxBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBatchSize",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetRequestTemplate(val *string) {
	if err := j.validateSetRequestTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTemplate",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetResponseTemplate(val *string) {
	if err := j.validateSetResponseTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseTemplate",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a AppsyncResolver resource upon running "cdktf plan <stack-name>".
func AppsyncResolver_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAppsyncResolver_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appsyncResolver.AppsyncResolver",
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
func AppsyncResolver_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncResolver_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appsyncResolver.AppsyncResolver",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncResolver_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncResolver_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appsyncResolver.AppsyncResolver",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncResolver_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncResolver_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appsyncResolver.AppsyncResolver",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncResolver_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.appsyncResolver.AppsyncResolver",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppsyncResolver) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppsyncResolver) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppsyncResolver) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncResolver) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncResolver) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncResolver) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncResolver) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncResolver) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncResolver) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncResolver) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncResolver) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncResolver) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncResolver) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppsyncResolver) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncResolver) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppsyncResolver) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppsyncResolver) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppsyncResolver) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncResolver) PutCachingConfig(value *AppsyncResolverCachingConfig) {
	if err := a.validatePutCachingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCachingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncResolver) PutPipelineConfig(value *AppsyncResolverPipelineConfig) {
	if err := a.validatePutPipelineConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPipelineConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncResolver) PutRuntime(value *AppsyncResolverRuntime) {
	if err := a.validatePutRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRuntime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncResolver) PutSyncConfig(value *AppsyncResolverSyncConfig) {
	if err := a.validatePutSyncConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSyncConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetCachingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCachingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetCode() {
	_jsii_.InvokeVoid(
		a,
		"resetCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetDataSource() {
	_jsii_.InvokeVoid(
		a,
		"resetDataSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetKind() {
	_jsii_.InvokeVoid(
		a,
		"resetKind",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetMaxBatchSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxBatchSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetPipelineConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPipelineConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetRequestTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetResponseTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetRuntime() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetSyncConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSyncConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncResolver) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncResolver) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncResolver) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncResolver) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncResolver) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

