package kinesisanalyticsapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/kinesisanalyticsapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference interface {
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
	InternalValue() *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormat
	SetInternalValue(val *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormat)
	MappingParameters() KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParametersOutputReference
	MappingParametersInput() *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParameters
	RecordFormatType() *string
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
	PutMappingParameters(value *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParameters)
	ResetMappingParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference
type jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) InternalValue() *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormat {
	var returns *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormat
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) MappingParameters() KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParametersOutputReference {
	var returns KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParametersOutputReference
	_jsii_.Get(
		j,
		"mappingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) MappingParametersInput() *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParameters {
	var returns *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParameters
	_jsii_.Get(
		j,
		"mappingParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) RecordFormatType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordFormatType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisAnalyticsApplication.KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference_Override(k KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisAnalyticsApplication.KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference)SetInternalValue(val *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormat) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) PutMappingParameters(value *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParameters) {
	if err := k.validatePutMappingParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putMappingParameters",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) ResetMappingParameters() {
	_jsii_.InvokeVoid(
		k,
		"resetMappingParameters",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

