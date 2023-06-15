package connectqueue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v15/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v15/connectqueue/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectQueueOutboundCallerConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ConnectQueueOutboundCallerConfig
	SetInternalValue(val *ConnectQueueOutboundCallerConfig)
	OutboundCallerIdName() *string
	SetOutboundCallerIdName(val *string)
	OutboundCallerIdNameInput() *string
	OutboundCallerIdNumberId() *string
	SetOutboundCallerIdNumberId(val *string)
	OutboundCallerIdNumberIdInput() *string
	OutboundFlowId() *string
	SetOutboundFlowId(val *string)
	OutboundFlowIdInput() *string
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
	ResetOutboundCallerIdName()
	ResetOutboundCallerIdNumberId()
	ResetOutboundFlowId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectQueueOutboundCallerConfigOutputReference
type jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) InternalValue() *ConnectQueueOutboundCallerConfig {
	var returns *ConnectQueueOutboundCallerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) OutboundCallerIdName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCallerIdName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) OutboundCallerIdNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCallerIdNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) OutboundCallerIdNumberId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCallerIdNumberId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) OutboundCallerIdNumberIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCallerIdNumberIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) OutboundFlowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundFlowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) OutboundFlowIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundFlowIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectQueueOutboundCallerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConnectQueueOutboundCallerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConnectQueueOutboundCallerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.connectQueue.ConnectQueueOutboundCallerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectQueueOutboundCallerConfigOutputReference_Override(c ConnectQueueOutboundCallerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.connectQueue.ConnectQueueOutboundCallerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetInternalValue(val *ConnectQueueOutboundCallerConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetOutboundCallerIdName(val *string) {
	if err := j.validateSetOutboundCallerIdNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundCallerIdName",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetOutboundCallerIdNumberId(val *string) {
	if err := j.validateSetOutboundCallerIdNumberIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundCallerIdNumberId",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetOutboundFlowId(val *string) {
	if err := j.validateSetOutboundFlowIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundFlowId",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ResetOutboundCallerIdName() {
	_jsii_.InvokeVoid(
		c,
		"resetOutboundCallerIdName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ResetOutboundCallerIdNumberId() {
	_jsii_.InvokeVoid(
		c,
		"resetOutboundCallerIdNumberId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ResetOutboundFlowId() {
	_jsii_.InvokeVoid(
		c,
		"resetOutboundFlowId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

