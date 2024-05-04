// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsreplicationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dmsreplicationconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/dms_replication_config aws_dms_replication_config}.
type DmsReplicationConfig interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputeConfig() DmsReplicationConfigComputeConfigOutputReference
	ComputeConfigInput() *DmsReplicationConfigComputeConfig
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	// Experimental.
	RawOverrides() interface{}
	ReplicationConfigIdentifier() *string
	SetReplicationConfigIdentifier(val *string)
	ReplicationConfigIdentifierInput() *string
	ReplicationSettings() *string
	SetReplicationSettings(val *string)
	ReplicationSettingsInput() *string
	ReplicationType() *string
	SetReplicationType(val *string)
	ReplicationTypeInput() *string
	ResourceIdentifier() *string
	SetResourceIdentifier(val *string)
	ResourceIdentifierInput() *string
	SourceEndpointArn() *string
	SetSourceEndpointArn(val *string)
	SourceEndpointArnInput() *string
	StartReplication() interface{}
	SetStartReplication(val interface{})
	StartReplicationInput() interface{}
	SupplementalSettings() *string
	SetSupplementalSettings(val *string)
	SupplementalSettingsInput() *string
	TableMappings() *string
	SetTableMappings(val *string)
	TableMappingsInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TargetEndpointArn() *string
	SetTargetEndpointArn(val *string)
	TargetEndpointArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DmsReplicationConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutComputeConfig(value *DmsReplicationConfigComputeConfig)
	PutTimeouts(value *DmsReplicationConfigTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReplicationSettings()
	ResetResourceIdentifier()
	ResetStartReplication()
	ResetSupplementalSettings()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
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

// The jsii proxy struct for DmsReplicationConfig
type jsiiProxy_DmsReplicationConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DmsReplicationConfig) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) ComputeConfig() DmsReplicationConfigComputeConfigOutputReference {
	var returns DmsReplicationConfigComputeConfigOutputReference
	_jsii_.Get(
		j,
		"computeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) ComputeConfigInput() *DmsReplicationConfigComputeConfig {
	var returns *DmsReplicationConfigComputeConfig
	_jsii_.Get(
		j,
		"computeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) ReplicationConfigIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationConfigIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) ReplicationConfigIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationConfigIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) ReplicationSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) ReplicationSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) ReplicationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) ReplicationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) ResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) ResourceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) SourceEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) SourceEndpointArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndpointArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) StartReplication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) StartReplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) SupplementalSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplementalSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) SupplementalSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplementalSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) TableMappings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) TableMappingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) TargetEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) TargetEndpointArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetEndpointArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) Timeouts() DmsReplicationConfigTimeoutsOutputReference {
	var returns DmsReplicationConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/dms_replication_config aws_dms_replication_config} Resource.
func NewDmsReplicationConfig(scope constructs.Construct, id *string, config *DmsReplicationConfigConfig) DmsReplicationConfig {
	_init_.Initialize()

	if err := validateNewDmsReplicationConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsReplicationConfig{}

	_jsii_.Create(
		"@cdktf/provider-aws.dmsReplicationConfig.DmsReplicationConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/dms_replication_config aws_dms_replication_config} Resource.
func NewDmsReplicationConfig_Override(d DmsReplicationConfig, scope constructs.Construct, id *string, config *DmsReplicationConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dmsReplicationConfig.DmsReplicationConfig",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetReplicationConfigIdentifier(val *string) {
	if err := j.validateSetReplicationConfigIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationConfigIdentifier",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetReplicationSettings(val *string) {
	if err := j.validateSetReplicationSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationSettings",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetReplicationType(val *string) {
	if err := j.validateSetReplicationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationType",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetResourceIdentifier(val *string) {
	if err := j.validateSetResourceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetSourceEndpointArn(val *string) {
	if err := j.validateSetSourceEndpointArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceEndpointArn",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetStartReplication(val interface{}) {
	if err := j.validateSetStartReplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startReplication",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetSupplementalSettings(val *string) {
	if err := j.validateSetSupplementalSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supplementalSettings",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetTableMappings(val *string) {
	if err := j.validateSetTableMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableMappings",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfig)SetTargetEndpointArn(val *string) {
	if err := j.validateSetTargetEndpointArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetEndpointArn",
		val,
	)
}

// Generates CDKTF code for importing a DmsReplicationConfig resource upon running "cdktf plan <stack-name>".
func DmsReplicationConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDmsReplicationConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dmsReplicationConfig.DmsReplicationConfig",
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
func DmsReplicationConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsReplicationConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dmsReplicationConfig.DmsReplicationConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsReplicationConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsReplicationConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dmsReplicationConfig.DmsReplicationConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsReplicationConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsReplicationConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dmsReplicationConfig.DmsReplicationConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsReplicationConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dmsReplicationConfig.DmsReplicationConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DmsReplicationConfig) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DmsReplicationConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DmsReplicationConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsReplicationConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DmsReplicationConfig) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsReplicationConfig) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DmsReplicationConfig) PutComputeConfig(value *DmsReplicationConfigComputeConfig) {
	if err := d.validatePutComputeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putComputeConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsReplicationConfig) PutTimeouts(value *DmsReplicationConfigTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsReplicationConfig) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfig) ResetReplicationSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicationSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfig) ResetResourceIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfig) ResetStartReplication() {
	_jsii_.InvokeVoid(
		d,
		"resetStartReplication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfig) ResetSupplementalSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetSupplementalSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfig) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfig) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

