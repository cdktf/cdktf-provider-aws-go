// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AwsProviderConfig struct {
	// The access key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#access_key AwsProvider#access_key}
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#alias AwsProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#allowed_account_ids AwsProvider#allowed_account_ids}.
	AllowedAccountIds *[]*string `field:"optional" json:"allowedAccountIds" yaml:"allowedAccountIds"`
	// assume_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#assume_role AwsProvider#assume_role}
	AssumeRole interface{} `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// assume_role_with_web_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#assume_role_with_web_identity AwsProvider#assume_role_with_web_identity}
	AssumeRoleWithWebIdentity interface{} `field:"optional" json:"assumeRoleWithWebIdentity" yaml:"assumeRoleWithWebIdentity"`
	// File containing custom root and intermediate certificates.
	//
	// Can also be configured using the `AWS_CA_BUNDLE` environment variable. (Setting `ca_bundle` in the shared config file is not supported.)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#custom_ca_bundle AwsProvider#custom_ca_bundle}
	CustomCaBundle *string `field:"optional" json:"customCaBundle" yaml:"customCaBundle"`
	// default_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#default_tags AwsProvider#default_tags}
	DefaultTags interface{} `field:"optional" json:"defaultTags" yaml:"defaultTags"`
	// Address of the EC2 metadata service endpoint to use. Can also be configured using the `AWS_EC2_METADATA_SERVICE_ENDPOINT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#ec2_metadata_service_endpoint AwsProvider#ec2_metadata_service_endpoint}
	Ec2MetadataServiceEndpoint *string `field:"optional" json:"ec2MetadataServiceEndpoint" yaml:"ec2MetadataServiceEndpoint"`
	// Protocol to use with EC2 metadata service endpoint.Valid values are `IPv4` and `IPv6`. Can also be configured using the `AWS_EC2_METADATA_SERVICE_ENDPOINT_MODE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#ec2_metadata_service_endpoint_mode AwsProvider#ec2_metadata_service_endpoint_mode}
	Ec2MetadataServiceEndpointMode *string `field:"optional" json:"ec2MetadataServiceEndpointMode" yaml:"ec2MetadataServiceEndpointMode"`
	// endpoints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#endpoints AwsProvider#endpoints}
	Endpoints interface{} `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#forbidden_account_ids AwsProvider#forbidden_account_ids}.
	ForbiddenAccountIds *[]*string `field:"optional" json:"forbiddenAccountIds" yaml:"forbiddenAccountIds"`
	// URL of a proxy to use for HTTP requests when accessing the AWS API.
	//
	// Can also be set using the `HTTP_PROXY` or `http_proxy` environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#http_proxy AwsProvider#http_proxy}
	HttpProxy *string `field:"optional" json:"httpProxy" yaml:"httpProxy"`
	// URL of a proxy to use for HTTPS requests when accessing the AWS API.
	//
	// Can also be set using the `HTTPS_PROXY` or `https_proxy` environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#https_proxy AwsProvider#https_proxy}
	HttpsProxy *string `field:"optional" json:"httpsProxy" yaml:"httpsProxy"`
	// ignore_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#ignore_tags AwsProvider#ignore_tags}
	IgnoreTags interface{} `field:"optional" json:"ignoreTags" yaml:"ignoreTags"`
	// Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#insecure AwsProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// The maximum number of times an AWS API request is being executed.
	//
	// If the API request still fails, an error is
	// thrown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#max_retries AwsProvider#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Comma-separated list of hosts that should not use HTTP or HTTPS proxies.
	//
	// Can also be set using the `NO_PROXY` or `no_proxy` environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#no_proxy AwsProvider#no_proxy}
	NoProxy *string `field:"optional" json:"noProxy" yaml:"noProxy"`
	// The profile for API operations. If not set, the default profile created with `aws configure` will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#profile AwsProvider#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#region AwsProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Specifies how retries are attempted.
	//
	// Valid values are `standard` and `adaptive`. Can also be configured using the `AWS_RETRY_MODE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#retry_mode AwsProvider#retry_mode}
	RetryMode *string `field:"optional" json:"retryMode" yaml:"retryMode"`
	// Specifies whether S3 API calls in the `us-east-1` region use the legacy global endpoint or a regional endpoint.
	//
	// Valid values are `legacy` or `regional`. Can also be configured using the `AWS_S3_US_EAST_1_REGIONAL_ENDPOINT` environment variable or the `s3_us_east_1_regional_endpoint` shared config file parameter
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#s3_us_east_1_regional_endpoint AwsProvider#s3_us_east_1_regional_endpoint}
	S3UsEast1RegionalEndpoint *string `field:"optional" json:"s3UsEast1RegionalEndpoint" yaml:"s3UsEast1RegionalEndpoint"`
	// Set this to true to enable the request to use path-style addressing, i.e., https://s3.amazonaws.com/BUCKET/KEY. By default, the S3 client will use virtual hosted bucket addressing when possible (https://BUCKET.s3.amazonaws.com/KEY). Specific to the Amazon S3 service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#s3_use_path_style AwsProvider#s3_use_path_style}
	S3UsePathStyle interface{} `field:"optional" json:"s3UsePathStyle" yaml:"s3UsePathStyle"`
	// The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#secret_key AwsProvider#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// List of paths to shared config files. If not set, defaults to [~/.aws/config].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#shared_config_files AwsProvider#shared_config_files}
	SharedConfigFiles *[]*string `field:"optional" json:"sharedConfigFiles" yaml:"sharedConfigFiles"`
	// List of paths to shared credentials files. If not set, defaults to [~/.aws/credentials].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#shared_credentials_files AwsProvider#shared_credentials_files}
	SharedCredentialsFiles *[]*string `field:"optional" json:"sharedCredentialsFiles" yaml:"sharedCredentialsFiles"`
	// Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#skip_credentials_validation AwsProvider#skip_credentials_validation}
	SkipCredentialsValidation interface{} `field:"optional" json:"skipCredentialsValidation" yaml:"skipCredentialsValidation"`
	// Skip the AWS Metadata API check. Used for AWS API implementations that do not have a metadata api endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#skip_metadata_api_check AwsProvider#skip_metadata_api_check}
	SkipMetadataApiCheck *string `field:"optional" json:"skipMetadataApiCheck" yaml:"skipMetadataApiCheck"`
	// Skip static validation of region name.
	//
	// Used by users of alternative AWS-like APIs or users w/ access to regions that are not public (yet).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#skip_region_validation AwsProvider#skip_region_validation}
	SkipRegionValidation interface{} `field:"optional" json:"skipRegionValidation" yaml:"skipRegionValidation"`
	// Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#skip_requesting_account_id AwsProvider#skip_requesting_account_id}
	SkipRequestingAccountId interface{} `field:"optional" json:"skipRequestingAccountId" yaml:"skipRequestingAccountId"`
	// The region where AWS STS operations will take place. Examples are us-east-1 and us-west-2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#sts_region AwsProvider#sts_region}
	StsRegion *string `field:"optional" json:"stsRegion" yaml:"stsRegion"`
	// session token. A session token is only required if you are using temporary security credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#token AwsProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// The capacity of the AWS SDK's token bucket rate limiter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#token_bucket_rate_limiter_capacity AwsProvider#token_bucket_rate_limiter_capacity}
	TokenBucketRateLimiterCapacity *float64 `field:"optional" json:"tokenBucketRateLimiterCapacity" yaml:"tokenBucketRateLimiterCapacity"`
	// Resolve an endpoint with DualStack capability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#use_dualstack_endpoint AwsProvider#use_dualstack_endpoint}
	UseDualstackEndpoint interface{} `field:"optional" json:"useDualstackEndpoint" yaml:"useDualstackEndpoint"`
	// Resolve an endpoint with FIPS capability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs#use_fips_endpoint AwsProvider#use_fips_endpoint}
	UseFipsEndpoint interface{} `field:"optional" json:"useFipsEndpoint" yaml:"useFipsEndpoint"`
}

