package emrcontainersjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/emrcontainersjobtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference interface {
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
	EntryPoint() *string
	SetEntryPoint(val *string)
	EntryPointInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriver
	SetInternalValue(val *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriver)
	SparkSqlParameters() *string
	SetSparkSqlParameters(val *string)
	SparkSqlParametersInput() *string
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
	ResetEntryPoint()
	ResetSparkSqlParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference
type jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) EntryPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) EntryPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) InternalValue() *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriver {
	var returns *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriver
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) SparkSqlParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkSqlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) SparkSqlParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkSqlParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference {
	_init_.Initialize()

	if err := validateNewEmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrcontainersJobTemplate.EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference_Override(e EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrcontainersJobTemplate.EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference)SetEntryPoint(val *string) {
	if err := j.validateSetEntryPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference)SetInternalValue(val *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriver) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference)SetSparkSqlParameters(val *string) {
	if err := j.validateSetSparkSqlParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkSqlParameters",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) ResetEntryPoint() {
	_jsii_.InvokeVoid(
		e,
		"resetEntryPoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) ResetSparkSqlParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetSparkSqlParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

