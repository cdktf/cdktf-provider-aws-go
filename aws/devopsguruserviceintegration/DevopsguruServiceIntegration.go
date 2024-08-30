// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devopsguruserviceintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/devopsguruserviceintegration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/devopsguru_service_integration aws_devopsguru_service_integration}.
type DevopsguruServiceIntegration interface {
	cdktf.TerraformResource
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
	KmsServerSideEncryption() DevopsguruServiceIntegrationKmsServerSideEncryptionList
	KmsServerSideEncryptionInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogsAnomalyDetection() DevopsguruServiceIntegrationLogsAnomalyDetectionList
	LogsAnomalyDetectionInput() interface{}
	// The tree node.
	Node() constructs.Node
	OpsCenter() DevopsguruServiceIntegrationOpsCenterList
	OpsCenterInput() interface{}
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutKmsServerSideEncryption(value interface{})
	PutLogsAnomalyDetection(value interface{})
	PutOpsCenter(value interface{})
	ResetKmsServerSideEncryption()
	ResetLogsAnomalyDetection()
	ResetOpsCenter()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DevopsguruServiceIntegration
type jsiiProxy_DevopsguruServiceIntegration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DevopsguruServiceIntegration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) KmsServerSideEncryption() DevopsguruServiceIntegrationKmsServerSideEncryptionList {
	var returns DevopsguruServiceIntegrationKmsServerSideEncryptionList
	_jsii_.Get(
		j,
		"kmsServerSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) KmsServerSideEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsServerSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) LogsAnomalyDetection() DevopsguruServiceIntegrationLogsAnomalyDetectionList {
	var returns DevopsguruServiceIntegrationLogsAnomalyDetectionList
	_jsii_.Get(
		j,
		"logsAnomalyDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) LogsAnomalyDetectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logsAnomalyDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) OpsCenter() DevopsguruServiceIntegrationOpsCenterList {
	var returns DevopsguruServiceIntegrationOpsCenterList
	_jsii_.Get(
		j,
		"opsCenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) OpsCenterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opsCenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevopsguruServiceIntegration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/devopsguru_service_integration aws_devopsguru_service_integration} Resource.
func NewDevopsguruServiceIntegration(scope constructs.Construct, id *string, config *DevopsguruServiceIntegrationConfig) DevopsguruServiceIntegration {
	_init_.Initialize()

	if err := validateNewDevopsguruServiceIntegrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DevopsguruServiceIntegration{}

	_jsii_.Create(
		"@cdktf/provider-aws.devopsguruServiceIntegration.DevopsguruServiceIntegration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/devopsguru_service_integration aws_devopsguru_service_integration} Resource.
func NewDevopsguruServiceIntegration_Override(d DevopsguruServiceIntegration, scope constructs.Construct, id *string, config *DevopsguruServiceIntegrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.devopsguruServiceIntegration.DevopsguruServiceIntegration",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DevopsguruServiceIntegration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DevopsguruServiceIntegration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DevopsguruServiceIntegration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DevopsguruServiceIntegration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DevopsguruServiceIntegration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DevopsguruServiceIntegration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DevopsguruServiceIntegration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a DevopsguruServiceIntegration resource upon running "cdktf plan <stack-name>".
func DevopsguruServiceIntegration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDevopsguruServiceIntegration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.devopsguruServiceIntegration.DevopsguruServiceIntegration",
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
func DevopsguruServiceIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevopsguruServiceIntegration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.devopsguruServiceIntegration.DevopsguruServiceIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DevopsguruServiceIntegration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevopsguruServiceIntegration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.devopsguruServiceIntegration.DevopsguruServiceIntegration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DevopsguruServiceIntegration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevopsguruServiceIntegration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.devopsguruServiceIntegration.DevopsguruServiceIntegration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DevopsguruServiceIntegration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.devopsguruServiceIntegration.DevopsguruServiceIntegration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DevopsguruServiceIntegration) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DevopsguruServiceIntegration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DevopsguruServiceIntegration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DevopsguruServiceIntegration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DevopsguruServiceIntegration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DevopsguruServiceIntegration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DevopsguruServiceIntegration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DevopsguruServiceIntegration) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DevopsguruServiceIntegration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DevopsguruServiceIntegration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsguruServiceIntegration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DevopsguruServiceIntegration) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) PutKmsServerSideEncryption(value interface{}) {
	if err := d.validatePutKmsServerSideEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKmsServerSideEncryption",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) PutLogsAnomalyDetection(value interface{}) {
	if err := d.validatePutLogsAnomalyDetectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLogsAnomalyDetection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) PutOpsCenter(value interface{}) {
	if err := d.validatePutOpsCenterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOpsCenter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) ResetKmsServerSideEncryption() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsServerSideEncryption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) ResetLogsAnomalyDetection() {
	_jsii_.InvokeVoid(
		d,
		"resetLogsAnomalyDetection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) ResetOpsCenter() {
	_jsii_.InvokeVoid(
		d,
		"resetOpsCenter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevopsguruServiceIntegration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsguruServiceIntegration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsguruServiceIntegration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsguruServiceIntegration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsguruServiceIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevopsguruServiceIntegration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

