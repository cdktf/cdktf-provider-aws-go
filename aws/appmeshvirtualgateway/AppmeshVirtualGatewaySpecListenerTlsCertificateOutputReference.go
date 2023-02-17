package appmeshvirtualgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/appmeshvirtualgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference interface {
	cdktf.ComplexObject
	Acm() AppmeshVirtualGatewaySpecListenerTlsCertificateAcmOutputReference
	AcmInput() *AppmeshVirtualGatewaySpecListenerTlsCertificateAcm
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
	File() AppmeshVirtualGatewaySpecListenerTlsCertificateFileOutputReference
	FileInput() *AppmeshVirtualGatewaySpecListenerTlsCertificateFile
	// Experimental.
	Fqn() *string
	InternalValue() *AppmeshVirtualGatewaySpecListenerTlsCertificate
	SetInternalValue(val *AppmeshVirtualGatewaySpecListenerTlsCertificate)
	Sds() AppmeshVirtualGatewaySpecListenerTlsCertificateSdsOutputReference
	SdsInput() *AppmeshVirtualGatewaySpecListenerTlsCertificateSds
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
	PutAcm(value *AppmeshVirtualGatewaySpecListenerTlsCertificateAcm)
	PutFile(value *AppmeshVirtualGatewaySpecListenerTlsCertificateFile)
	PutSds(value *AppmeshVirtualGatewaySpecListenerTlsCertificateSds)
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

// The jsii proxy struct for AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference
type jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) Acm() AppmeshVirtualGatewaySpecListenerTlsCertificateAcmOutputReference {
	var returns AppmeshVirtualGatewaySpecListenerTlsCertificateAcmOutputReference
	_jsii_.Get(
		j,
		"acm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) AcmInput() *AppmeshVirtualGatewaySpecListenerTlsCertificateAcm {
	var returns *AppmeshVirtualGatewaySpecListenerTlsCertificateAcm
	_jsii_.Get(
		j,
		"acmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) File() AppmeshVirtualGatewaySpecListenerTlsCertificateFileOutputReference {
	var returns AppmeshVirtualGatewaySpecListenerTlsCertificateFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) FileInput() *AppmeshVirtualGatewaySpecListenerTlsCertificateFile {
	var returns *AppmeshVirtualGatewaySpecListenerTlsCertificateFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) InternalValue() *AppmeshVirtualGatewaySpecListenerTlsCertificate {
	var returns *AppmeshVirtualGatewaySpecListenerTlsCertificate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) Sds() AppmeshVirtualGatewaySpecListenerTlsCertificateSdsOutputReference {
	var returns AppmeshVirtualGatewaySpecListenerTlsCertificateSdsOutputReference
	_jsii_.Get(
		j,
		"sds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) SdsInput() *AppmeshVirtualGatewaySpecListenerTlsCertificateSds {
	var returns *AppmeshVirtualGatewaySpecListenerTlsCertificateSds
	_jsii_.Get(
		j,
		"sdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference {
	_init_.Initialize()

	if err := validateNewAppmeshVirtualGatewaySpecListenerTlsCertificateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualGateway.AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference_Override(a AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualGateway.AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference)SetInternalValue(val *AppmeshVirtualGatewaySpecListenerTlsCertificate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) PutAcm(value *AppmeshVirtualGatewaySpecListenerTlsCertificateAcm) {
	if err := a.validatePutAcmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAcm",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) PutFile(value *AppmeshVirtualGatewaySpecListenerTlsCertificateFile) {
	if err := a.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFile",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) PutSds(value *AppmeshVirtualGatewaySpecListenerTlsCertificateSds) {
	if err := a.validatePutSdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSds",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) ResetAcm() {
	_jsii_.InvokeVoid(
		a,
		"resetAcm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		a,
		"resetFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) ResetSds() {
	_jsii_.InvokeVoid(
		a,
		"resetSds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

