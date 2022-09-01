package config

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/config/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigConfigurationRecorderRecordingGroupOutputReference interface {
	cdktf.ComplexObject
	AllSupported() interface{}
	SetAllSupported(val interface{})
	AllSupportedInput() interface{}
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
	// Experimental.
	Fqn() *string
	IncludeGlobalResourceTypes() interface{}
	SetIncludeGlobalResourceTypes(val interface{})
	IncludeGlobalResourceTypesInput() interface{}
	InternalValue() *ConfigConfigurationRecorderRecordingGroup
	SetInternalValue(val *ConfigConfigurationRecorderRecordingGroup)
	ResourceTypes() *[]*string
	SetResourceTypes(val *[]*string)
	ResourceTypesInput() *[]*string
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
	ResetAllSupported()
	ResetIncludeGlobalResourceTypes()
	ResetResourceTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigConfigurationRecorderRecordingGroupOutputReference
type jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) AllSupported() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) AllSupportedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allSupportedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) IncludeGlobalResourceTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalResourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) IncludeGlobalResourceTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalResourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) InternalValue() *ConfigConfigurationRecorderRecordingGroup {
	var returns *ConfigConfigurationRecorderRecordingGroup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigConfigurationRecorderRecordingGroupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConfigConfigurationRecorderRecordingGroupOutputReference {
	_init_.Initialize()

	if err := validateNewConfigConfigurationRecorderRecordingGroupOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.config.ConfigConfigurationRecorderRecordingGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConfigConfigurationRecorderRecordingGroupOutputReference_Override(c ConfigConfigurationRecorderRecordingGroupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.config.ConfigConfigurationRecorderRecordingGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetAllSupported(val interface{}) {
	if err := j.validateSetAllSupportedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allSupported",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetIncludeGlobalResourceTypes(val interface{}) {
	if err := j.validateSetIncludeGlobalResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeGlobalResourceTypes",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetInternalValue(val *ConfigConfigurationRecorderRecordingGroup) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetResourceTypes(val *[]*string) {
	if err := j.validateSetResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypes",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResetAllSupported() {
	_jsii_.InvokeVoid(
		c,
		"resetAllSupported",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResetIncludeGlobalResourceTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeGlobalResourceTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResetResourceTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

