package iot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/iot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotIndexingConfigurationThingIndexingConfigurationOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomField() IotIndexingConfigurationThingIndexingConfigurationCustomFieldList
	CustomFieldInput() interface{}
	DeviceDefenderIndexingMode() *string
	SetDeviceDefenderIndexingMode(val *string)
	DeviceDefenderIndexingModeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *IotIndexingConfigurationThingIndexingConfiguration
	SetInternalValue(val *IotIndexingConfigurationThingIndexingConfiguration)
	ManagedField() IotIndexingConfigurationThingIndexingConfigurationManagedFieldList
	ManagedFieldInput() interface{}
	NamedShadowIndexingMode() *string
	SetNamedShadowIndexingMode(val *string)
	NamedShadowIndexingModeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThingConnectivityIndexingMode() *string
	SetThingConnectivityIndexingMode(val *string)
	ThingConnectivityIndexingModeInput() *string
	ThingIndexingMode() *string
	SetThingIndexingMode(val *string)
	ThingIndexingModeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutCustomField(value interface{})
	PutManagedField(value interface{})
	ResetCustomField()
	ResetDeviceDefenderIndexingMode()
	ResetManagedField()
	ResetNamedShadowIndexingMode()
	ResetThingConnectivityIndexingMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotIndexingConfigurationThingIndexingConfigurationOutputReference
type jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) CustomField() IotIndexingConfigurationThingIndexingConfigurationCustomFieldList {
	var returns IotIndexingConfigurationThingIndexingConfigurationCustomFieldList
	_jsii_.Get(
		j,
		"customField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) CustomFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) DeviceDefenderIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceDefenderIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) DeviceDefenderIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceDefenderIndexingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) InternalValue() *IotIndexingConfigurationThingIndexingConfiguration {
	var returns *IotIndexingConfigurationThingIndexingConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ManagedField() IotIndexingConfigurationThingIndexingConfigurationManagedFieldList {
	var returns IotIndexingConfigurationThingIndexingConfigurationManagedFieldList
	_jsii_.Get(
		j,
		"managedField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ManagedFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) NamedShadowIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namedShadowIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) NamedShadowIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namedShadowIndexingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ThingConnectivityIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingConnectivityIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ThingConnectivityIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingConnectivityIndexingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ThingIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ThingIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingIndexingModeInput",
		&returns,
	)
	return returns
}


func NewIotIndexingConfigurationThingIndexingConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotIndexingConfigurationThingIndexingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewIotIndexingConfigurationThingIndexingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.iot.IotIndexingConfigurationThingIndexingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotIndexingConfigurationThingIndexingConfigurationOutputReference_Override(i IotIndexingConfigurationThingIndexingConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.iot.IotIndexingConfigurationThingIndexingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference)SetDeviceDefenderIndexingMode(val *string) {
	if err := j.validateSetDeviceDefenderIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceDefenderIndexingMode",
		val,
	)
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference)SetInternalValue(val *IotIndexingConfigurationThingIndexingConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference)SetNamedShadowIndexingMode(val *string) {
	if err := j.validateSetNamedShadowIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namedShadowIndexingMode",
		val,
	)
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference)SetThingConnectivityIndexingMode(val *string) {
	if err := j.validateSetThingConnectivityIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thingConnectivityIndexingMode",
		val,
	)
}

func (j *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference)SetThingIndexingMode(val *string) {
	if err := j.validateSetThingIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thingIndexingMode",
		val,
	)
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) PutCustomField(value interface{}) {
	if err := i.validatePutCustomFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCustomField",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) PutManagedField(value interface{}) {
	if err := i.validatePutManagedFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putManagedField",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ResetCustomField() {
	_jsii_.InvokeVoid(
		i,
		"resetCustomField",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ResetDeviceDefenderIndexingMode() {
	_jsii_.InvokeVoid(
		i,
		"resetDeviceDefenderIndexingMode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ResetManagedField() {
	_jsii_.InvokeVoid(
		i,
		"resetManagedField",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ResetNamedShadowIndexingMode() {
	_jsii_.InvokeVoid(
		i,
		"resetNamedShadowIndexingMode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ResetThingConnectivityIndexingMode() {
	_jsii_.InvokeVoid(
		i,
		"resetThingConnectivityIndexingMode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingIndexingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

