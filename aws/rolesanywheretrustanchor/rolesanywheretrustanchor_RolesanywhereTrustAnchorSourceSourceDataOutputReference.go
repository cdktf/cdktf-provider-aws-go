package rolesanywheretrustanchor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/rolesanywheretrustanchor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RolesanywhereTrustAnchorSourceSourceDataOutputReference interface {
	cdktf.ComplexObject
	AcmPcaArn() *string
	SetAcmPcaArn(val *string)
	AcmPcaArnInput() *string
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
	InternalValue() *RolesanywhereTrustAnchorSourceSourceData
	SetInternalValue(val *RolesanywhereTrustAnchorSourceSourceData)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	X509CertificateData() *string
	SetX509CertificateData(val *string)
	X509CertificateDataInput() *string
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
	ResetAcmPcaArn()
	ResetX509CertificateData()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RolesanywhereTrustAnchorSourceSourceDataOutputReference
type jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) AcmPcaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmPcaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) AcmPcaArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmPcaArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) InternalValue() *RolesanywhereTrustAnchorSourceSourceData {
	var returns *RolesanywhereTrustAnchorSourceSourceData
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) X509CertificateData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"x509CertificateData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) X509CertificateDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"x509CertificateDataInput",
		&returns,
	)
	return returns
}


func NewRolesanywhereTrustAnchorSourceSourceDataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RolesanywhereTrustAnchorSourceSourceDataOutputReference {
	_init_.Initialize()

	if err := validateNewRolesanywhereTrustAnchorSourceSourceDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.rolesanywhereTrustAnchor.RolesanywhereTrustAnchorSourceSourceDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRolesanywhereTrustAnchorSourceSourceDataOutputReference_Override(r RolesanywhereTrustAnchorSourceSourceDataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.rolesanywhereTrustAnchor.RolesanywhereTrustAnchorSourceSourceDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference)SetAcmPcaArn(val *string) {
	if err := j.validateSetAcmPcaArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acmPcaArn",
		val,
	)
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference)SetInternalValue(val *RolesanywhereTrustAnchorSourceSourceData) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference)SetX509CertificateData(val *string) {
	if err := j.validateSetX509CertificateDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"x509CertificateData",
		val,
	)
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) ResetAcmPcaArn() {
	_jsii_.InvokeVoid(
		r,
		"resetAcmPcaArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) ResetX509CertificateData() {
	_jsii_.InvokeVoid(
		r,
		"resetX509CertificateData",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RolesanywhereTrustAnchorSourceSourceDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

