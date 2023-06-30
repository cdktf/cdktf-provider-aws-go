package timestreamwritetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/timestreamwritetable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TimestreamwriteTableMagneticStoreWritePropertiesOutputReference interface {
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
	EnableMagneticStoreWrites() interface{}
	SetEnableMagneticStoreWrites(val interface{})
	EnableMagneticStoreWritesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *TimestreamwriteTableMagneticStoreWriteProperties
	SetInternalValue(val *TimestreamwriteTableMagneticStoreWriteProperties)
	MagneticStoreRejectedDataLocation() TimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationOutputReference
	MagneticStoreRejectedDataLocationInput() *TimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocation
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutMagneticStoreRejectedDataLocation(value *TimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocation)
	ResetEnableMagneticStoreWrites()
	ResetMagneticStoreRejectedDataLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TimestreamwriteTableMagneticStoreWritePropertiesOutputReference
type jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) EnableMagneticStoreWrites() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMagneticStoreWrites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) EnableMagneticStoreWritesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMagneticStoreWritesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) InternalValue() *TimestreamwriteTableMagneticStoreWriteProperties {
	var returns *TimestreamwriteTableMagneticStoreWriteProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) MagneticStoreRejectedDataLocation() TimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationOutputReference {
	var returns TimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationOutputReference
	_jsii_.Get(
		j,
		"magneticStoreRejectedDataLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) MagneticStoreRejectedDataLocationInput() *TimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocation {
	var returns *TimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocation
	_jsii_.Get(
		j,
		"magneticStoreRejectedDataLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTimestreamwriteTableMagneticStoreWritePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TimestreamwriteTableMagneticStoreWritePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewTimestreamwriteTableMagneticStoreWritePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.timestreamwriteTable.TimestreamwriteTableMagneticStoreWritePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTimestreamwriteTableMagneticStoreWritePropertiesOutputReference_Override(t TimestreamwriteTableMagneticStoreWritePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.timestreamwriteTable.TimestreamwriteTableMagneticStoreWritePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference)SetEnableMagneticStoreWrites(val interface{}) {
	if err := j.validateSetEnableMagneticStoreWritesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMagneticStoreWrites",
		val,
	)
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference)SetInternalValue(val *TimestreamwriteTableMagneticStoreWriteProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) PutMagneticStoreRejectedDataLocation(value *TimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocation) {
	if err := t.validatePutMagneticStoreRejectedDataLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMagneticStoreRejectedDataLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) ResetEnableMagneticStoreWrites() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableMagneticStoreWrites",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) ResetMagneticStoreRejectedDataLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetMagneticStoreRejectedDataLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableMagneticStoreWritePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

