package fisexperimenttemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/fisexperimenttemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FisExperimentTemplateLogConfigurationOutputReference interface {
	cdktf.ComplexObject
	CloudwatchLogsConfiguration() FisExperimentTemplateLogConfigurationCloudwatchLogsConfigurationOutputReference
	CloudwatchLogsConfigurationInput() *FisExperimentTemplateLogConfigurationCloudwatchLogsConfiguration
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
	InternalValue() *FisExperimentTemplateLogConfiguration
	SetInternalValue(val *FisExperimentTemplateLogConfiguration)
	LogSchemaVersion() *float64
	SetLogSchemaVersion(val *float64)
	LogSchemaVersionInput() *float64
	S3Configuration() FisExperimentTemplateLogConfigurationS3ConfigurationOutputReference
	S3ConfigurationInput() *FisExperimentTemplateLogConfigurationS3Configuration
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
	PutCloudwatchLogsConfiguration(value *FisExperimentTemplateLogConfigurationCloudwatchLogsConfiguration)
	PutS3Configuration(value *FisExperimentTemplateLogConfigurationS3Configuration)
	ResetCloudwatchLogsConfiguration()
	ResetS3Configuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FisExperimentTemplateLogConfigurationOutputReference
type jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) CloudwatchLogsConfiguration() FisExperimentTemplateLogConfigurationCloudwatchLogsConfigurationOutputReference {
	var returns FisExperimentTemplateLogConfigurationCloudwatchLogsConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) CloudwatchLogsConfigurationInput() *FisExperimentTemplateLogConfigurationCloudwatchLogsConfiguration {
	var returns *FisExperimentTemplateLogConfigurationCloudwatchLogsConfiguration
	_jsii_.Get(
		j,
		"cloudwatchLogsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) InternalValue() *FisExperimentTemplateLogConfiguration {
	var returns *FisExperimentTemplateLogConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) LogSchemaVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logSchemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) LogSchemaVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logSchemaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) S3Configuration() FisExperimentTemplateLogConfigurationS3ConfigurationOutputReference {
	var returns FisExperimentTemplateLogConfigurationS3ConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) S3ConfigurationInput() *FisExperimentTemplateLogConfigurationS3Configuration {
	var returns *FisExperimentTemplateLogConfigurationS3Configuration
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFisExperimentTemplateLogConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FisExperimentTemplateLogConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFisExperimentTemplateLogConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.fisExperimentTemplate.FisExperimentTemplateLogConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFisExperimentTemplateLogConfigurationOutputReference_Override(f FisExperimentTemplateLogConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fisExperimentTemplate.FisExperimentTemplateLogConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference)SetInternalValue(val *FisExperimentTemplateLogConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference)SetLogSchemaVersion(val *float64) {
	if err := j.validateSetLogSchemaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logSchemaVersion",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) PutCloudwatchLogsConfiguration(value *FisExperimentTemplateLogConfigurationCloudwatchLogsConfiguration) {
	if err := f.validatePutCloudwatchLogsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putCloudwatchLogsConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) PutS3Configuration(value *FisExperimentTemplateLogConfigurationS3Configuration) {
	if err := f.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) ResetCloudwatchLogsConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetCloudwatchLogsConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) ResetS3Configuration() {
	_jsii_.InvokeVoid(
		f,
		"resetS3Configuration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FisExperimentTemplateLogConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

