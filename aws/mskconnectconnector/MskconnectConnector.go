// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskconnectconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/mskconnectconnector/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/mskconnect_connector aws_mskconnect_connector}.
type MskconnectConnector interface {
	cdktf.TerraformResource
	Arn() *string
	Capacity() MskconnectConnectorCapacityOutputReference
	CapacityInput() *MskconnectConnectorCapacity
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectorConfiguration() *map[string]*string
	SetConnectorConfiguration(val *map[string]*string)
	ConnectorConfigurationInput() *map[string]*string
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
	KafkaCluster() MskconnectConnectorKafkaClusterOutputReference
	KafkaClusterClientAuthentication() MskconnectConnectorKafkaClusterClientAuthenticationOutputReference
	KafkaClusterClientAuthenticationInput() *MskconnectConnectorKafkaClusterClientAuthentication
	KafkaClusterEncryptionInTransit() MskconnectConnectorKafkaClusterEncryptionInTransitOutputReference
	KafkaClusterEncryptionInTransitInput() *MskconnectConnectorKafkaClusterEncryptionInTransit
	KafkaClusterInput() *MskconnectConnectorKafkaCluster
	KafkaconnectVersion() *string
	SetKafkaconnectVersion(val *string)
	KafkaconnectVersionInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogDelivery() MskconnectConnectorLogDeliveryOutputReference
	LogDeliveryInput() *MskconnectConnectorLogDelivery
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Plugin() MskconnectConnectorPluginList
	PluginInput() interface{}
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
	ServiceExecutionRoleArn() *string
	SetServiceExecutionRoleArn(val *string)
	ServiceExecutionRoleArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MskconnectConnectorTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	WorkerConfiguration() MskconnectConnectorWorkerConfigurationOutputReference
	WorkerConfigurationInput() *MskconnectConnectorWorkerConfiguration
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
	PutCapacity(value *MskconnectConnectorCapacity)
	PutKafkaCluster(value *MskconnectConnectorKafkaCluster)
	PutKafkaClusterClientAuthentication(value *MskconnectConnectorKafkaClusterClientAuthentication)
	PutKafkaClusterEncryptionInTransit(value *MskconnectConnectorKafkaClusterEncryptionInTransit)
	PutLogDelivery(value *MskconnectConnectorLogDelivery)
	PutPlugin(value interface{})
	PutTimeouts(value *MskconnectConnectorTimeouts)
	PutWorkerConfiguration(value *MskconnectConnectorWorkerConfiguration)
	ResetDescription()
	ResetId()
	ResetLogDelivery()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetWorkerConfiguration()
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

// The jsii proxy struct for MskconnectConnector
type jsiiProxy_MskconnectConnector struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MskconnectConnector) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Capacity() MskconnectConnectorCapacityOutputReference {
	var returns MskconnectConnectorCapacityOutputReference
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) CapacityInput() *MskconnectConnectorCapacity {
	var returns *MskconnectConnectorCapacity
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) ConnectorConfiguration() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"connectorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) ConnectorConfigurationInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"connectorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) KafkaCluster() MskconnectConnectorKafkaClusterOutputReference {
	var returns MskconnectConnectorKafkaClusterOutputReference
	_jsii_.Get(
		j,
		"kafkaCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) KafkaClusterClientAuthentication() MskconnectConnectorKafkaClusterClientAuthenticationOutputReference {
	var returns MskconnectConnectorKafkaClusterClientAuthenticationOutputReference
	_jsii_.Get(
		j,
		"kafkaClusterClientAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) KafkaClusterClientAuthenticationInput() *MskconnectConnectorKafkaClusterClientAuthentication {
	var returns *MskconnectConnectorKafkaClusterClientAuthentication
	_jsii_.Get(
		j,
		"kafkaClusterClientAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) KafkaClusterEncryptionInTransit() MskconnectConnectorKafkaClusterEncryptionInTransitOutputReference {
	var returns MskconnectConnectorKafkaClusterEncryptionInTransitOutputReference
	_jsii_.Get(
		j,
		"kafkaClusterEncryptionInTransit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) KafkaClusterEncryptionInTransitInput() *MskconnectConnectorKafkaClusterEncryptionInTransit {
	var returns *MskconnectConnectorKafkaClusterEncryptionInTransit
	_jsii_.Get(
		j,
		"kafkaClusterEncryptionInTransitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) KafkaClusterInput() *MskconnectConnectorKafkaCluster {
	var returns *MskconnectConnectorKafkaCluster
	_jsii_.Get(
		j,
		"kafkaClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) KafkaconnectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaconnectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) KafkaconnectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaconnectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) LogDelivery() MskconnectConnectorLogDeliveryOutputReference {
	var returns MskconnectConnectorLogDeliveryOutputReference
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) LogDeliveryInput() *MskconnectConnectorLogDelivery {
	var returns *MskconnectConnectorLogDelivery
	_jsii_.Get(
		j,
		"logDeliveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Plugin() MskconnectConnectorPluginList {
	var returns MskconnectConnectorPluginList
	_jsii_.Get(
		j,
		"plugin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) PluginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pluginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) ServiceExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) ServiceExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Timeouts() MskconnectConnectorTimeoutsOutputReference {
	var returns MskconnectConnectorTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) WorkerConfiguration() MskconnectConnectorWorkerConfigurationOutputReference {
	var returns MskconnectConnectorWorkerConfigurationOutputReference
	_jsii_.Get(
		j,
		"workerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnector) WorkerConfigurationInput() *MskconnectConnectorWorkerConfiguration {
	var returns *MskconnectConnectorWorkerConfiguration
	_jsii_.Get(
		j,
		"workerConfigurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/mskconnect_connector aws_mskconnect_connector} Resource.
func NewMskconnectConnector(scope constructs.Construct, id *string, config *MskconnectConnectorConfig) MskconnectConnector {
	_init_.Initialize()

	if err := validateNewMskconnectConnectorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MskconnectConnector{}

	_jsii_.Create(
		"@cdktf/provider-aws.mskconnectConnector.MskconnectConnector",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/mskconnect_connector aws_mskconnect_connector} Resource.
func NewMskconnectConnector_Override(m MskconnectConnector, scope constructs.Construct, id *string, config *MskconnectConnectorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.mskconnectConnector.MskconnectConnector",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MskconnectConnector)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnector)SetConnectorConfiguration(val *map[string]*string) {
	if err := j.validateSetConnectorConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorConfiguration",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnector)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnector)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnector)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnector)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnector)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnector)SetKafkaconnectVersion(val *string) {
	if err := j.validateSetKafkaconnectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kafkaconnectVersion",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnector)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnector)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnector)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnector)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnector)SetServiceExecutionRoleArn(val *string) {
	if err := j.validateSetServiceExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceExecutionRoleArn",
		val,
	)
}

// Generates CDKTF code for importing a MskconnectConnector resource upon running "cdktf plan <stack-name>".
func MskconnectConnector_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMskconnectConnector_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.mskconnectConnector.MskconnectConnector",
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
func MskconnectConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMskconnectConnector_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.mskconnectConnector.MskconnectConnector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MskconnectConnector_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMskconnectConnector_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.mskconnectConnector.MskconnectConnector",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MskconnectConnector_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMskconnectConnector_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.mskconnectConnector.MskconnectConnector",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MskconnectConnector_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.mskconnectConnector.MskconnectConnector",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MskconnectConnector) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MskconnectConnector) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MskconnectConnector) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MskconnectConnector) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MskconnectConnector) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MskconnectConnector) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MskconnectConnector) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MskconnectConnector) PutCapacity(value *MskconnectConnectorCapacity) {
	if err := m.validatePutCapacityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCapacity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskconnectConnector) PutKafkaCluster(value *MskconnectConnectorKafkaCluster) {
	if err := m.validatePutKafkaClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putKafkaCluster",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskconnectConnector) PutKafkaClusterClientAuthentication(value *MskconnectConnectorKafkaClusterClientAuthentication) {
	if err := m.validatePutKafkaClusterClientAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putKafkaClusterClientAuthentication",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskconnectConnector) PutKafkaClusterEncryptionInTransit(value *MskconnectConnectorKafkaClusterEncryptionInTransit) {
	if err := m.validatePutKafkaClusterEncryptionInTransitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putKafkaClusterEncryptionInTransit",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskconnectConnector) PutLogDelivery(value *MskconnectConnectorLogDelivery) {
	if err := m.validatePutLogDeliveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLogDelivery",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskconnectConnector) PutPlugin(value interface{}) {
	if err := m.validatePutPluginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPlugin",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskconnectConnector) PutTimeouts(value *MskconnectConnectorTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskconnectConnector) PutWorkerConfiguration(value *MskconnectConnectorWorkerConfiguration) {
	if err := m.validatePutWorkerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWorkerConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskconnectConnector) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskconnectConnector) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskconnectConnector) ResetLogDelivery() {
	_jsii_.InvokeVoid(
		m,
		"resetLogDelivery",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskconnectConnector) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskconnectConnector) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskconnectConnector) ResetWorkerConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkerConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskconnectConnector) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnector) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

