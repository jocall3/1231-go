// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jocall3_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/jocall3/1231-go"
	"github.com/jocall3/1231-go/internal/testutil"
	"github.com/jocall3/1231-go/option"
)

func TestAIAdGenerateAdvancedWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := jocall3.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.AI.Ads.Generate.Advanced(context.TODO(), jocall3.AIAdGenerateAdvancedParams{
		LengthSeconds:        jocall3.F[any](15),
		Prompt:               jocall3.F[any]("A captivating ad featuring a young entrepreneur using 's AI tools to grow their startup. Focus on innovation and ease of use."),
		Style:                jocall3.F(jocall3.AIAdGenerateAdvancedParamsStyleCinematic),
		AspectRatio:          jocall3.F(jocall3.AIAdGenerateAdvancedParamsAspectRatio16_9),
		AudienceTarget:       jocall3.F(jocall3.AIAdGenerateAdvancedParamsAudienceTargetCorporate),
		BackgroundMusicGenre: jocall3.F(jocall3.AIAdGenerateAdvancedParamsBackgroundMusicGenreCorporate),
		BrandAssets:          jocall3.F([]interface{}{"https://demobank.com/assets/corporate_logo.png"}),
		BrandColors:          jocall3.F([]interface{}{"#0000FF", "#FFD700"}),
		CallToAction: jocall3.F(jocall3.AIAdGenerateAdvancedParamsCallToAction{
			DisplayTimeSeconds: jocall3.F[any](5),
			Text:               jocall3.F[any]("Learn more at DemoBank.com/business"),
			URL:                jocall3.F[any]("https://demobank.com/business"),
		}),
		Keywords:       jocall3.F([]interface{}{"innovation", "fintech", "startup"}),
		VoiceoverStyle: jocall3.F(jocall3.AIAdGenerateAdvancedParamsVoiceoverStyleMaleProfessional),
		VoiceoverText:  jocall3.F[any](": Your business, powered by intelligent finance."),
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAdGenerateStandardWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := jocall3.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.AI.Ads.Generate.Standard(context.TODO(), jocall3.AIAdGenerateStandardParams{
		GenerateVideo: jocall3.GenerateVideoParam{
			LengthSeconds: jocall3.F[any](15),
			Prompt:        jocall3.F[any]("A captivating ad featuring a young entrepreneur using 's AI tools to grow their startup. Focus on innovation and ease of use."),
			Style:         jocall3.F(jocall3.GenerateVideoStyleCinematic),
			AspectRatio:   jocall3.F(jocall3.GenerateVideoAspectRatio16_9),
			BrandColors:   jocall3.F([]interface{}{"#0000FF", "#FFD700"}),
			Keywords:      jocall3.F([]interface{}{"innovation", "fintech", "startup"}),
		},
	})
	if err != nil {
		var apierr *jocall3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
