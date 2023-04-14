package transferserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v13/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v13/transferserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferServerProtocolDetailsOutputReference interface {
	cdktf.ComplexObject
	As2Transports() *[]*string
	SetAs2Transports(val *[]*string)
	As2TransportsInput() *[]*string
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
	InternalValue() *TransferServerProtocolDetails
	SetInternalValue(val *TransferServerProtocolDetails)
	PassiveIp() *string
	SetPassiveIp(val *string)
	PassiveIpInput() *string
	SetStatOption() *string
	SetSetStatOption(val *string)
	SetStatOptionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsSessionResumptionMode() *string
	SetTlsSessionResumptionMode(val *string)
	TlsSessionResumptionModeInput() *string
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
	ResetAs2Transports()
	ResetPassiveIp()
	ResetSetStatOption()
	ResetTlsSessionResumptionMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TransferServerProtocolDetailsOutputReference
type jsiiProxy_TransferServerProtocolDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) As2Transports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"as2Transports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) As2TransportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"as2TransportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) InternalValue() *TransferServerProtocolDetails {
	var returns *TransferServerProtocolDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) PassiveIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passiveIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) PassiveIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passiveIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) SetStatOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setStatOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) SetStatOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setStatOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) TlsSessionResumptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsSessionResumptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference) TlsSessionResumptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsSessionResumptionModeInput",
		&returns,
	)
	return returns
}


func NewTransferServerProtocolDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TransferServerProtocolDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewTransferServerProtocolDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransferServerProtocolDetailsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.transferServer.TransferServerProtocolDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTransferServerProtocolDetailsOutputReference_Override(t TransferServerProtocolDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.transferServer.TransferServerProtocolDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference)SetAs2Transports(val *[]*string) {
	if err := j.validateSetAs2TransportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"as2Transports",
		val,
	)
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference)SetInternalValue(val *TransferServerProtocolDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference)SetPassiveIp(val *string) {
	if err := j.validateSetPassiveIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passiveIp",
		val,
	)
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference)SetSetStatOption(val *string) {
	if err := j.validateSetSetStatOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setStatOption",
		val,
	)
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TransferServerProtocolDetailsOutputReference)SetTlsSessionResumptionMode(val *string) {
	if err := j.validateSetTlsSessionResumptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsSessionResumptionMode",
		val,
	)
}

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) ResetAs2Transports() {
	_jsii_.InvokeVoid(
		t,
		"resetAs2Transports",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) ResetPassiveIp() {
	_jsii_.InvokeVoid(
		t,
		"resetPassiveIp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) ResetSetStatOption() {
	_jsii_.InvokeVoid(
		t,
		"resetSetStatOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) ResetTlsSessionResumptionMode() {
	_jsii_.InvokeVoid(
		t,
		"resetTlsSessionResumptionMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (t *jsiiProxy_TransferServerProtocolDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

