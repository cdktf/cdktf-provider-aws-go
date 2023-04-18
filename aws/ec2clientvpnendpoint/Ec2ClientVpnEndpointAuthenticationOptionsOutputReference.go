package ec2clientvpnendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/ec2clientvpnendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2ClientVpnEndpointAuthenticationOptionsOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectoryId() *string
	SetActiveDirectoryId(val *string)
	ActiveDirectoryIdInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RootCertificateChainArn() *string
	SetRootCertificateChainArn(val *string)
	RootCertificateChainArnInput() *string
	SamlProviderArn() *string
	SetSamlProviderArn(val *string)
	SamlProviderArnInput() *string
	SelfServiceSamlProviderArn() *string
	SetSelfServiceSamlProviderArn(val *string)
	SelfServiceSamlProviderArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetActiveDirectoryId()
	ResetRootCertificateChainArn()
	ResetSamlProviderArn()
	ResetSelfServiceSamlProviderArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2ClientVpnEndpointAuthenticationOptionsOutputReference
type jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ActiveDirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ActiveDirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) RootCertificateChainArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootCertificateChainArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) RootCertificateChainArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootCertificateChainArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) SamlProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) SamlProviderArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlProviderArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) SelfServiceSamlProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServiceSamlProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) SelfServiceSamlProviderArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServiceSamlProviderArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewEc2ClientVpnEndpointAuthenticationOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2ClientVpnEndpointAuthenticationOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2ClientVpnEndpointAuthenticationOptionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthenticationOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2ClientVpnEndpointAuthenticationOptionsOutputReference_Override(e Ec2ClientVpnEndpointAuthenticationOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthenticationOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetActiveDirectoryId(val *string) {
	if err := j.validateSetActiveDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDirectoryId",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetRootCertificateChainArn(val *string) {
	if err := j.validateSetRootCertificateChainArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootCertificateChainArn",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetSamlProviderArn(val *string) {
	if err := j.validateSetSamlProviderArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlProviderArn",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetSelfServiceSamlProviderArn(val *string) {
	if err := j.validateSetSelfServiceSamlProviderArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfServiceSamlProviderArn",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ResetActiveDirectoryId() {
	_jsii_.InvokeVoid(
		e,
		"resetActiveDirectoryId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ResetRootCertificateChainArn() {
	_jsii_.InvokeVoid(
		e,
		"resetRootCertificateChainArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ResetSamlProviderArn() {
	_jsii_.InvokeVoid(
		e,
		"resetSamlProviderArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ResetSelfServiceSamlProviderArn() {
	_jsii_.InvokeVoid(
		e,
		"resetSelfServiceSamlProviderArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

