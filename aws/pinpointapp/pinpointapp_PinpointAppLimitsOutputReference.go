package pinpointapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/pinpointapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PinpointAppLimitsOutputReference interface {
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
	Daily() *float64
	SetDaily(val *float64)
	DailyInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *PinpointAppLimits
	SetInternalValue(val *PinpointAppLimits)
	MaximumDuration() *float64
	SetMaximumDuration(val *float64)
	MaximumDurationInput() *float64
	MessagesPerSecond() *float64
	SetMessagesPerSecond(val *float64)
	MessagesPerSecondInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Total() *float64
	SetTotal(val *float64)
	TotalInput() *float64
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
	ResetDaily()
	ResetMaximumDuration()
	ResetMessagesPerSecond()
	ResetTotal()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PinpointAppLimitsOutputReference
type jsiiProxy_PinpointAppLimitsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) Daily() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) DailyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) InternalValue() *PinpointAppLimits {
	var returns *PinpointAppLimits
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) MaximumDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) MaximumDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) MessagesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messagesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) MessagesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messagesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) Total() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"total",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference) TotalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalInput",
		&returns,
	)
	return returns
}


func NewPinpointAppLimitsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PinpointAppLimitsOutputReference {
	_init_.Initialize()

	if err := validateNewPinpointAppLimitsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PinpointAppLimitsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.pinpointApp.PinpointAppLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPinpointAppLimitsOutputReference_Override(p PinpointAppLimitsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pinpointApp.PinpointAppLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference)SetDaily(val *float64) {
	if err := j.validateSetDailyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daily",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference)SetInternalValue(val *PinpointAppLimits) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference)SetMaximumDuration(val *float64) {
	if err := j.validateSetMaximumDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumDuration",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference)SetMessagesPerSecond(val *float64) {
	if err := j.validateSetMessagesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messagesPerSecond",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PinpointAppLimitsOutputReference)SetTotal(val *float64) {
	if err := j.validateSetTotalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"total",
		val,
	)
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) ResetDaily() {
	_jsii_.InvokeVoid(
		p,
		"resetDaily",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) ResetMaximumDuration() {
	_jsii_.InvokeVoid(
		p,
		"resetMaximumDuration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) ResetMessagesPerSecond() {
	_jsii_.InvokeVoid(
		p,
		"resetMessagesPerSecond",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) ResetTotal() {
	_jsii_.InvokeVoid(
		p,
		"resetTotal",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointAppLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

