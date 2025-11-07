// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/fisexperimenttemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FisExperimentTemplateExperimentReportConfigurationOutputReference interface {
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
	DataSources() FisExperimentTemplateExperimentReportConfigurationDataSourcesOutputReference
	DataSourcesInput() *FisExperimentTemplateExperimentReportConfigurationDataSources
	// Experimental.
	Fqn() *string
	InternalValue() *FisExperimentTemplateExperimentReportConfiguration
	SetInternalValue(val *FisExperimentTemplateExperimentReportConfiguration)
	Outputs() FisExperimentTemplateExperimentReportConfigurationOutputsOutputReference
	OutputsInput() *FisExperimentTemplateExperimentReportConfigurationOutputs
	PostExperimentDuration() *string
	SetPostExperimentDuration(val *string)
	PostExperimentDurationInput() *string
	PreExperimentDuration() *string
	SetPreExperimentDuration(val *string)
	PreExperimentDurationInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutDataSources(value *FisExperimentTemplateExperimentReportConfigurationDataSources)
	PutOutputs(value *FisExperimentTemplateExperimentReportConfigurationOutputs)
	ResetDataSources()
	ResetOutputs()
	ResetPostExperimentDuration()
	ResetPreExperimentDuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FisExperimentTemplateExperimentReportConfigurationOutputReference
type jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) DataSources() FisExperimentTemplateExperimentReportConfigurationDataSourcesOutputReference {
	var returns FisExperimentTemplateExperimentReportConfigurationDataSourcesOutputReference
	_jsii_.Get(
		j,
		"dataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) DataSourcesInput() *FisExperimentTemplateExperimentReportConfigurationDataSources {
	var returns *FisExperimentTemplateExperimentReportConfigurationDataSources
	_jsii_.Get(
		j,
		"dataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) InternalValue() *FisExperimentTemplateExperimentReportConfiguration {
	var returns *FisExperimentTemplateExperimentReportConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) Outputs() FisExperimentTemplateExperimentReportConfigurationOutputsOutputReference {
	var returns FisExperimentTemplateExperimentReportConfigurationOutputsOutputReference
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) OutputsInput() *FisExperimentTemplateExperimentReportConfigurationOutputs {
	var returns *FisExperimentTemplateExperimentReportConfigurationOutputs
	_jsii_.Get(
		j,
		"outputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) PostExperimentDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postExperimentDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) PostExperimentDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postExperimentDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) PreExperimentDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preExperimentDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) PreExperimentDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preExperimentDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFisExperimentTemplateExperimentReportConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FisExperimentTemplateExperimentReportConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFisExperimentTemplateExperimentReportConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.fisExperimentTemplate.FisExperimentTemplateExperimentReportConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFisExperimentTemplateExperimentReportConfigurationOutputReference_Override(f FisExperimentTemplateExperimentReportConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fisExperimentTemplate.FisExperimentTemplateExperimentReportConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference)SetInternalValue(val *FisExperimentTemplateExperimentReportConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference)SetPostExperimentDuration(val *string) {
	if err := j.validateSetPostExperimentDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postExperimentDuration",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference)SetPreExperimentDuration(val *string) {
	if err := j.validateSetPreExperimentDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preExperimentDuration",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) PutDataSources(value *FisExperimentTemplateExperimentReportConfigurationDataSources) {
	if err := f.validatePutDataSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putDataSources",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) PutOutputs(value *FisExperimentTemplateExperimentReportConfigurationOutputs) {
	if err := f.validatePutOutputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putOutputs",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) ResetDataSources() {
	_jsii_.InvokeVoid(
		f,
		"resetDataSources",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) ResetOutputs() {
	_jsii_.InvokeVoid(
		f,
		"resetOutputs",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) ResetPostExperimentDuration() {
	_jsii_.InvokeVoid(
		f,
		"resetPostExperimentDuration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) ResetPreExperimentDuration() {
	_jsii_.InvokeVoid(
		f,
		"resetPreExperimentDuration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateExperimentReportConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

