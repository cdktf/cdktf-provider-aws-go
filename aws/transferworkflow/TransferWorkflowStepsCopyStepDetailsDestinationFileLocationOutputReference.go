package transferworkflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/transferworkflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference interface {
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
	EfsFileLocation() TransferWorkflowStepsCopyStepDetailsDestinationFileLocationEfsFileLocationOutputReference
	EfsFileLocationInput() *TransferWorkflowStepsCopyStepDetailsDestinationFileLocationEfsFileLocation
	// Experimental.
	Fqn() *string
	InternalValue() *TransferWorkflowStepsCopyStepDetailsDestinationFileLocation
	SetInternalValue(val *TransferWorkflowStepsCopyStepDetailsDestinationFileLocation)
	S3FileLocation() TransferWorkflowStepsCopyStepDetailsDestinationFileLocationS3FileLocationOutputReference
	S3FileLocationInput() *TransferWorkflowStepsCopyStepDetailsDestinationFileLocationS3FileLocation
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
	PutEfsFileLocation(value *TransferWorkflowStepsCopyStepDetailsDestinationFileLocationEfsFileLocation)
	PutS3FileLocation(value *TransferWorkflowStepsCopyStepDetailsDestinationFileLocationS3FileLocation)
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

// The jsii proxy struct for TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference
type jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) EfsFileLocation() TransferWorkflowStepsCopyStepDetailsDestinationFileLocationEfsFileLocationOutputReference {
	var returns TransferWorkflowStepsCopyStepDetailsDestinationFileLocationEfsFileLocationOutputReference
	_jsii_.Get(
		j,
		"efsFileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) EfsFileLocationInput() *TransferWorkflowStepsCopyStepDetailsDestinationFileLocationEfsFileLocation {
	var returns *TransferWorkflowStepsCopyStepDetailsDestinationFileLocationEfsFileLocation
	_jsii_.Get(
		j,
		"efsFileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) InternalValue() *TransferWorkflowStepsCopyStepDetailsDestinationFileLocation {
	var returns *TransferWorkflowStepsCopyStepDetailsDestinationFileLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) S3FileLocation() TransferWorkflowStepsCopyStepDetailsDestinationFileLocationS3FileLocationOutputReference {
	var returns TransferWorkflowStepsCopyStepDetailsDestinationFileLocationS3FileLocationOutputReference
	_jsii_.Get(
		j,
		"s3FileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) S3FileLocationInput() *TransferWorkflowStepsCopyStepDetailsDestinationFileLocationS3FileLocation {
	var returns *TransferWorkflowStepsCopyStepDetailsDestinationFileLocationS3FileLocation
	_jsii_.Get(
		j,
		"s3FileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference {
	_init_.Initialize()

	if err := validateNewTransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.transferWorkflow.TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference_Override(t TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.transferWorkflow.TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference)SetInternalValue(val *TransferWorkflowStepsCopyStepDetailsDestinationFileLocation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) PutEfsFileLocation(value *TransferWorkflowStepsCopyStepDetailsDestinationFileLocationEfsFileLocation) {
	if err := t.validatePutEfsFileLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEfsFileLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) PutS3FileLocation(value *TransferWorkflowStepsCopyStepDetailsDestinationFileLocationS3FileLocation) {
	if err := t.validatePutS3FileLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3FileLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) ResetEfsFileLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetEfsFileLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) ResetS3FileLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetS3FileLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (t *jsiiProxy_TransferWorkflowStepsCopyStepDetailsDestinationFileLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

