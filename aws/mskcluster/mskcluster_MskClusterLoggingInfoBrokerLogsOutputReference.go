package mskcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/mskcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MskClusterLoggingInfoBrokerLogsOutputReference interface {
	cdktf.ComplexObject
	CloudwatchLogs() MskClusterLoggingInfoBrokerLogsCloudwatchLogsOutputReference
	CloudwatchLogsInput() *MskClusterLoggingInfoBrokerLogsCloudwatchLogs
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
	Firehose() MskClusterLoggingInfoBrokerLogsFirehoseOutputReference
	FirehoseInput() *MskClusterLoggingInfoBrokerLogsFirehose
	// Experimental.
	Fqn() *string
	InternalValue() *MskClusterLoggingInfoBrokerLogs
	SetInternalValue(val *MskClusterLoggingInfoBrokerLogs)
	S3() MskClusterLoggingInfoBrokerLogsS3OutputReference
	S3Input() *MskClusterLoggingInfoBrokerLogsS3
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
	PutCloudwatchLogs(value *MskClusterLoggingInfoBrokerLogsCloudwatchLogs)
	PutFirehose(value *MskClusterLoggingInfoBrokerLogsFirehose)
	PutS3(value *MskClusterLoggingInfoBrokerLogsS3)
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

// The jsii proxy struct for MskClusterLoggingInfoBrokerLogsOutputReference
type jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) CloudwatchLogs() MskClusterLoggingInfoBrokerLogsCloudwatchLogsOutputReference {
	var returns MskClusterLoggingInfoBrokerLogsCloudwatchLogsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) CloudwatchLogsInput() *MskClusterLoggingInfoBrokerLogsCloudwatchLogs {
	var returns *MskClusterLoggingInfoBrokerLogsCloudwatchLogs
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) Firehose() MskClusterLoggingInfoBrokerLogsFirehoseOutputReference {
	var returns MskClusterLoggingInfoBrokerLogsFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) FirehoseInput() *MskClusterLoggingInfoBrokerLogsFirehose {
	var returns *MskClusterLoggingInfoBrokerLogsFirehose
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) InternalValue() *MskClusterLoggingInfoBrokerLogs {
	var returns *MskClusterLoggingInfoBrokerLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) S3() MskClusterLoggingInfoBrokerLogsS3OutputReference {
	var returns MskClusterLoggingInfoBrokerLogsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) S3Input() *MskClusterLoggingInfoBrokerLogsS3 {
	var returns *MskClusterLoggingInfoBrokerLogsS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMskClusterLoggingInfoBrokerLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MskClusterLoggingInfoBrokerLogsOutputReference {
	_init_.Initialize()

	if err := validateNewMskClusterLoggingInfoBrokerLogsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.mskCluster.MskClusterLoggingInfoBrokerLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMskClusterLoggingInfoBrokerLogsOutputReference_Override(m MskClusterLoggingInfoBrokerLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.mskCluster.MskClusterLoggingInfoBrokerLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference)SetInternalValue(val *MskClusterLoggingInfoBrokerLogs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) PutCloudwatchLogs(value *MskClusterLoggingInfoBrokerLogsCloudwatchLogs) {
	if err := m.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) PutFirehose(value *MskClusterLoggingInfoBrokerLogsFirehose) {
	if err := m.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFirehose",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) PutS3(value *MskClusterLoggingInfoBrokerLogsS3) {
	if err := m.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putS3",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		m,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		m,
		"resetFirehose",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		m,
		"resetS3",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MskClusterLoggingInfoBrokerLogsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

