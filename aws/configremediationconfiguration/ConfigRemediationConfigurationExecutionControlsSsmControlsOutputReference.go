package configremediationconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/configremediationconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference interface {
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
	ConcurrentExecutionRatePercentage() *float64
	SetConcurrentExecutionRatePercentage(val *float64)
	ConcurrentExecutionRatePercentageInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ErrorPercentage() *float64
	SetErrorPercentage(val *float64)
	ErrorPercentageInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *ConfigRemediationConfigurationExecutionControlsSsmControls
	SetInternalValue(val *ConfigRemediationConfigurationExecutionControlsSsmControls)
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
	ResetConcurrentExecutionRatePercentage()
	ResetErrorPercentage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference
type jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) ConcurrentExecutionRatePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrentExecutionRatePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) ConcurrentExecutionRatePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrentExecutionRatePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) ErrorPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) ErrorPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) InternalValue() *ConfigRemediationConfigurationExecutionControlsSsmControls {
	var returns *ConfigRemediationConfigurationExecutionControlsSsmControls
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference {
	_init_.Initialize()

	if err := validateNewConfigRemediationConfigurationExecutionControlsSsmControlsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.configRemediationConfiguration.ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference_Override(c ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.configRemediationConfiguration.ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference)SetConcurrentExecutionRatePercentage(val *float64) {
	if err := j.validateSetConcurrentExecutionRatePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"concurrentExecutionRatePercentage",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference)SetErrorPercentage(val *float64) {
	if err := j.validateSetErrorPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorPercentage",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference)SetInternalValue(val *ConfigRemediationConfigurationExecutionControlsSsmControls) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) ResetConcurrentExecutionRatePercentage() {
	_jsii_.InvokeVoid(
		c,
		"resetConcurrentExecutionRatePercentage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) ResetErrorPercentage() {
	_jsii_.InvokeVoid(
		c,
		"resetErrorPercentage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConfigRemediationConfigurationExecutionControlsSsmControlsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

