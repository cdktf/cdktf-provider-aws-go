package lexbot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LexBotConfig struct {
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
	// abort_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#abort_statement LexBot#abort_statement}
	AbortStatement *LexBotAbortStatement `field:"required" json:"abortStatement" yaml:"abortStatement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#child_directed LexBot#child_directed}.
	ChildDirected interface{} `field:"required" json:"childDirected" yaml:"childDirected"`
	// intent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#intent LexBot#intent}
	Intent interface{} `field:"required" json:"intent" yaml:"intent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#name LexBot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// clarification_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#clarification_prompt LexBot#clarification_prompt}
	ClarificationPrompt *LexBotClarificationPrompt `field:"optional" json:"clarificationPrompt" yaml:"clarificationPrompt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#create_version LexBot#create_version}.
	CreateVersion interface{} `field:"optional" json:"createVersion" yaml:"createVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#description LexBot#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#detect_sentiment LexBot#detect_sentiment}.
	DetectSentiment interface{} `field:"optional" json:"detectSentiment" yaml:"detectSentiment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#enable_model_improvements LexBot#enable_model_improvements}.
	EnableModelImprovements interface{} `field:"optional" json:"enableModelImprovements" yaml:"enableModelImprovements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#id LexBot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#idle_session_ttl_in_seconds LexBot#idle_session_ttl_in_seconds}.
	IdleSessionTtlInSeconds *float64 `field:"optional" json:"idleSessionTtlInSeconds" yaml:"idleSessionTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#locale LexBot#locale}.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#nlu_intent_confidence_threshold LexBot#nlu_intent_confidence_threshold}.
	NluIntentConfidenceThreshold *float64 `field:"optional" json:"nluIntentConfidenceThreshold" yaml:"nluIntentConfidenceThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#process_behavior LexBot#process_behavior}.
	ProcessBehavior *string `field:"optional" json:"processBehavior" yaml:"processBehavior"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#timeouts LexBot#timeouts}
	Timeouts *LexBotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lex_bot#voice_id LexBot#voice_id}.
	VoiceId *string `field:"optional" json:"voiceId" yaml:"voiceId"`
}

