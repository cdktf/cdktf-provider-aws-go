package globalacceleratoraccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/globalacceleratoraccelerator/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlobalacceleratorAcceleratorAttributesOutputReference interface {
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
	InternalValue() *GlobalacceleratorAcceleratorAttributes
	SetInternalValue(val *GlobalacceleratorAcceleratorAttributes)
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

// The jsii proxy struct for GlobalacceleratorAcceleratorAttributesOutputReference
type jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) FlowLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flowLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) FlowLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flowLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) FlowLogsS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) FlowLogsS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) FlowLogsS3Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) FlowLogsS3PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogsS3PrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) InternalValue() *GlobalacceleratorAcceleratorAttributes {
	var returns *GlobalacceleratorAcceleratorAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlobalacceleratorAcceleratorAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GlobalacceleratorAcceleratorAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewGlobalacceleratorAcceleratorAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.globalacceleratorAccelerator.GlobalacceleratorAcceleratorAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlobalacceleratorAcceleratorAttributesOutputReference_Override(g GlobalacceleratorAcceleratorAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.globalacceleratorAccelerator.GlobalacceleratorAcceleratorAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference)SetFlowLogsEnabled(val interface{}) {
	if err := j.validateSetFlowLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference)SetFlowLogsS3Bucket(val *string) {
	if err := j.validateSetFlowLogsS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowLogsS3Bucket",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference)SetFlowLogsS3Prefix(val *string) {
	if err := j.validateSetFlowLogsS3PrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowLogsS3Prefix",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference)SetInternalValue(val *GlobalacceleratorAcceleratorAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) ResetFlowLogsEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetFlowLogsEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) ResetFlowLogsS3Bucket() {
	_jsii_.InvokeVoid(
		g,
		"resetFlowLogsS3Bucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) ResetFlowLogsS3Prefix() {
	_jsii_.InvokeVoid(
		g,
		"resetFlowLogsS3Prefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlobalacceleratorAcceleratorAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

