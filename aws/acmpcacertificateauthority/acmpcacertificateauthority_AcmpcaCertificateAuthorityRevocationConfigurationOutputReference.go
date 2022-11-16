package acmpcacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/acmpcacertificateauthority/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AcmpcaCertificateAuthorityRevocationConfigurationOutputReference interface {
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
	CrlConfiguration() AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference
	CrlConfigurationInput() *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration
	// Experimental.
	Fqn() *string
	InternalValue() *AcmpcaCertificateAuthorityRevocationConfiguration
	SetInternalValue(val *AcmpcaCertificateAuthorityRevocationConfiguration)
	OcspConfiguration() AcmpcaCertificateAuthorityRevocationConfigurationOcspConfigurationOutputReference
	OcspConfigurationInput() *AcmpcaCertificateAuthorityRevocationConfigurationOcspConfiguration
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
	PutCrlConfiguration(value *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration)
	PutOcspConfiguration(value *AcmpcaCertificateAuthorityRevocationConfigurationOcspConfiguration)
	ResetCrlConfiguration()
	ResetOcspConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AcmpcaCertificateAuthorityRevocationConfigurationOutputReference
type jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) CrlConfiguration() AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference {
	var returns AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference
	_jsii_.Get(
		j,
		"crlConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) CrlConfigurationInput() *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration {
	var returns *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration
	_jsii_.Get(
		j,
		"crlConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) InternalValue() *AcmpcaCertificateAuthorityRevocationConfiguration {
	var returns *AcmpcaCertificateAuthorityRevocationConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) OcspConfiguration() AcmpcaCertificateAuthorityRevocationConfigurationOcspConfigurationOutputReference {
	var returns AcmpcaCertificateAuthorityRevocationConfigurationOcspConfigurationOutputReference
	_jsii_.Get(
		j,
		"ocspConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) OcspConfigurationInput() *AcmpcaCertificateAuthorityRevocationConfigurationOcspConfiguration {
	var returns *AcmpcaCertificateAuthorityRevocationConfigurationOcspConfiguration
	_jsii_.Get(
		j,
		"ocspConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAcmpcaCertificateAuthorityRevocationConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AcmpcaCertificateAuthorityRevocationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewAcmpcaCertificateAuthorityRevocationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.acmpcaCertificateAuthority.AcmpcaCertificateAuthorityRevocationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAcmpcaCertificateAuthorityRevocationConfigurationOutputReference_Override(a AcmpcaCertificateAuthorityRevocationConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.acmpcaCertificateAuthority.AcmpcaCertificateAuthorityRevocationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference)SetInternalValue(val *AcmpcaCertificateAuthorityRevocationConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) PutCrlConfiguration(value *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration) {
	if err := a.validatePutCrlConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCrlConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) PutOcspConfiguration(value *AcmpcaCertificateAuthorityRevocationConfigurationOcspConfiguration) {
	if err := a.validatePutOcspConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOcspConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) ResetCrlConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCrlConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) ResetOcspConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetOcspConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

