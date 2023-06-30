package globalacceleratorcustomroutingaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/globalacceleratorcustomroutingaccelerator/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference interface {
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
	FlowLogsEnabled() interface{}
	SetFlowLogsEnabled(val interface{})
	FlowLogsEnabledInput() interface{}
	FlowLogsS3Bucket() *string
	SetFlowLogsS3Bucket(val *string)
	FlowLogsS3BucketInput() *string
	FlowLogsS3Prefix() *string
	SetFlowLogsS3Prefix(val *string)
	FlowLogsS3PrefixInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GlobalacceleratorCustomRoutingAcceleratorAttributes
	SetInternalValue(val *GlobalacceleratorCustomRoutingAcceleratorAttributes)
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
	ResetFlowLogsEnabled()
	ResetFlowLogsS3Bucket()
	ResetFlowLogsS3Prefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference
type jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) FlowLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flowLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) FlowLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flowLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) FlowLogsS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) FlowLogsS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) FlowLogsS3Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) FlowLogsS3PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3PrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) InternalValue() *GlobalacceleratorCustomRoutingAcceleratorAttributes {
	var returns *GlobalacceleratorCustomRoutingAcceleratorAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewGlobalacceleratorCustomRoutingAcceleratorAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.globalacceleratorCustomRoutingAccelerator.GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference_Override(g GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.globalacceleratorCustomRoutingAccelerator.GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference)SetFlowLogsEnabled(val interface{}) {
	if err := j.validateSetFlowLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference)SetFlowLogsS3Bucket(val *string) {
	if err := j.validateSetFlowLogsS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowLogsS3Bucket",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference)SetFlowLogsS3Prefix(val *string) {
	if err := j.validateSetFlowLogsS3PrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowLogsS3Prefix",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference)SetInternalValue(val *GlobalacceleratorCustomRoutingAcceleratorAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) ResetFlowLogsEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetFlowLogsEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) ResetFlowLogsS3Bucket() {
	_jsii_.InvokeVoid(
		g,
		"resetFlowLogsS3Bucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) ResetFlowLogsS3Prefix() {
	_jsii_.InvokeVoid(
		g,
		"resetFlowLogsS3Prefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorCustomRoutingAcceleratorAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

