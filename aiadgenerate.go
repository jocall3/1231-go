// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jocall3

import (
	"context"
	"net/http"
	"slices"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// AIAdGenerateService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAdGenerateService] method instead.
type AIAdGenerateService struct {
	Options []option.RequestOption
}

// NewAIAdGenerateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAdGenerateService(opts ...option.RequestOption) (r *AIAdGenerateService) {
	r = &AIAdGenerateService{}
	r.Options = opts
	return
}

// Submits a highly customized request to generate a video ad, allowing
// fine-grained control over artistic style, aspect ratio, voiceover, background
// music, target audience, and call-to-action elements for professional-grade
// productions.
func (r *AIAdGenerateService) Advanced(ctx context.Context, body AIAdGenerateAdvancedParams, opts ...option.RequestOption) (res *AIAdGenerateAdvancedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/ads/generate/advanced"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Submits a request to generate a high-quality video ad using the advanced Veo 2.0
// generative AI model. This is an asynchronous operation, suitable for standard ad
// content creation.
func (r *AIAdGenerateService) Standard(ctx context.Context, body AIAdGenerateStandardParams, opts ...option.RequestOption) (res *AIAdGenerateStandardResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/ads/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GenerateVideoParam struct {
	// Desired length of the video in seconds.
	LengthSeconds param.Field[interface{}] `json:"lengthSeconds,required"`
	// The textual prompt to guide the AI video generation.
	Prompt param.Field[interface{}] `json:"prompt,required"`
	// Artistic style of the video.
	Style param.Field[GenerateVideoStyle] `json:"style,required"`
	// Aspect ratio of the video (e.g., 16:9 for widescreen, 9:16 for vertical shorts).
	AspectRatio param.Field[GenerateVideoAspectRatio] `json:"aspectRatio"`
	// Optional: Hex color codes to influence the video's aesthetic.
	BrandColors param.Field[[]interface{}] `json:"brandColors"`
	// Optional: Additional keywords to guide the AI's content generation.
	Keywords param.Field[[]interface{}] `json:"keywords"`
}

func (r GenerateVideoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Artistic style of the video.
type GenerateVideoStyle string

const (
	GenerateVideoStyleCinematic   GenerateVideoStyle = "Cinematic"
	GenerateVideoStyleExplainer   GenerateVideoStyle = "Explainer"
	GenerateVideoStyleDocumentary GenerateVideoStyle = "Documentary"
	GenerateVideoStyleAbstract    GenerateVideoStyle = "Abstract"
	GenerateVideoStyleMinimalist  GenerateVideoStyle = "Minimalist"
)

func (r GenerateVideoStyle) IsKnown() bool {
	switch r {
	case GenerateVideoStyleCinematic, GenerateVideoStyleExplainer, GenerateVideoStyleDocumentary, GenerateVideoStyleAbstract, GenerateVideoStyleMinimalist:
		return true
	}
	return false
}

// Aspect ratio of the video (e.g., 16:9 for widescreen, 9:16 for vertical shorts).
type GenerateVideoAspectRatio string

const (
	GenerateVideoAspectRatio16_9 GenerateVideoAspectRatio = "16:9"
	GenerateVideoAspectRatio9_16 GenerateVideoAspectRatio = "9:16"
	GenerateVideoAspectRatio1_1  GenerateVideoAspectRatio = "1:1"
)

func (r GenerateVideoAspectRatio) IsKnown() bool {
	switch r {
	case GenerateVideoAspectRatio16_9, GenerateVideoAspectRatio9_16, GenerateVideoAspectRatio1_1:
		return true
	}
	return false
}

type AIAdGenerateAdvancedResponse struct {
	// Estimated time until advanced video generation is complete. May be longer than
	// standard generation.
	EstimatedCompletionTimeSeconds interface{} `json:"estimatedCompletionTimeSeconds"`
	// The unique identifier for the advanced video generation operation.
	OperationID interface{}                      `json:"operationId"`
	JSON        aiAdGenerateAdvancedResponseJSON `json:"-"`
}

// aiAdGenerateAdvancedResponseJSON contains the JSON metadata for the struct
// [AIAdGenerateAdvancedResponse]
type aiAdGenerateAdvancedResponseJSON struct {
	EstimatedCompletionTimeSeconds apijson.Field
	OperationID                    apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AIAdGenerateAdvancedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdGenerateAdvancedResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdGenerateStandardResponse struct {
	// Estimated time until video generation is complete.
	EstimatedCompletionTimeSeconds interface{} `json:"estimatedCompletionTimeSeconds"`
	// The unique identifier for the video generation operation.
	OperationID interface{}                      `json:"operationId"`
	JSON        aiAdGenerateStandardResponseJSON `json:"-"`
}

// aiAdGenerateStandardResponseJSON contains the JSON metadata for the struct
// [AIAdGenerateStandardResponse]
type aiAdGenerateStandardResponseJSON struct {
	EstimatedCompletionTimeSeconds apijson.Field
	OperationID                    apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AIAdGenerateStandardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdGenerateStandardResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdGenerateAdvancedParams struct {
	// Desired length of the video in seconds.
	LengthSeconds param.Field[interface{}] `json:"lengthSeconds,required"`
	// The textual prompt to guide the AI video generation.
	Prompt param.Field[interface{}] `json:"prompt,required"`
	// Artistic style of the video.
	Style param.Field[AIAdGenerateAdvancedParamsStyle] `json:"style,required"`
	// Aspect ratio of the video (e.g., 16:9 for widescreen, 9:16 for vertical shorts).
	AspectRatio param.Field[AIAdGenerateAdvancedParamsAspectRatio] `json:"aspectRatio"`
	// Target audience for the ad, influencing tone and visuals.
	AudienceTarget param.Field[AIAdGenerateAdvancedParamsAudienceTarget] `json:"audienceTarget"`
	// Genre of background music.
	BackgroundMusicGenre param.Field[AIAdGenerateAdvancedParamsBackgroundMusicGenre] `json:"backgroundMusicGenre"`
	// URLs to brand assets (e.g., logos, specific imagery) to be incorporated.
	BrandAssets param.Field[[]interface{}] `json:"brandAssets"`
	// Optional: Hex color codes to influence the video's aesthetic.
	BrandColors param.Field[[]interface{}] `json:"brandColors"`
	// Call-to-action text and URL to be displayed.
	CallToAction param.Field[AIAdGenerateAdvancedParamsCallToAction] `json:"callToAction"`
	// Optional: Additional keywords to guide the AI's content generation.
	Keywords param.Field[[]interface{}] `json:"keywords"`
	// Style/tone for the AI voiceover.
	VoiceoverStyle param.Field[AIAdGenerateAdvancedParamsVoiceoverStyle] `json:"voiceoverStyle"`
	// Optional: Text for an AI-generated voiceover.
	VoiceoverText param.Field[interface{}] `json:"voiceoverText"`
}

func (r AIAdGenerateAdvancedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Artistic style of the video.
type AIAdGenerateAdvancedParamsStyle string

const (
	AIAdGenerateAdvancedParamsStyleCinematic   AIAdGenerateAdvancedParamsStyle = "Cinematic"
	AIAdGenerateAdvancedParamsStyleExplainer   AIAdGenerateAdvancedParamsStyle = "Explainer"
	AIAdGenerateAdvancedParamsStyleDocumentary AIAdGenerateAdvancedParamsStyle = "Documentary"
	AIAdGenerateAdvancedParamsStyleAbstract    AIAdGenerateAdvancedParamsStyle = "Abstract"
	AIAdGenerateAdvancedParamsStyleMinimalist  AIAdGenerateAdvancedParamsStyle = "Minimalist"
)

func (r AIAdGenerateAdvancedParamsStyle) IsKnown() bool {
	switch r {
	case AIAdGenerateAdvancedParamsStyleCinematic, AIAdGenerateAdvancedParamsStyleExplainer, AIAdGenerateAdvancedParamsStyleDocumentary, AIAdGenerateAdvancedParamsStyleAbstract, AIAdGenerateAdvancedParamsStyleMinimalist:
		return true
	}
	return false
}

// Aspect ratio of the video (e.g., 16:9 for widescreen, 9:16 for vertical shorts).
type AIAdGenerateAdvancedParamsAspectRatio string

const (
	AIAdGenerateAdvancedParamsAspectRatio16_9 AIAdGenerateAdvancedParamsAspectRatio = "16:9"
	AIAdGenerateAdvancedParamsAspectRatio9_16 AIAdGenerateAdvancedParamsAspectRatio = "9:16"
	AIAdGenerateAdvancedParamsAspectRatio1_1  AIAdGenerateAdvancedParamsAspectRatio = "1:1"
)

func (r AIAdGenerateAdvancedParamsAspectRatio) IsKnown() bool {
	switch r {
	case AIAdGenerateAdvancedParamsAspectRatio16_9, AIAdGenerateAdvancedParamsAspectRatio9_16, AIAdGenerateAdvancedParamsAspectRatio1_1:
		return true
	}
	return false
}

// Target audience for the ad, influencing tone and visuals.
type AIAdGenerateAdvancedParamsAudienceTarget string

const (
	AIAdGenerateAdvancedParamsAudienceTargetGeneral   AIAdGenerateAdvancedParamsAudienceTarget = "general"
	AIAdGenerateAdvancedParamsAudienceTargetCorporate AIAdGenerateAdvancedParamsAudienceTarget = "corporate"
	AIAdGenerateAdvancedParamsAudienceTargetInvestor  AIAdGenerateAdvancedParamsAudienceTarget = "investor"
	AIAdGenerateAdvancedParamsAudienceTargetYouth     AIAdGenerateAdvancedParamsAudienceTarget = "youth"
)

func (r AIAdGenerateAdvancedParamsAudienceTarget) IsKnown() bool {
	switch r {
	case AIAdGenerateAdvancedParamsAudienceTargetGeneral, AIAdGenerateAdvancedParamsAudienceTargetCorporate, AIAdGenerateAdvancedParamsAudienceTargetInvestor, AIAdGenerateAdvancedParamsAudienceTargetYouth:
		return true
	}
	return false
}

// Genre of background music.
type AIAdGenerateAdvancedParamsBackgroundMusicGenre string

const (
	AIAdGenerateAdvancedParamsBackgroundMusicGenreCorporate AIAdGenerateAdvancedParamsBackgroundMusicGenre = "corporate"
	AIAdGenerateAdvancedParamsBackgroundMusicGenreUplifting AIAdGenerateAdvancedParamsBackgroundMusicGenre = "uplifting"
	AIAdGenerateAdvancedParamsBackgroundMusicGenreAmbient   AIAdGenerateAdvancedParamsBackgroundMusicGenre = "ambient"
	AIAdGenerateAdvancedParamsBackgroundMusicGenreCinematic AIAdGenerateAdvancedParamsBackgroundMusicGenre = "cinematic"
	AIAdGenerateAdvancedParamsBackgroundMusicGenreNone      AIAdGenerateAdvancedParamsBackgroundMusicGenre = "none"
)

func (r AIAdGenerateAdvancedParamsBackgroundMusicGenre) IsKnown() bool {
	switch r {
	case AIAdGenerateAdvancedParamsBackgroundMusicGenreCorporate, AIAdGenerateAdvancedParamsBackgroundMusicGenreUplifting, AIAdGenerateAdvancedParamsBackgroundMusicGenreAmbient, AIAdGenerateAdvancedParamsBackgroundMusicGenreCinematic, AIAdGenerateAdvancedParamsBackgroundMusicGenreNone:
		return true
	}
	return false
}

// Call-to-action text and URL to be displayed.
type AIAdGenerateAdvancedParamsCallToAction struct {
	DisplayTimeSeconds param.Field[interface{}] `json:"displayTimeSeconds"`
	Text               param.Field[interface{}] `json:"text"`
	URL                param.Field[interface{}] `json:"url"`
}

func (r AIAdGenerateAdvancedParamsCallToAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Style/tone for the AI voiceover.
type AIAdGenerateAdvancedParamsVoiceoverStyle string

const (
	AIAdGenerateAdvancedParamsVoiceoverStyleMaleProfessional AIAdGenerateAdvancedParamsVoiceoverStyle = "male_professional"
	AIAdGenerateAdvancedParamsVoiceoverStyleFemaleFriendly   AIAdGenerateAdvancedParamsVoiceoverStyle = "female_friendly"
	AIAdGenerateAdvancedParamsVoiceoverStyleNeutralCalm      AIAdGenerateAdvancedParamsVoiceoverStyle = "neutral_calm"
)

func (r AIAdGenerateAdvancedParamsVoiceoverStyle) IsKnown() bool {
	switch r {
	case AIAdGenerateAdvancedParamsVoiceoverStyleMaleProfessional, AIAdGenerateAdvancedParamsVoiceoverStyleFemaleFriendly, AIAdGenerateAdvancedParamsVoiceoverStyleNeutralCalm:
		return true
	}
	return false
}

type AIAdGenerateStandardParams struct {
	GenerateVideo GenerateVideoParam `json:"generate_video,required"`
}

func (r AIAdGenerateStandardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.GenerateVideo)
}
