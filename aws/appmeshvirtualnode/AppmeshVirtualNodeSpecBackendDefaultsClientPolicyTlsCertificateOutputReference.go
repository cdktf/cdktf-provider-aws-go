package appmeshvirtualnode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/appmeshvirtualnode/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference interface {
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
	File() AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFileOutputReference
	FileInput() *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFile
	// Experimental.
	Fqn() *string
	InternalValue() *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate
	SetInternalValue(val *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate)
	Sds() AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSdsOutputReference
	SdsInput() *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSds
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
	PutFile(value *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFile)
	PutSds(value *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSds)
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

// The jsii proxy struct for AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference
type jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) File() AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFileOutputReference {
	var returns AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) FileInput() *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFile {
	var returns *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) InternalValue() *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate {
	var returns *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) Sds() AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSdsOutputReference {
	var returns AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSdsOutputReference
	_jsii_.Get(
		j,
		"sds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) SdsInput() *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSds {
	var returns *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSds
	_jsii_.Get(
		j,
		"sdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference {
	_init_.Initialize()

	if err := validateNewAppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualNode.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference_Override(a AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualNode.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference)SetInternalValue(val *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) PutFile(value *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFile) {
	if err := a.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFile",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) PutSds(value *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSds) {
	if err := a.validatePutSdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSds",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		a,
		"resetFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) ResetSds() {
	_jsii_.InvokeVoid(
		a,
		"resetSds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

