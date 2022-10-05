package mwaaenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/mwaaenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwaaEnvironmentLoggingConfigurationOutputReference interface {
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
	DagProcessingLogs() MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference
	DagProcessingLogsInput() *MwaaEnvironmentLoggingConfigurationDagProcessingLogs
	// Experimental.
	Fqn() *string
	InternalValue() *MwaaEnvironmentLoggingConfiguration
	SetInternalValue(val *MwaaEnvironmentLoggingConfiguration)
	SchedulerLogs() MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference
	SchedulerLogsInput() *MwaaEnvironmentLoggingConfigurationSchedulerLogs
	TaskLogs() MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference
	TaskLogsInput() *MwaaEnvironmentLoggingConfigurationTaskLogs
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebserverLogs() MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference
	WebserverLogsInput() *MwaaEnvironmentLoggingConfigurationWebserverLogs
	WorkerLogs() MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference
	WorkerLogsInput() *MwaaEnvironmentLoggingConfigurationWorkerLogs
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
	PutDagProcessingLogs(value *MwaaEnvironmentLoggingConfigurationDagProcessingLogs)
	PutSchedulerLogs(value *MwaaEnvironmentLoggingConfigurationSchedulerLogs)
	PutTaskLogs(value *MwaaEnvironmentLoggingConfigurationTaskLogs)
	PutWebserverLogs(value *MwaaEnvironmentLoggingConfigurationWebserverLogs)
	PutWorkerLogs(value *MwaaEnvironmentLoggingConfigurationWorkerLogs)
	ResetDagProcessingLogs()
	ResetSchedulerLogs()
	ResetTaskLogs()
	ResetWebserverLogs()
	ResetWorkerLogs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MwaaEnvironmentLoggingConfigurationOutputReference
type jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) DagProcessingLogs() MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference {
	var returns MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference
	_jsii_.Get(
		j,
		"dagProcessingLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) DagProcessingLogsInput() *MwaaEnvironmentLoggingConfigurationDagProcessingLogs {
	var returns *MwaaEnvironmentLoggingConfigurationDagProcessingLogs
	_jsii_.Get(
		j,
		"dagProcessingLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) InternalValue() *MwaaEnvironmentLoggingConfiguration {
	var returns *MwaaEnvironmentLoggingConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) SchedulerLogs() MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference {
	var returns MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference
	_jsii_.Get(
		j,
		"schedulerLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) SchedulerLogsInput() *MwaaEnvironmentLoggingConfigurationSchedulerLogs {
	var returns *MwaaEnvironmentLoggingConfigurationSchedulerLogs
	_jsii_.Get(
		j,
		"schedulerLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) TaskLogs() MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference {
	var returns MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference
	_jsii_.Get(
		j,
		"taskLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) TaskLogsInput() *MwaaEnvironmentLoggingConfigurationTaskLogs {
	var returns *MwaaEnvironmentLoggingConfigurationTaskLogs
	_jsii_.Get(
		j,
		"taskLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) WebserverLogs() MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference {
	var returns MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference
	_jsii_.Get(
		j,
		"webserverLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) WebserverLogsInput() *MwaaEnvironmentLoggingConfigurationWebserverLogs {
	var returns *MwaaEnvironmentLoggingConfigurationWebserverLogs
	_jsii_.Get(
		j,
		"webserverLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) WorkerLogs() MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference {
	var returns MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference
	_jsii_.Get(
		j,
		"workerLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) WorkerLogsInput() *MwaaEnvironmentLoggingConfigurationWorkerLogs {
	var returns *MwaaEnvironmentLoggingConfigurationWorkerLogs
	_jsii_.Get(
		j,
		"workerLogsInput",
		&returns,
	)
	return returns
}


func NewMwaaEnvironmentLoggingConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MwaaEnvironmentLoggingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMwaaEnvironmentLoggingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.mwaaEnvironment.MwaaEnvironmentLoggingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMwaaEnvironmentLoggingConfigurationOutputReference_Override(m MwaaEnvironmentLoggingConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.mwaaEnvironment.MwaaEnvironmentLoggingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference)SetInternalValue(val *MwaaEnvironmentLoggingConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) PutDagProcessingLogs(value *MwaaEnvironmentLoggingConfigurationDagProcessingLogs) {
	if err := m.validatePutDagProcessingLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDagProcessingLogs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) PutSchedulerLogs(value *MwaaEnvironmentLoggingConfigurationSchedulerLogs) {
	if err := m.validatePutSchedulerLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSchedulerLogs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) PutTaskLogs(value *MwaaEnvironmentLoggingConfigurationTaskLogs) {
	if err := m.validatePutTaskLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTaskLogs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) PutWebserverLogs(value *MwaaEnvironmentLoggingConfigurationWebserverLogs) {
	if err := m.validatePutWebserverLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWebserverLogs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) PutWorkerLogs(value *MwaaEnvironmentLoggingConfigurationWorkerLogs) {
	if err := m.validatePutWorkerLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWorkerLogs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ResetDagProcessingLogs() {
	_jsii_.InvokeVoid(
		m,
		"resetDagProcessingLogs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ResetSchedulerLogs() {
	_jsii_.InvokeVoid(
		m,
		"resetSchedulerLogs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ResetTaskLogs() {
	_jsii_.InvokeVoid(
		m,
		"resetTaskLogs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ResetWebserverLogs() {
	_jsii_.InvokeVoid(
		m,
		"resetWebserverLogs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ResetWorkerLogs() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkerLogs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

