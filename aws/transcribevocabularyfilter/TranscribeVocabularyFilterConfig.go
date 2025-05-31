// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcribevocabularyfilter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TranscribeVocabularyFilterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transcribe_vocabulary_filter#language_code TranscribeVocabularyFilter#language_code}.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transcribe_vocabulary_filter#vocabulary_filter_name TranscribeVocabularyFilter#vocabulary_filter_name}.
	VocabularyFilterName *string `field:"required" json:"vocabularyFilterName" yaml:"vocabularyFilterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transcribe_vocabulary_filter#id TranscribeVocabularyFilter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transcribe_vocabulary_filter#tags TranscribeVocabularyFilter#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transcribe_vocabulary_filter#tags_all TranscribeVocabularyFilter#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transcribe_vocabulary_filter#vocabulary_filter_file_uri TranscribeVocabularyFilter#vocabulary_filter_file_uri}.
	VocabularyFilterFileUri *string `field:"optional" json:"vocabularyFilterFileUri" yaml:"vocabularyFilterFileUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/transcribe_vocabulary_filter#words TranscribeVocabularyFilter#words}.
	Words *[]*string `field:"optional" json:"words" yaml:"words"`
}

