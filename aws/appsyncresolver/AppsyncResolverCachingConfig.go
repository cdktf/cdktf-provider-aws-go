package appsyncresolver


type AppsyncResolverCachingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/appsync_resolver#caching_keys AppsyncResolver#caching_keys}.
	CachingKeys *[]*string `field:"optional" json:"cachingKeys" yaml:"cachingKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/appsync_resolver#ttl AppsyncResolver#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

