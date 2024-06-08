// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datazoneenvironmentblueprintconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/datazoneenvironmentblueprintconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/datazone_environment_blueprint_configuration aws_datazone_environment_blueprint_configuration}.
type DatazoneEnvironmentBlueprintConfiguration interface {
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
	DomainId() *string
	SetDomainId(val *string)
	DomainIdInput() *string
	EnabledRegions() *[]*string
	SetEnabledRegions(val *[]*string)
	EnabledRegionsInput() *[]*string
	EnvironmentBlueprintId() *string
	SetEnvironmentBlueprintId(val *string)
	EnvironmentBlueprintIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManageAccessRoleArn() *string
	SetManageAccessRoleArn(val *string)
	ManageAccessRoleArnInput() *string
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
	ProvisioningRoleArn() *string
	SetProvisioningRoleArn(val *string)
	ProvisioningRoleArnInput() *string
	// Experimental.
	RawOverrides() interface{}
	RegionalParameters() interface{}
	SetRegionalParameters(val interface{})
	RegionalParametersInput() interface{}
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
	ResetManageAccessRoleArn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProvisioningRoleArn()
	ResetRegionalParameters()
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

// The jsii proxy struct for DatazoneEnvironmentBlueprintConfiguration
type jsiiProxy_DatazoneEnvironmentBlueprintConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) DomainIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) EnabledRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) EnabledRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) EnvironmentBlueprintId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentBlueprintId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) EnvironmentBlueprintIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentBlueprintIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ManageAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manageAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ManageAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manageAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ProvisioningRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ProvisioningRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) RegionalParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionalParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) RegionalParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionalParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/datazone_environment_blueprint_configuration aws_datazone_environment_blueprint_configuration} Resource.
func NewDatazoneEnvironmentBlueprintConfiguration(scope constructs.Construct, id *string, config *DatazoneEnvironmentBlueprintConfigurationConfig) DatazoneEnvironmentBlueprintConfiguration {
	_init_.Initialize()

	if err := validateNewDatazoneEnvironmentBlueprintConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazoneEnvironmentBlueprintConfiguration{}

	_jsii_.Create(
		"@cdktf/provider-aws.datazoneEnvironmentBlueprintConfiguration.DatazoneEnvironmentBlueprintConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/datazone_environment_blueprint_configuration aws_datazone_environment_blueprint_configuration} Resource.
func NewDatazoneEnvironmentBlueprintConfiguration_Override(d DatazoneEnvironmentBlueprintConfiguration, scope constructs.Construct, id *string, config *DatazoneEnvironmentBlueprintConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.datazoneEnvironmentBlueprintConfiguration.DatazoneEnvironmentBlueprintConfiguration",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration)SetDomainId(val *string) {
	if err := j.validateSetDomainIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainId",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration)SetEnabledRegions(val *[]*string) {
	if err := j.validateSetEnabledRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledRegions",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration)SetEnvironmentBlueprintId(val *string) {
	if err := j.validateSetEnvironmentBlueprintIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentBlueprintId",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration)SetManageAccessRoleArn(val *string) {
	if err := j.validateSetManageAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration)SetProvisioningRoleArn(val *string) {
	if err := j.validateSetProvisioningRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningRoleArn",
		val,
	)
}

func (j *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration)SetRegionalParameters(val interface{}) {
	if err := j.validateSetRegionalParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionalParameters",
		val,
	)
}

// Generates CDKTF code for importing a DatazoneEnvironmentBlueprintConfiguration resource upon running "cdktf plan <stack-name>".
func DatazoneEnvironmentBlueprintConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatazoneEnvironmentBlueprintConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.datazoneEnvironmentBlueprintConfiguration.DatazoneEnvironmentBlueprintConfiguration",
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
func DatazoneEnvironmentBlueprintConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatazoneEnvironmentBlueprintConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.datazoneEnvironmentBlueprintConfiguration.DatazoneEnvironmentBlueprintConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatazoneEnvironmentBlueprintConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatazoneEnvironmentBlueprintConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.datazoneEnvironmentBlueprintConfiguration.DatazoneEnvironmentBlueprintConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatazoneEnvironmentBlueprintConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatazoneEnvironmentBlueprintConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.datazoneEnvironmentBlueprintConfiguration.DatazoneEnvironmentBlueprintConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatazoneEnvironmentBlueprintConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.datazoneEnvironmentBlueprintConfiguration.DatazoneEnvironmentBlueprintConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ResetManageAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetManageAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ResetProvisioningRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetProvisioningRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ResetRegionalParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetRegionalParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneEnvironmentBlueprintConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

