// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dmsendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsEndpointKafkaSettingsOutputReference interface {
	cdktf.ComplexObject
	Broker() *string
	SetBroker(val *string)
	BrokerInput() *string
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
	IncludeControlDetails() interface{}
	SetIncludeControlDetails(val interface{})
	IncludeControlDetailsInput() interface{}
	IncludeNullAndEmpty() interface{}
	SetIncludeNullAndEmpty(val interface{})
	IncludeNullAndEmptyInput() interface{}
	IncludePartitionValue() interface{}
	SetIncludePartitionValue(val interface{})
	IncludePartitionValueInput() interface{}
	IncludeTableAlterOperations() interface{}
	SetIncludeTableAlterOperations(val interface{})
	IncludeTableAlterOperationsInput() interface{}
	IncludeTransactionDetails() interface{}
	SetIncludeTransactionDetails(val interface{})
	IncludeTransactionDetailsInput() interface{}
	InternalValue() *DmsEndpointKafkaSettings
	SetInternalValue(val *DmsEndpointKafkaSettings)
	MessageFormat() *string
	SetMessageFormat(val *string)
	MessageFormatInput() *string
	MessageMaxBytes() *float64
	SetMessageMaxBytes(val *float64)
	MessageMaxBytesInput() *float64
	NoHexPrefix() interface{}
	SetNoHexPrefix(val interface{})
	NoHexPrefixInput() interface{}
	PartitionIncludeSchemaTable() interface{}
	SetPartitionIncludeSchemaTable(val interface{})
	PartitionIncludeSchemaTableInput() interface{}
	SaslMechanism() *string
	SetSaslMechanism(val *string)
	SaslMechanismInput() *string
	SaslPassword() *string
	SetSaslPassword(val *string)
	SaslPasswordInput() *string
	SaslUsername() *string
	SetSaslUsername(val *string)
	SaslUsernameInput() *string
	SecurityProtocol() *string
	SetSecurityProtocol(val *string)
	SecurityProtocolInput() *string
	SslCaCertificateArn() *string
	SetSslCaCertificateArn(val *string)
	SslCaCertificateArnInput() *string
	SslClientCertificateArn() *string
	SetSslClientCertificateArn(val *string)
	SslClientCertificateArnInput() *string
	SslClientKeyArn() *string
	SetSslClientKeyArn(val *string)
	SslClientKeyArnInput() *string
	SslClientKeyPassword() *string
	SetSslClientKeyPassword(val *string)
	SslClientKeyPasswordInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Topic() *string
	SetTopic(val *string)
	TopicInput() *string
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
	ResetIncludeControlDetails()
	ResetIncludeNullAndEmpty()
	ResetIncludePartitionValue()
	ResetIncludeTableAlterOperations()
	ResetIncludeTransactionDetails()
	ResetMessageFormat()
	ResetMessageMaxBytes()
	ResetNoHexPrefix()
	ResetPartitionIncludeSchemaTable()
	ResetSaslMechanism()
	ResetSaslPassword()
	ResetSaslUsername()
	ResetSecurityProtocol()
	ResetSslCaCertificateArn()
	ResetSslClientCertificateArn()
	ResetSslClientKeyArn()
	ResetSslClientKeyPassword()
	ResetTopic()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsEndpointKafkaSettingsOutputReference
type jsiiProxy_DmsEndpointKafkaSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) Broker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"broker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) BrokerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeControlDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeControlDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeNullAndEmpty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeNullAndEmptyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmptyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludePartitionValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludePartitionValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeTableAlterOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeTableAlterOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeTransactionDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeTransactionDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) InternalValue() *DmsEndpointKafkaSettings {
	var returns *DmsEndpointKafkaSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) MessageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) MessageFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) MessageMaxBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageMaxBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) MessageMaxBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageMaxBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) NoHexPrefix() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHexPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) NoHexPrefixInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHexPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) PartitionIncludeSchemaTable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) PartitionIncludeSchemaTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SaslMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SaslMechanismInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SaslPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SaslPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SaslUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SaslUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslCaCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslCaCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslClientCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslClientCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslClientKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslClientKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslClientKeyPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslClientKeyPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) TopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}


func NewDmsEndpointKafkaSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DmsEndpointKafkaSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsEndpointKafkaSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpointKafkaSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpointKafkaSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsEndpointKafkaSettingsOutputReference_Override(d DmsEndpointKafkaSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpointKafkaSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetBroker(val *string) {
	if err := j.validateSetBrokerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"broker",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetIncludeControlDetails(val interface{}) {
	if err := j.validateSetIncludeControlDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeControlDetails",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetIncludeNullAndEmpty(val interface{}) {
	if err := j.validateSetIncludeNullAndEmptyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeNullAndEmpty",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetIncludePartitionValue(val interface{}) {
	if err := j.validateSetIncludePartitionValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePartitionValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetIncludeTableAlterOperations(val interface{}) {
	if err := j.validateSetIncludeTableAlterOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTableAlterOperations",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetIncludeTransactionDetails(val interface{}) {
	if err := j.validateSetIncludeTransactionDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTransactionDetails",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetInternalValue(val *DmsEndpointKafkaSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetMessageFormat(val *string) {
	if err := j.validateSetMessageFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageFormat",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetMessageMaxBytes(val *float64) {
	if err := j.validateSetMessageMaxBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageMaxBytes",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetNoHexPrefix(val interface{}) {
	if err := j.validateSetNoHexPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noHexPrefix",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetPartitionIncludeSchemaTable(val interface{}) {
	if err := j.validateSetPartitionIncludeSchemaTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionIncludeSchemaTable",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetSaslMechanism(val *string) {
	if err := j.validateSetSaslMechanismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslMechanism",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetSaslPassword(val *string) {
	if err := j.validateSetSaslPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslPassword",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetSaslUsername(val *string) {
	if err := j.validateSetSaslUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslUsername",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetSecurityProtocol(val *string) {
	if err := j.validateSetSecurityProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetSslCaCertificateArn(val *string) {
	if err := j.validateSetSslCaCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaCertificateArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetSslClientCertificateArn(val *string) {
	if err := j.validateSetSslClientCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslClientCertificateArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetSslClientKeyArn(val *string) {
	if err := j.validateSetSslClientKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslClientKeyArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetSslClientKeyPassword(val *string) {
	if err := j.validateSetSslClientKeyPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslClientKeyPassword",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference)SetTopic(val *string) {
	if err := j.validateSetTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topic",
		val,
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetIncludeControlDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeControlDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetIncludeNullAndEmpty() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeNullAndEmpty",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetIncludePartitionValue() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludePartitionValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetIncludeTableAlterOperations() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeTableAlterOperations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetIncludeTransactionDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeTransactionDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetMessageFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetMessageFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetMessageMaxBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetMessageMaxBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetNoHexPrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetNoHexPrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetPartitionIncludeSchemaTable() {
	_jsii_.InvokeVoid(
		d,
		"resetPartitionIncludeSchemaTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSaslMechanism() {
	_jsii_.InvokeVoid(
		d,
		"resetSaslMechanism",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSaslPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetSaslPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSaslUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetSaslUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSslCaCertificateArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSslCaCertificateArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSslClientCertificateArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSslClientCertificateArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSslClientKeyArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSslClientKeyArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSslClientKeyPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetSslClientKeyPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetTopic() {
	_jsii_.InvokeVoid(
		d,
		"resetTopic",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

