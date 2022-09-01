// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AwsProviderConfig struct {
	// The access key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#access_key AwsProvider#access_key}
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#alias AwsProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#allowed_account_ids AwsProvider#allowed_account_ids}.
	AllowedAccountIds *[]*string `field:"optional" json:"allowedAccountIds" yaml:"allowedAccountIds"`
	// assume_role block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#assume_role AwsProvider#assume_role}
	AssumeRole *AwsProviderAssumeRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// assume_role_with_web_identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#assume_role_with_web_identity AwsProvider#assume_role_with_web_identity}
	AssumeRoleWithWebIdentity *AwsProviderAssumeRoleWithWebIdentity `field:"optional" json:"assumeRoleWithWebIdentity" yaml:"assumeRoleWithWebIdentity"`
	// File containing custom root and intermediate certificates.
	//
	// Can also be configured using the `AWS_CA_BUNDLE` environment variable. (Setting `ca_bundle` in the shared config file is not supported.)
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#custom_ca_bundle AwsProvider#custom_ca_bundle}
	CustomCaBundle *string `field:"optional" json:"customCaBundle" yaml:"customCaBundle"`
	// default_tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#default_tags AwsProvider#default_tags}
	DefaultTags *AwsProviderDefaultTags `field:"optional" json:"defaultTags" yaml:"defaultTags"`
	// Address of the EC2 metadata service endpoint to use. Can also be configured using the `AWS_EC2_METADATA_SERVICE_ENDPOINT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ec2_metadata_service_endpoint AwsProvider#ec2_metadata_service_endpoint}
	Ec2MetadataServiceEndpoint *string `field:"optional" json:"ec2MetadataServiceEndpoint" yaml:"ec2MetadataServiceEndpoint"`
	// Protocol to use with EC2 metadata service endpoint.Valid values are `IPv4` and `IPv6`. Can also be configured using the `AWS_EC2_METADATA_SERVICE_ENDPOINT_MODE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ec2_metadata_service_endpoint_mode AwsProvider#ec2_metadata_service_endpoint_mode}
	Ec2MetadataServiceEndpointMode *string `field:"optional" json:"ec2MetadataServiceEndpointMode" yaml:"ec2MetadataServiceEndpointMode"`
	// endpoints block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#endpoints AwsProvider#endpoints}
	Endpoints interface{} `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#forbidden_account_ids AwsProvider#forbidden_account_ids}.
	ForbiddenAccountIds *[]*string `field:"optional" json:"forbiddenAccountIds" yaml:"forbiddenAccountIds"`
	// The address of an HTTP proxy to use when accessing the AWS API.
	//
	// Can also be configured using the `HTTP_PROXY` or `HTTPS_PROXY` environment variables.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#http_proxy AwsProvider#http_proxy}
	HttpProxy *string `field:"optional" json:"httpProxy" yaml:"httpProxy"`
	// ignore_tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ignore_tags AwsProvider#ignore_tags}
	IgnoreTags *AwsProviderIgnoreTags `field:"optional" json:"ignoreTags" yaml:"ignoreTags"`
	// Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#insecure AwsProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// The maximum number of times an AWS API request is being executed.
	//
	// If the API request still fails, an error is
	// thrown.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#max_retries AwsProvider#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// The profile for API operations. If not set, the default profile created with `aws configure` will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#profile AwsProvider#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#region AwsProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Set this to true to enable the request to use path-style addressing, i.e., https://s3.amazonaws.com/BUCKET/KEY. By default, the S3 client will use virtual hosted bucket addressing when possible (https://BUCKET.s3.amazonaws.com/KEY). Specific to the Amazon S3 service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#s3_force_path_style AwsProvider#s3_force_path_style}
	S3ForcePathStyle interface{} `field:"optional" json:"s3ForcePathStyle" yaml:"s3ForcePathStyle"`
	// Set this to true to enable the request to use path-style addressing, i.e., https://s3.amazonaws.com/BUCKET/KEY. By default, the S3 client will use virtual hosted bucket addressing when possible (https://BUCKET.s3.amazonaws.com/KEY). Specific to the Amazon S3 service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#s3_use_path_style AwsProvider#s3_use_path_style}
	S3UsePathStyle interface{} `field:"optional" json:"s3UsePathStyle" yaml:"s3UsePathStyle"`
	// The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#secret_key AwsProvider#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// List of paths to shared config files. If not set, defaults to [~/.aws/config].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#shared_config_files AwsProvider#shared_config_files}
	SharedConfigFiles *[]*string `field:"optional" json:"sharedConfigFiles" yaml:"sharedConfigFiles"`
	// The path to the shared credentials file. If not set, defaults to ~/.aws/credentials.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#shared_credentials_file AwsProvider#shared_credentials_file}
	SharedCredentialsFile *string `field:"optional" json:"sharedCredentialsFile" yaml:"sharedCredentialsFile"`
	// List of paths to shared credentials files. If not set, defaults to [~/.aws/credentials].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#shared_credentials_files AwsProvider#shared_credentials_files}
	SharedCredentialsFiles *[]*string `field:"optional" json:"sharedCredentialsFiles" yaml:"sharedCredentialsFiles"`
	// Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#skip_credentials_validation AwsProvider#skip_credentials_validation}
	SkipCredentialsValidation interface{} `field:"optional" json:"skipCredentialsValidation" yaml:"skipCredentialsValidation"`
	// Skip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes permissions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#skip_get_ec2_platforms AwsProvider#skip_get_ec2_platforms}
	SkipGetEc2Platforms interface{} `field:"optional" json:"skipGetEc2Platforms" yaml:"skipGetEc2Platforms"`
	// Skip the AWS Metadata API check. Used for AWS API implementations that do not have a metadata api endpoint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#skip_metadata_api_check AwsProvider#skip_metadata_api_check}
	SkipMetadataApiCheck *string `field:"optional" json:"skipMetadataApiCheck" yaml:"skipMetadataApiCheck"`
	// Skip static validation of region name.
	//
	// Used by users of alternative AWS-like APIs or users w/ access to regions that are not public (yet).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#skip_region_validation AwsProvider#skip_region_validation}
	SkipRegionValidation interface{} `field:"optional" json:"skipRegionValidation" yaml:"skipRegionValidation"`
	// Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#skip_requesting_account_id AwsProvider#skip_requesting_account_id}
	SkipRequestingAccountId interface{} `field:"optional" json:"skipRequestingAccountId" yaml:"skipRequestingAccountId"`
	// The region where AWS STS operations will take place. Examples are us-east-1 and us-west-2.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#sts_region AwsProvider#sts_region}
	StsRegion *string `field:"optional" json:"stsRegion" yaml:"stsRegion"`
	// session token. A session token is only required if you are using temporary security credentials.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#token AwsProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Resolve an endpoint with DualStack capability.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#use_dualstack_endpoint AwsProvider#use_dualstack_endpoint}
	UseDualstackEndpoint interface{} `field:"optional" json:"useDualstackEndpoint" yaml:"useDualstackEndpoint"`
	// Resolve an endpoint with FIPS capability.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#use_fips_endpoint AwsProvider#use_fips_endpoint}
	UseFipsEndpoint interface{} `field:"optional" json:"useFipsEndpoint" yaml:"useFipsEndpoint"`
}

