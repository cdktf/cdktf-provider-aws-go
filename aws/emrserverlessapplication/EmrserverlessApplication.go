// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/emrserverlessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/emrserverless_application aws_emrserverless_application}.
type EmrserverlessApplication interface {
	cdktf.TerraformResource
	Architecture() *string
	SetArchitecture(val *string)
	ArchitectureInput() *string
	Arn() *string
	AutoStartConfiguration() EmrserverlessApplicationAutoStartConfigurationOutputReference
	AutoStartConfigurationInput() *EmrserverlessApplicationAutoStartConfiguration
	AutoStopConfiguration() EmrserverlessApplicationAutoStopConfigurationOutputReference
	AutoStopConfigurationInput() *EmrserverlessApplicationAutoStopConfiguration
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
	ImageConfiguration() EmrserverlessApplicationImageConfigurationOutputReference
	ImageConfigurationInput() *EmrserverlessApplicationImageConfiguration
	InitialCapacity() EmrserverlessApplicationInitialCapacityList
	InitialCapacityInput() interface{}
	InteractiveConfiguration() EmrserverlessApplicationInteractiveConfigurationOutputReference
	InteractiveConfigurationInput() *EmrserverlessApplicationInteractiveConfiguration
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumCapacity() EmrserverlessApplicationMaximumCapacityOutputReference
	MaximumCapacityInput() *EmrserverlessApplicationMaximumCapacity
	MonitoringConfiguration() EmrserverlessApplicationMonitoringConfigurationOutputReference
	MonitoringConfigurationInput() *EmrserverlessApplicationMonitoringConfiguration
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfiguration() EmrserverlessApplicationNetworkConfigurationOutputReference
	NetworkConfigurationInput() *EmrserverlessApplicationNetworkConfiguration
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReleaseLabel() *string
	SetReleaseLabel(val *string)
	ReleaseLabelInput() *string
	RuntimeConfiguration() EmrserverlessApplicationRuntimeConfigurationList
	RuntimeConfigurationInput() interface{}
	SchedulerConfiguration() EmrserverlessApplicationSchedulerConfigurationOutputReference
	SchedulerConfigurationInput() *EmrserverlessApplicationSchedulerConfiguration
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
	PutAutoStartConfiguration(value *EmrserverlessApplicationAutoStartConfiguration)
	PutAutoStopConfiguration(value *EmrserverlessApplicationAutoStopConfiguration)
	PutImageConfiguration(value *EmrserverlessApplicationImageConfiguration)
	PutInitialCapacity(value interface{})
	PutInteractiveConfiguration(value *EmrserverlessApplicationInteractiveConfiguration)
	PutMaximumCapacity(value *EmrserverlessApplicationMaximumCapacity)
	PutMonitoringConfiguration(value *EmrserverlessApplicationMonitoringConfiguration)
	PutNetworkConfiguration(value *EmrserverlessApplicationNetworkConfiguration)
	PutRuntimeConfiguration(value interface{})
	PutSchedulerConfiguration(value *EmrserverlessApplicationSchedulerConfiguration)
	ResetArchitecture()
	ResetAutoStartConfiguration()
	ResetAutoStopConfiguration()
	ResetId()
	ResetImageConfiguration()
	ResetInitialCapacity()
	ResetInteractiveConfiguration()
	ResetMaximumCapacity()
	ResetMonitoringConfiguration()
	ResetNetworkConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetRuntimeConfiguration()
	ResetSchedulerConfiguration()
	ResetTags()
	ResetTagsAll()
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

// The jsii proxy struct for EmrserverlessApplication
type jsiiProxy_EmrserverlessApplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EmrserverlessApplication) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) AutoStartConfiguration() EmrserverlessApplicationAutoStartConfigurationOutputReference {
	var returns EmrserverlessApplicationAutoStartConfigurationOutputReference
	_jsii_.Get(
		j,
		"autoStartConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) AutoStartConfigurationInput() *EmrserverlessApplicationAutoStartConfiguration {
	var returns *EmrserverlessApplicationAutoStartConfiguration
	_jsii_.Get(
		j,
		"autoStartConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) AutoStopConfiguration() EmrserverlessApplicationAutoStopConfigurationOutputReference {
	var returns EmrserverlessApplicationAutoStopConfigurationOutputReference
	_jsii_.Get(
		j,
		"autoStopConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) AutoStopConfigurationInput() *EmrserverlessApplicationAutoStopConfiguration {
	var returns *EmrserverlessApplicationAutoStopConfiguration
	_jsii_.Get(
		j,
		"autoStopConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) ImageConfiguration() EmrserverlessApplicationImageConfigurationOutputReference {
	var returns EmrserverlessApplicationImageConfigurationOutputReference
	_jsii_.Get(
		j,
		"imageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) ImageConfigurationInput() *EmrserverlessApplicationImageConfiguration {
	var returns *EmrserverlessApplicationImageConfiguration
	_jsii_.Get(
		j,
		"imageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) InitialCapacity() EmrserverlessApplicationInitialCapacityList {
	var returns EmrserverlessApplicationInitialCapacityList
	_jsii_.Get(
		j,
		"initialCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) InitialCapacityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) InteractiveConfiguration() EmrserverlessApplicationInteractiveConfigurationOutputReference {
	var returns EmrserverlessApplicationInteractiveConfigurationOutputReference
	_jsii_.Get(
		j,
		"interactiveConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) InteractiveConfigurationInput() *EmrserverlessApplicationInteractiveConfiguration {
	var returns *EmrserverlessApplicationInteractiveConfiguration
	_jsii_.Get(
		j,
		"interactiveConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) MaximumCapacity() EmrserverlessApplicationMaximumCapacityOutputReference {
	var returns EmrserverlessApplicationMaximumCapacityOutputReference
	_jsii_.Get(
		j,
		"maximumCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) MaximumCapacityInput() *EmrserverlessApplicationMaximumCapacity {
	var returns *EmrserverlessApplicationMaximumCapacity
	_jsii_.Get(
		j,
		"maximumCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) MonitoringConfiguration() EmrserverlessApplicationMonitoringConfigurationOutputReference {
	var returns EmrserverlessApplicationMonitoringConfigurationOutputReference
	_jsii_.Get(
		j,
		"monitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) MonitoringConfigurationInput() *EmrserverlessApplicationMonitoringConfiguration {
	var returns *EmrserverlessApplicationMonitoringConfiguration
	_jsii_.Get(
		j,
		"monitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) NetworkConfiguration() EmrserverlessApplicationNetworkConfigurationOutputReference {
	var returns EmrserverlessApplicationNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) NetworkConfigurationInput() *EmrserverlessApplicationNetworkConfiguration {
	var returns *EmrserverlessApplicationNetworkConfiguration
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) ReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) RuntimeConfiguration() EmrserverlessApplicationRuntimeConfigurationList {
	var returns EmrserverlessApplicationRuntimeConfigurationList
	_jsii_.Get(
		j,
		"runtimeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) RuntimeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) SchedulerConfiguration() EmrserverlessApplicationSchedulerConfigurationOutputReference {
	var returns EmrserverlessApplicationSchedulerConfigurationOutputReference
	_jsii_.Get(
		j,
		"schedulerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) SchedulerConfigurationInput() *EmrserverlessApplicationSchedulerConfiguration {
	var returns *EmrserverlessApplicationSchedulerConfiguration
	_jsii_.Get(
		j,
		"schedulerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplication) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/emrserverless_application aws_emrserverless_application} Resource.
func NewEmrserverlessApplication(scope constructs.Construct, id *string, config *EmrserverlessApplicationConfig) EmrserverlessApplication {
	_init_.Initialize()

	if err := validateNewEmrserverlessApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrserverlessApplication{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrserverlessApplication.EmrserverlessApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/emrserverless_application aws_emrserverless_application} Resource.
func NewEmrserverlessApplication_Override(e EmrserverlessApplication, scope constructs.Construct, id *string, config *EmrserverlessApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrserverlessApplication.EmrserverlessApplication",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetArchitecture(val *string) {
	if err := j.validateSetArchitectureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetReleaseLabel(val *string) {
	if err := j.validateSetReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplication)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a EmrserverlessApplication resource upon running "cdktf plan <stack-name>".
func EmrserverlessApplication_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEmrserverlessApplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.emrserverlessApplication.EmrserverlessApplication",
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
func EmrserverlessApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrserverlessApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.emrserverlessApplication.EmrserverlessApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EmrserverlessApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrserverlessApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.emrserverlessApplication.EmrserverlessApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EmrserverlessApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrserverlessApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.emrserverlessApplication.EmrserverlessApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EmrserverlessApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.emrserverlessApplication.EmrserverlessApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) PutAutoStartConfiguration(value *EmrserverlessApplicationAutoStartConfiguration) {
	if err := e.validatePutAutoStartConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutoStartConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) PutAutoStopConfiguration(value *EmrserverlessApplicationAutoStopConfiguration) {
	if err := e.validatePutAutoStopConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutoStopConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) PutImageConfiguration(value *EmrserverlessApplicationImageConfiguration) {
	if err := e.validatePutImageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putImageConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) PutInitialCapacity(value interface{}) {
	if err := e.validatePutInitialCapacityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInitialCapacity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) PutInteractiveConfiguration(value *EmrserverlessApplicationInteractiveConfiguration) {
	if err := e.validatePutInteractiveConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInteractiveConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) PutMaximumCapacity(value *EmrserverlessApplicationMaximumCapacity) {
	if err := e.validatePutMaximumCapacityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMaximumCapacity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) PutMonitoringConfiguration(value *EmrserverlessApplicationMonitoringConfiguration) {
	if err := e.validatePutMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) PutNetworkConfiguration(value *EmrserverlessApplicationNetworkConfiguration) {
	if err := e.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) PutRuntimeConfiguration(value interface{}) {
	if err := e.validatePutRuntimeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRuntimeConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) PutSchedulerConfiguration(value *EmrserverlessApplicationSchedulerConfiguration) {
	if err := e.validatePutSchedulerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSchedulerConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetArchitecture() {
	_jsii_.InvokeVoid(
		e,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetAutoStartConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoStartConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetAutoStopConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoStopConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetImageConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetImageConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetInitialCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetInitialCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetInteractiveConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetInteractiveConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetMaximumCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetMaximumCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetMonitoringConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetRuntimeConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetRuntimeConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetSchedulerConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetSchedulerConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

