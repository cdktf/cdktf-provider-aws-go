// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/quicksightdatasource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightDataSourceParametersOutputReference interface {
	cdktf.ComplexObject
	AmazonElasticsearch() QuicksightDataSourceParametersAmazonElasticsearchOutputReference
	AmazonElasticsearchInput() *QuicksightDataSourceParametersAmazonElasticsearch
	Athena() QuicksightDataSourceParametersAthenaOutputReference
	AthenaInput() *QuicksightDataSourceParametersAthena
	Aurora() QuicksightDataSourceParametersAuroraOutputReference
	AuroraInput() *QuicksightDataSourceParametersAurora
	AuroraPostgresql() QuicksightDataSourceParametersAuroraPostgresqlOutputReference
	AuroraPostgresqlInput() *QuicksightDataSourceParametersAuroraPostgresql
	AwsIotAnalytics() QuicksightDataSourceParametersAwsIotAnalyticsOutputReference
	AwsIotAnalyticsInput() *QuicksightDataSourceParametersAwsIotAnalytics
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
	InternalValue() *QuicksightDataSourceParameters
	SetInternalValue(val *QuicksightDataSourceParameters)
	Jira() QuicksightDataSourceParametersJiraOutputReference
	JiraInput() *QuicksightDataSourceParametersJira
	MariaDb() QuicksightDataSourceParametersMariaDbOutputReference
	MariaDbInput() *QuicksightDataSourceParametersMariaDb
	Mysql() QuicksightDataSourceParametersMysqlOutputReference
	MysqlInput() *QuicksightDataSourceParametersMysql
	Oracle() QuicksightDataSourceParametersOracleOutputReference
	OracleInput() *QuicksightDataSourceParametersOracle
	Postgresql() QuicksightDataSourceParametersPostgresqlOutputReference
	PostgresqlInput() *QuicksightDataSourceParametersPostgresql
	Presto() QuicksightDataSourceParametersPrestoOutputReference
	PrestoInput() *QuicksightDataSourceParametersPresto
	Rds() QuicksightDataSourceParametersRdsOutputReference
	RdsInput() *QuicksightDataSourceParametersRds
	Redshift() QuicksightDataSourceParametersRedshiftOutputReference
	RedshiftInput() *QuicksightDataSourceParametersRedshift
	S3() QuicksightDataSourceParametersS3OutputReference
	S3Input() *QuicksightDataSourceParametersS3
	ServiceNow() QuicksightDataSourceParametersServiceNowOutputReference
	ServiceNowInput() *QuicksightDataSourceParametersServiceNow
	Snowflake() QuicksightDataSourceParametersSnowflakeOutputReference
	SnowflakeInput() *QuicksightDataSourceParametersSnowflake
	Spark() QuicksightDataSourceParametersSparkOutputReference
	SparkInput() *QuicksightDataSourceParametersSpark
	SqlServer() QuicksightDataSourceParametersSqlServerOutputReference
	SqlServerInput() *QuicksightDataSourceParametersSqlServer
	Teradata() QuicksightDataSourceParametersTeradataOutputReference
	TeradataInput() *QuicksightDataSourceParametersTeradata
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Twitter() QuicksightDataSourceParametersTwitterOutputReference
	TwitterInput() *QuicksightDataSourceParametersTwitter
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
	PutAmazonElasticsearch(value *QuicksightDataSourceParametersAmazonElasticsearch)
	PutAthena(value *QuicksightDataSourceParametersAthena)
	PutAurora(value *QuicksightDataSourceParametersAurora)
	PutAuroraPostgresql(value *QuicksightDataSourceParametersAuroraPostgresql)
	PutAwsIotAnalytics(value *QuicksightDataSourceParametersAwsIotAnalytics)
	PutJira(value *QuicksightDataSourceParametersJira)
	PutMariaDb(value *QuicksightDataSourceParametersMariaDb)
	PutMysql(value *QuicksightDataSourceParametersMysql)
	PutOracle(value *QuicksightDataSourceParametersOracle)
	PutPostgresql(value *QuicksightDataSourceParametersPostgresql)
	PutPresto(value *QuicksightDataSourceParametersPresto)
	PutRds(value *QuicksightDataSourceParametersRds)
	PutRedshift(value *QuicksightDataSourceParametersRedshift)
	PutS3(value *QuicksightDataSourceParametersS3)
	PutServiceNow(value *QuicksightDataSourceParametersServiceNow)
	PutSnowflake(value *QuicksightDataSourceParametersSnowflake)
	PutSpark(value *QuicksightDataSourceParametersSpark)
	PutSqlServer(value *QuicksightDataSourceParametersSqlServer)
	PutTeradata(value *QuicksightDataSourceParametersTeradata)
	PutTwitter(value *QuicksightDataSourceParametersTwitter)
	ResetAmazonElasticsearch()
	ResetAthena()
	ResetAurora()
	ResetAuroraPostgresql()
	ResetAwsIotAnalytics()
	ResetJira()
	ResetMariaDb()
	ResetMysql()
	ResetOracle()
	ResetPostgresql()
	ResetPresto()
	ResetRds()
	ResetRedshift()
	ResetS3()
	ResetServiceNow()
	ResetSnowflake()
	ResetSpark()
	ResetSqlServer()
	ResetTeradata()
	ResetTwitter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDataSourceParametersOutputReference
type jsiiProxy_QuicksightDataSourceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AmazonElasticsearch() QuicksightDataSourceParametersAmazonElasticsearchOutputReference {
	var returns QuicksightDataSourceParametersAmazonElasticsearchOutputReference
	_jsii_.Get(
		j,
		"amazonElasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AmazonElasticsearchInput() *QuicksightDataSourceParametersAmazonElasticsearch {
	var returns *QuicksightDataSourceParametersAmazonElasticsearch
	_jsii_.Get(
		j,
		"amazonElasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Athena() QuicksightDataSourceParametersAthenaOutputReference {
	var returns QuicksightDataSourceParametersAthenaOutputReference
	_jsii_.Get(
		j,
		"athena",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AthenaInput() *QuicksightDataSourceParametersAthena {
	var returns *QuicksightDataSourceParametersAthena
	_jsii_.Get(
		j,
		"athenaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Aurora() QuicksightDataSourceParametersAuroraOutputReference {
	var returns QuicksightDataSourceParametersAuroraOutputReference
	_jsii_.Get(
		j,
		"aurora",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AuroraInput() *QuicksightDataSourceParametersAurora {
	var returns *QuicksightDataSourceParametersAurora
	_jsii_.Get(
		j,
		"auroraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AuroraPostgresql() QuicksightDataSourceParametersAuroraPostgresqlOutputReference {
	var returns QuicksightDataSourceParametersAuroraPostgresqlOutputReference
	_jsii_.Get(
		j,
		"auroraPostgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AuroraPostgresqlInput() *QuicksightDataSourceParametersAuroraPostgresql {
	var returns *QuicksightDataSourceParametersAuroraPostgresql
	_jsii_.Get(
		j,
		"auroraPostgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AwsIotAnalytics() QuicksightDataSourceParametersAwsIotAnalyticsOutputReference {
	var returns QuicksightDataSourceParametersAwsIotAnalyticsOutputReference
	_jsii_.Get(
		j,
		"awsIotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AwsIotAnalyticsInput() *QuicksightDataSourceParametersAwsIotAnalytics {
	var returns *QuicksightDataSourceParametersAwsIotAnalytics
	_jsii_.Get(
		j,
		"awsIotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) InternalValue() *QuicksightDataSourceParameters {
	var returns *QuicksightDataSourceParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Jira() QuicksightDataSourceParametersJiraOutputReference {
	var returns QuicksightDataSourceParametersJiraOutputReference
	_jsii_.Get(
		j,
		"jira",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) JiraInput() *QuicksightDataSourceParametersJira {
	var returns *QuicksightDataSourceParametersJira
	_jsii_.Get(
		j,
		"jiraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) MariaDb() QuicksightDataSourceParametersMariaDbOutputReference {
	var returns QuicksightDataSourceParametersMariaDbOutputReference
	_jsii_.Get(
		j,
		"mariaDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) MariaDbInput() *QuicksightDataSourceParametersMariaDb {
	var returns *QuicksightDataSourceParametersMariaDb
	_jsii_.Get(
		j,
		"mariaDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Mysql() QuicksightDataSourceParametersMysqlOutputReference {
	var returns QuicksightDataSourceParametersMysqlOutputReference
	_jsii_.Get(
		j,
		"mysql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) MysqlInput() *QuicksightDataSourceParametersMysql {
	var returns *QuicksightDataSourceParametersMysql
	_jsii_.Get(
		j,
		"mysqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Oracle() QuicksightDataSourceParametersOracleOutputReference {
	var returns QuicksightDataSourceParametersOracleOutputReference
	_jsii_.Get(
		j,
		"oracle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) OracleInput() *QuicksightDataSourceParametersOracle {
	var returns *QuicksightDataSourceParametersOracle
	_jsii_.Get(
		j,
		"oracleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Postgresql() QuicksightDataSourceParametersPostgresqlOutputReference {
	var returns QuicksightDataSourceParametersPostgresqlOutputReference
	_jsii_.Get(
		j,
		"postgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) PostgresqlInput() *QuicksightDataSourceParametersPostgresql {
	var returns *QuicksightDataSourceParametersPostgresql
	_jsii_.Get(
		j,
		"postgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Presto() QuicksightDataSourceParametersPrestoOutputReference {
	var returns QuicksightDataSourceParametersPrestoOutputReference
	_jsii_.Get(
		j,
		"presto",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) PrestoInput() *QuicksightDataSourceParametersPresto {
	var returns *QuicksightDataSourceParametersPresto
	_jsii_.Get(
		j,
		"prestoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Rds() QuicksightDataSourceParametersRdsOutputReference {
	var returns QuicksightDataSourceParametersRdsOutputReference
	_jsii_.Get(
		j,
		"rds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) RdsInput() *QuicksightDataSourceParametersRds {
	var returns *QuicksightDataSourceParametersRds
	_jsii_.Get(
		j,
		"rdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Redshift() QuicksightDataSourceParametersRedshiftOutputReference {
	var returns QuicksightDataSourceParametersRedshiftOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) RedshiftInput() *QuicksightDataSourceParametersRedshift {
	var returns *QuicksightDataSourceParametersRedshift
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) S3() QuicksightDataSourceParametersS3OutputReference {
	var returns QuicksightDataSourceParametersS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) S3Input() *QuicksightDataSourceParametersS3 {
	var returns *QuicksightDataSourceParametersS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) ServiceNow() QuicksightDataSourceParametersServiceNowOutputReference {
	var returns QuicksightDataSourceParametersServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) ServiceNowInput() *QuicksightDataSourceParametersServiceNow {
	var returns *QuicksightDataSourceParametersServiceNow
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Snowflake() QuicksightDataSourceParametersSnowflakeOutputReference {
	var returns QuicksightDataSourceParametersSnowflakeOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) SnowflakeInput() *QuicksightDataSourceParametersSnowflake {
	var returns *QuicksightDataSourceParametersSnowflake
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Spark() QuicksightDataSourceParametersSparkOutputReference {
	var returns QuicksightDataSourceParametersSparkOutputReference
	_jsii_.Get(
		j,
		"spark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) SparkInput() *QuicksightDataSourceParametersSpark {
	var returns *QuicksightDataSourceParametersSpark
	_jsii_.Get(
		j,
		"sparkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) SqlServer() QuicksightDataSourceParametersSqlServerOutputReference {
	var returns QuicksightDataSourceParametersSqlServerOutputReference
	_jsii_.Get(
		j,
		"sqlServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) SqlServerInput() *QuicksightDataSourceParametersSqlServer {
	var returns *QuicksightDataSourceParametersSqlServer
	_jsii_.Get(
		j,
		"sqlServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Teradata() QuicksightDataSourceParametersTeradataOutputReference {
	var returns QuicksightDataSourceParametersTeradataOutputReference
	_jsii_.Get(
		j,
		"teradata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) TeradataInput() *QuicksightDataSourceParametersTeradata {
	var returns *QuicksightDataSourceParametersTeradata
	_jsii_.Get(
		j,
		"teradataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Twitter() QuicksightDataSourceParametersTwitterOutputReference {
	var returns QuicksightDataSourceParametersTwitterOutputReference
	_jsii_.Get(
		j,
		"twitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) TwitterInput() *QuicksightDataSourceParametersTwitter {
	var returns *QuicksightDataSourceParametersTwitter
	_jsii_.Get(
		j,
		"twitterInput",
		&returns,
	)
	return returns
}


func NewQuicksightDataSourceParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QuicksightDataSourceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSourceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSourceParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDataSource.QuicksightDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersOutputReference_Override(q QuicksightDataSourceParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDataSource.QuicksightDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference)SetInternalValue(val *QuicksightDataSourceParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutAmazonElasticsearch(value *QuicksightDataSourceParametersAmazonElasticsearch) {
	if err := q.validatePutAmazonElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAmazonElasticsearch",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutAthena(value *QuicksightDataSourceParametersAthena) {
	if err := q.validatePutAthenaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAthena",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutAurora(value *QuicksightDataSourceParametersAurora) {
	if err := q.validatePutAuroraParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAurora",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutAuroraPostgresql(value *QuicksightDataSourceParametersAuroraPostgresql) {
	if err := q.validatePutAuroraPostgresqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAuroraPostgresql",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutAwsIotAnalytics(value *QuicksightDataSourceParametersAwsIotAnalytics) {
	if err := q.validatePutAwsIotAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAwsIotAnalytics",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutJira(value *QuicksightDataSourceParametersJira) {
	if err := q.validatePutJiraParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putJira",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutMariaDb(value *QuicksightDataSourceParametersMariaDb) {
	if err := q.validatePutMariaDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putMariaDb",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutMysql(value *QuicksightDataSourceParametersMysql) {
	if err := q.validatePutMysqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putMysql",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutOracle(value *QuicksightDataSourceParametersOracle) {
	if err := q.validatePutOracleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putOracle",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutPostgresql(value *QuicksightDataSourceParametersPostgresql) {
	if err := q.validatePutPostgresqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPostgresql",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutPresto(value *QuicksightDataSourceParametersPresto) {
	if err := q.validatePutPrestoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPresto",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutRds(value *QuicksightDataSourceParametersRds) {
	if err := q.validatePutRdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRds",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutRedshift(value *QuicksightDataSourceParametersRedshift) {
	if err := q.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRedshift",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutS3(value *QuicksightDataSourceParametersS3) {
	if err := q.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putS3",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutServiceNow(value *QuicksightDataSourceParametersServiceNow) {
	if err := q.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutSnowflake(value *QuicksightDataSourceParametersSnowflake) {
	if err := q.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutSpark(value *QuicksightDataSourceParametersSpark) {
	if err := q.validatePutSparkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSpark",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutSqlServer(value *QuicksightDataSourceParametersSqlServer) {
	if err := q.validatePutSqlServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSqlServer",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutTeradata(value *QuicksightDataSourceParametersTeradata) {
	if err := q.validatePutTeradataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTeradata",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutTwitter(value *QuicksightDataSourceParametersTwitter) {
	if err := q.validatePutTwitterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTwitter",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetAmazonElasticsearch() {
	_jsii_.InvokeVoid(
		q,
		"resetAmazonElasticsearch",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetAthena() {
	_jsii_.InvokeVoid(
		q,
		"resetAthena",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetAurora() {
	_jsii_.InvokeVoid(
		q,
		"resetAurora",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetAuroraPostgresql() {
	_jsii_.InvokeVoid(
		q,
		"resetAuroraPostgresql",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetAwsIotAnalytics() {
	_jsii_.InvokeVoid(
		q,
		"resetAwsIotAnalytics",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetJira() {
	_jsii_.InvokeVoid(
		q,
		"resetJira",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetMariaDb() {
	_jsii_.InvokeVoid(
		q,
		"resetMariaDb",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetMysql() {
	_jsii_.InvokeVoid(
		q,
		"resetMysql",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetOracle() {
	_jsii_.InvokeVoid(
		q,
		"resetOracle",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetPostgresql() {
	_jsii_.InvokeVoid(
		q,
		"resetPostgresql",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetPresto() {
	_jsii_.InvokeVoid(
		q,
		"resetPresto",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetRds() {
	_jsii_.InvokeVoid(
		q,
		"resetRds",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		q,
		"resetRedshift",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		q,
		"resetS3",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		q,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		q,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetSpark() {
	_jsii_.InvokeVoid(
		q,
		"resetSpark",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetSqlServer() {
	_jsii_.InvokeVoid(
		q,
		"resetSqlServer",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetTeradata() {
	_jsii_.InvokeVoid(
		q,
		"resetTeradata",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetTwitter() {
	_jsii_.InvokeVoid(
		q,
		"resetTwitter",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

