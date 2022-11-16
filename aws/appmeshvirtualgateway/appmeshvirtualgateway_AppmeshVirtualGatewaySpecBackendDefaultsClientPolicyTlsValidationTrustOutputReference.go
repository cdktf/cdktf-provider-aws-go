package appmeshvirtualgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/appmeshvirtualgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference interface {
	cdktf.ComplexObject
	Acm() AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcmOutputReference
	AcmInput() *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcm
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
	File() AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFileOutputReference
	FileInput() *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFile
	// Experimental.
	Fqn() *string
	InternalValue() *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrust
	SetInternalValue(val *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrust)
	Sds() AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSdsOutputReference
	SdsInput() *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSds
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
	PutAcm(value *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcm)
	PutFile(value *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFile)
	PutSds(value *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSds)
	ResetAcm()
	ResetFile()
	ResetSds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference
type jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) Acm() AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcmOutputReference {
	var returns AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcmOutputReference
	_jsii_.Get(
		j,
		"acm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) AcmInput() *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcm {
	var returns *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcm
	_jsii_.Get(
		j,
		"acmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) File() AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFileOutputReference {
	var returns AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) FileInput() *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFile {
	var returns *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) InternalValue() *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrust {
	var returns *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrust
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) Sds() AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSdsOutputReference {
	var returns AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSdsOutputReference
	_jsii_.Get(
		j,
		"sds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) SdsInput() *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSds {
	var returns *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSds
	_jsii_.Get(
		j,
		"sdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference {
	_init_.Initialize()

	if err := validateNewAppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualGateway.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference_Override(a AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualGateway.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference)SetInternalValue(val *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrust) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) PutAcm(value *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcm) {
	if err := a.validatePutAcmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAcm",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) PutFile(value *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFile) {
	if err := a.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFile",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) PutSds(value *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSds) {
	if err := a.validatePutSdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSds",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) ResetAcm() {
	_jsii_.InvokeVoid(
		a,
		"resetAcm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		a,
		"resetFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) ResetSds() {
	_jsii_.InvokeVoid(
		a,
		"resetSds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

