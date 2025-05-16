// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/imagebuilderdistributionconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagebuilderDistributionConfigurationDistributionOutputReference interface {
	cdktf.ComplexObject
	AmiDistributionConfiguration() ImagebuilderDistributionConfigurationDistributionAmiDistributionConfigurationOutputReference
	AmiDistributionConfigurationInput() *ImagebuilderDistributionConfigurationDistributionAmiDistributionConfiguration
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
	ContainerDistributionConfiguration() ImagebuilderDistributionConfigurationDistributionContainerDistributionConfigurationOutputReference
	ContainerDistributionConfigurationInput() *ImagebuilderDistributionConfigurationDistributionContainerDistributionConfiguration
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FastLaunchConfiguration() ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationList
	FastLaunchConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LaunchTemplateConfiguration() ImagebuilderDistributionConfigurationDistributionLaunchTemplateConfigurationList
	LaunchTemplateConfigurationInput() interface{}
	LicenseConfigurationArns() *[]*string
	SetLicenseConfigurationArns(val *[]*string)
	LicenseConfigurationArnsInput() *[]*string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	S3ExportConfiguration() ImagebuilderDistributionConfigurationDistributionS3ExportConfigurationOutputReference
	S3ExportConfigurationInput() *ImagebuilderDistributionConfigurationDistributionS3ExportConfiguration
	SsmParameterConfiguration() ImagebuilderDistributionConfigurationDistributionSsmParameterConfigurationList
	SsmParameterConfigurationInput() interface{}
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
	PutAmiDistributionConfiguration(value *ImagebuilderDistributionConfigurationDistributionAmiDistributionConfiguration)
	PutContainerDistributionConfiguration(value *ImagebuilderDistributionConfigurationDistributionContainerDistributionConfiguration)
	PutFastLaunchConfiguration(value interface{})
	PutLaunchTemplateConfiguration(value interface{})
	PutS3ExportConfiguration(value *ImagebuilderDistributionConfigurationDistributionS3ExportConfiguration)
	PutSsmParameterConfiguration(value interface{})
	ResetAmiDistributionConfiguration()
	ResetContainerDistributionConfiguration()
	ResetFastLaunchConfiguration()
	ResetLaunchTemplateConfiguration()
	ResetLicenseConfigurationArns()
	ResetS3ExportConfiguration()
	ResetSsmParameterConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ImagebuilderDistributionConfigurationDistributionOutputReference
type jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) AmiDistributionConfiguration() ImagebuilderDistributionConfigurationDistributionAmiDistributionConfigurationOutputReference {
	var returns ImagebuilderDistributionConfigurationDistributionAmiDistributionConfigurationOutputReference
	_jsii_.Get(
		j,
		"amiDistributionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) AmiDistributionConfigurationInput() *ImagebuilderDistributionConfigurationDistributionAmiDistributionConfiguration {
	var returns *ImagebuilderDistributionConfigurationDistributionAmiDistributionConfiguration
	_jsii_.Get(
		j,
		"amiDistributionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) ContainerDistributionConfiguration() ImagebuilderDistributionConfigurationDistributionContainerDistributionConfigurationOutputReference {
	var returns ImagebuilderDistributionConfigurationDistributionContainerDistributionConfigurationOutputReference
	_jsii_.Get(
		j,
		"containerDistributionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) ContainerDistributionConfigurationInput() *ImagebuilderDistributionConfigurationDistributionContainerDistributionConfiguration {
	var returns *ImagebuilderDistributionConfigurationDistributionContainerDistributionConfiguration
	_jsii_.Get(
		j,
		"containerDistributionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) FastLaunchConfiguration() ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationList {
	var returns ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationList
	_jsii_.Get(
		j,
		"fastLaunchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) FastLaunchConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fastLaunchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) LaunchTemplateConfiguration() ImagebuilderDistributionConfigurationDistributionLaunchTemplateConfigurationList {
	var returns ImagebuilderDistributionConfigurationDistributionLaunchTemplateConfigurationList
	_jsii_.Get(
		j,
		"launchTemplateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) LaunchTemplateConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) LicenseConfigurationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenseConfigurationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) LicenseConfigurationArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenseConfigurationArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) S3ExportConfiguration() ImagebuilderDistributionConfigurationDistributionS3ExportConfigurationOutputReference {
	var returns ImagebuilderDistributionConfigurationDistributionS3ExportConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3ExportConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) S3ExportConfigurationInput() *ImagebuilderDistributionConfigurationDistributionS3ExportConfiguration {
	var returns *ImagebuilderDistributionConfigurationDistributionS3ExportConfiguration
	_jsii_.Get(
		j,
		"s3ExportConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) SsmParameterConfiguration() ImagebuilderDistributionConfigurationDistributionSsmParameterConfigurationList {
	var returns ImagebuilderDistributionConfigurationDistributionSsmParameterConfigurationList
	_jsii_.Get(
		j,
		"ssmParameterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) SsmParameterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssmParameterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewImagebuilderDistributionConfigurationDistributionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ImagebuilderDistributionConfigurationDistributionOutputReference {
	_init_.Initialize()

	if err := validateNewImagebuilderDistributionConfigurationDistributionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilderDistributionConfiguration.ImagebuilderDistributionConfigurationDistributionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewImagebuilderDistributionConfigurationDistributionOutputReference_Override(i ImagebuilderDistributionConfigurationDistributionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilderDistributionConfiguration.ImagebuilderDistributionConfigurationDistributionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference)SetLicenseConfigurationArns(val *[]*string) {
	if err := j.validateSetLicenseConfigurationArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseConfigurationArns",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) PutAmiDistributionConfiguration(value *ImagebuilderDistributionConfigurationDistributionAmiDistributionConfiguration) {
	if err := i.validatePutAmiDistributionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAmiDistributionConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) PutContainerDistributionConfiguration(value *ImagebuilderDistributionConfigurationDistributionContainerDistributionConfiguration) {
	if err := i.validatePutContainerDistributionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putContainerDistributionConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) PutFastLaunchConfiguration(value interface{}) {
	if err := i.validatePutFastLaunchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFastLaunchConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) PutLaunchTemplateConfiguration(value interface{}) {
	if err := i.validatePutLaunchTemplateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLaunchTemplateConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) PutS3ExportConfiguration(value *ImagebuilderDistributionConfigurationDistributionS3ExportConfiguration) {
	if err := i.validatePutS3ExportConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putS3ExportConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) PutSsmParameterConfiguration(value interface{}) {
	if err := i.validatePutSsmParameterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSsmParameterConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) ResetAmiDistributionConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetAmiDistributionConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) ResetContainerDistributionConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetContainerDistributionConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) ResetFastLaunchConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetFastLaunchConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) ResetLaunchTemplateConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetLaunchTemplateConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) ResetLicenseConfigurationArns() {
	_jsii_.InvokeVoid(
		i,
		"resetLicenseConfigurationArns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) ResetS3ExportConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetS3ExportConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) ResetSsmParameterConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetSsmParameterConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

