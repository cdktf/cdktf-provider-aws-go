package dataawsacmpcacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/dataawsacmpcacertificateauthority/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference interface {
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
	CrlConfiguration() DataAwsAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationList
	CrlConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OcspConfiguration() DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOcspConfigurationList
	OcspConfigurationInput() interface{}
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
	PutCrlConfiguration(value interface{})
	PutOcspConfiguration(value interface{})
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

// The jsii proxy struct for DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference
type jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) CrlConfiguration() DataAwsAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationList {
	var returns DataAwsAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationList
	_jsii_.Get(
		j,
		"crlConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) CrlConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crlConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) OcspConfiguration() DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOcspConfigurationList {
	var returns DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOcspConfigurationList
	_jsii_.Get(
		j,
		"ocspConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) OcspConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocspConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsAcmpcaCertificateAuthority.DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference_Override(d DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsAcmpcaCertificateAuthority.DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) PutCrlConfiguration(value interface{}) {
	if err := d.validatePutCrlConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCrlConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) PutOcspConfiguration(value interface{}) {
	if err := d.validatePutOcspConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOcspConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) ResetCrlConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetCrlConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) ResetOcspConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetOcspConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthorityRevocationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
