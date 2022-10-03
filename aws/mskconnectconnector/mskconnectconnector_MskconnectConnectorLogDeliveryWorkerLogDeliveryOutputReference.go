package mskconnectconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/mskconnectconnector/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference interface {
	cdktf.ComplexObject
	CloudwatchLogs() MskconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogsOutputReference
	CloudwatchLogsInput() *MskconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs
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
	Firehose() MskconnectConnectorLogDeliveryWorkerLogDeliveryFirehoseOutputReference
	FirehoseInput() *MskconnectConnectorLogDeliveryWorkerLogDeliveryFirehose
	// Experimental.
	Fqn() *string
	InternalValue() *MskconnectConnectorLogDeliveryWorkerLogDelivery
	SetInternalValue(val *MskconnectConnectorLogDeliveryWorkerLogDelivery)
	S3() MskconnectConnectorLogDeliveryWorkerLogDeliveryS3OutputReference
	S3Input() *MskconnectConnectorLogDeliveryWorkerLogDeliveryS3
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
	PutCloudwatchLogs(value *MskconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs)
	PutFirehose(value *MskconnectConnectorLogDeliveryWorkerLogDeliveryFirehose)
	PutS3(value *MskconnectConnectorLogDeliveryWorkerLogDeliveryS3)
	ResetCloudwatchLogs()
	ResetFirehose()
	ResetS3()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference
type jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) CloudwatchLogs() MskconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogsOutputReference {
	var returns MskconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) CloudwatchLogsInput() *MskconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs {
	var returns *MskconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) Firehose() MskconnectConnectorLogDeliveryWorkerLogDeliveryFirehoseOutputReference {
	var returns MskconnectConnectorLogDeliveryWorkerLogDeliveryFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) FirehoseInput() *MskconnectConnectorLogDeliveryWorkerLogDeliveryFirehose {
	var returns *MskconnectConnectorLogDeliveryWorkerLogDeliveryFirehose
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) InternalValue() *MskconnectConnectorLogDeliveryWorkerLogDelivery {
	var returns *MskconnectConnectorLogDeliveryWorkerLogDelivery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) S3() MskconnectConnectorLogDeliveryWorkerLogDeliveryS3OutputReference {
	var returns MskconnectConnectorLogDeliveryWorkerLogDeliveryS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) S3Input() *MskconnectConnectorLogDeliveryWorkerLogDeliveryS3 {
	var returns *MskconnectConnectorLogDeliveryWorkerLogDeliveryS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference {
	_init_.Initialize()

	if err := validateNewMskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.mskconnectConnector.MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference_Override(m MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.mskconnectConnector.MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference)SetInternalValue(val *MskconnectConnectorLogDeliveryWorkerLogDelivery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) PutCloudwatchLogs(value *MskconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs) {
	if err := m.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) PutFirehose(value *MskconnectConnectorLogDeliveryWorkerLogDeliveryFirehose) {
	if err := m.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFirehose",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) PutS3(value *MskconnectConnectorLogDeliveryWorkerLogDeliveryS3) {
	if err := m.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putS3",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		m,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		m,
		"resetFirehose",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		m,
		"resetS3",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

