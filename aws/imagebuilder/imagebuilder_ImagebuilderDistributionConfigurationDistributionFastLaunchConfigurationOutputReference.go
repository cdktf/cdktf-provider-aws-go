package imagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/imagebuilder/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference interface {
	cdktf.ComplexObject
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LaunchTemplate() ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationLaunchTemplateOutputReference
	LaunchTemplateInput() *ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationLaunchTemplate
	MaxParallelLaunches() *float64
	SetMaxParallelLaunches(val *float64)
	MaxParallelLaunchesInput() *float64
	SnapshotConfiguration() ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationSnapshotConfigurationOutputReference
	SnapshotConfigurationInput() *ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationSnapshotConfiguration
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
	PutLaunchTemplate(value *ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationLaunchTemplate)
	PutSnapshotConfiguration(value *ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationSnapshotConfiguration)
	ResetLaunchTemplate()
	ResetMaxParallelLaunches()
	ResetSnapshotConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference
type jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) LaunchTemplate() ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationLaunchTemplateOutputReference {
	var returns ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) LaunchTemplateInput() *ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationLaunchTemplate {
	var returns *ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationLaunchTemplate
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) MaxParallelLaunches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelLaunches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) MaxParallelLaunchesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelLaunchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) SnapshotConfiguration() ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationSnapshotConfigurationOutputReference {
	var returns ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationSnapshotConfigurationOutputReference
	_jsii_.Get(
		j,
		"snapshotConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) SnapshotConfigurationInput() *ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationSnapshotConfiguration {
	var returns *ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationSnapshotConfiguration
	_jsii_.Get(
		j,
		"snapshotConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilder.ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference_Override(i ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilder.ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference)SetMaxParallelLaunches(val *float64) {
	if err := j.validateSetMaxParallelLaunchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelLaunches",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) PutLaunchTemplate(value *ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationLaunchTemplate) {
	if err := i.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) PutSnapshotConfiguration(value *ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationSnapshotConfiguration) {
	if err := i.validatePutSnapshotConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSnapshotConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		i,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) ResetMaxParallelLaunches() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxParallelLaunches",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) ResetSnapshotConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetSnapshotConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

