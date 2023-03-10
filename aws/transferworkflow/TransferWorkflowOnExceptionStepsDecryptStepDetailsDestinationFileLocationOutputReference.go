package transferworkflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/transferworkflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference interface {
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
	EfsFileLocation() TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocationOutputReference
	EfsFileLocationInput() *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocation
	// Experimental.
	Fqn() *string
	InternalValue() *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocation
	SetInternalValue(val *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocation)
	S3FileLocation() TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocationOutputReference
	S3FileLocationInput() *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocation
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
	PutEfsFileLocation(value *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocation)
	PutS3FileLocation(value *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocation)
	ResetEfsFileLocation()
	ResetS3FileLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference
type jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) EfsFileLocation() TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocationOutputReference {
	var returns TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocationOutputReference
	_jsii_.Get(
		j,
		"efsFileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) EfsFileLocationInput() *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocation {
	var returns *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocation
	_jsii_.Get(
		j,
		"efsFileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) InternalValue() *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocation {
	var returns *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) S3FileLocation() TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocationOutputReference {
	var returns TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocationOutputReference
	_jsii_.Get(
		j,
		"s3FileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) S3FileLocationInput() *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocation {
	var returns *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocation
	_jsii_.Get(
		j,
		"s3FileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference {
	_init_.Initialize()

	if err := validateNewTransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.transferWorkflow.TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference_Override(t TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.transferWorkflow.TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference)SetInternalValue(val *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) PutEfsFileLocation(value *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocation) {
	if err := t.validatePutEfsFileLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEfsFileLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) PutS3FileLocation(value *TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocation) {
	if err := t.validatePutS3FileLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3FileLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) ResetEfsFileLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetEfsFileLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) ResetS3FileLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetS3FileLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsDecryptStepDetailsDestinationFileLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

