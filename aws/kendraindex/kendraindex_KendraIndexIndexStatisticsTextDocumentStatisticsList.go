package kendraindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/kendraindex/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraIndexIndexStatisticsTextDocumentStatisticsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) KendraIndexIndexStatisticsTextDocumentStatisticsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraIndexIndexStatisticsTextDocumentStatisticsList
type jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewKendraIndexIndexStatisticsTextDocumentStatisticsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) KendraIndexIndexStatisticsTextDocumentStatisticsList {
	_init_.Initialize()

	if err := validateNewKendraIndexIndexStatisticsTextDocumentStatisticsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList{}

	_jsii_.Create(
		"@cdktf/provider-aws.kendraIndex.KendraIndexIndexStatisticsTextDocumentStatisticsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewKendraIndexIndexStatisticsTextDocumentStatisticsList_Override(k KendraIndexIndexStatisticsTextDocumentStatisticsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kendraIndex.KendraIndexIndexStatisticsTextDocumentStatisticsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		k,
	)
}

func (j *jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (k *jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList) Get(index *float64) KendraIndexIndexStatisticsTextDocumentStatisticsOutputReference {
	if err := k.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns KendraIndexIndexStatisticsTextDocumentStatisticsOutputReference

	_jsii_.Invoke(
		k,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexIndexStatisticsTextDocumentStatisticsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

