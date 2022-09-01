package s3


type S3BucketCorsConfigurationCorsRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_cors_configuration#allowed_methods S3BucketCorsConfiguration#allowed_methods}.
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_cors_configuration#allowed_origins S3BucketCorsConfiguration#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_cors_configuration#allowed_headers S3BucketCorsConfiguration#allowed_headers}.
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_cors_configuration#expose_headers S3BucketCorsConfiguration#expose_headers}.
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_cors_configuration#id S3BucketCorsConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_cors_configuration#max_age_seconds S3BucketCorsConfiguration#max_age_seconds}.
	MaxAgeSeconds *float64 `field:"optional" json:"maxAgeSeconds" yaml:"maxAgeSeconds"`
}

