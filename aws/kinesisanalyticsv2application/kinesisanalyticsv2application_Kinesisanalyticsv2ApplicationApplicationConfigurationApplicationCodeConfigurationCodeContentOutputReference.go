package kinesisanalyticsv2application

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/kinesisanalyticsv2application/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContent
	SetInternalValue(val *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContent)
	S3ContentLocation() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationOutputReference
	S3ContentLocationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocation
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextContent() *string
	SetTextContent(val *string)
	TextContentInput() *string
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
	PutS3ContentLocation(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocation)
	ResetS3ContentLocation()
	ResetTextContent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference
type jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) InternalValue() *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContent {
	var returns *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContent
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) S3ContentLocation() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationOutputReference
	_jsii_.Get(
		j,
		"s3ContentLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) S3ContentLocationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocation {
	var returns *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocation
	_jsii_.Get(
		j,
		"s3ContentLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) TextContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) TextContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textContentInput",
		&returns,
	)
	return returns
}


func NewKinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisanalyticsv2Application.Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference_Override(k Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisanalyticsv2Application.Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference)SetInternalValue(val *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContent) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference)SetTextContent(val *string) {
	if err := j.validateSetTextContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textContent",
		val,
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) PutS3ContentLocation(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocation) {
	if err := k.validatePutS3ContentLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3ContentLocation",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) ResetS3ContentLocation() {
	_jsii_.InvokeVoid(
		k,
		"resetS3ContentLocation",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) ResetTextContent() {
	_jsii_.InvokeVoid(
		k,
		"resetTextContent",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

