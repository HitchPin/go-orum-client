package goorumclient

import (
	"context"
	"fmt"
	"github.com/HitchPin/go-orum-client/models"
	"github.com/HitchPin/go-orum-client/oauth"
	"github.com/HitchPin/go-orum-client/webhooks"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	bundle "github.com/microsoft/kiota-bundle-go"
	khttp "github.com/microsoft/kiota-http-go"
	"log"
	nethttp "net/http"
	"testing"
)

type VersionMiddleware struct {
	Version string
}

func (v VersionMiddleware) Intercept(pipeline khttp.Pipeline, i int, request *nethttp.Request) (*nethttp.Response, error) {
	request.Header.Set("Orum-Version", v.Version)
	log.Printf("Request to %s, %s", request.URL.RequestURI(), request.Method)
	res, err := pipeline.Next(request, i+1)
	log.Printf("Auth header: %s", request.Header.Get("Authorization"))
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		log.Printf("Response: %v\n", res.StatusCode)
	}
	return res, err
}

func TestAuth(t *testing.T) {
	authProvider := authentication.AnonymousAuthenticationProvider{}
	hclient := khttp.GetDefaultClient(VersionMiddleware{
		Version: "v2022-09-21",
	})
	adapter, err := bundle.NewDefaultRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		&authProvider,
		nil,
		nil,
		hclient,
	)
	if err != nil {
		log.Panicf("Error creating request adapter: %v\n", err)
	}
	adapter.SetBaseUrl("https://api-sandbox.orum.io")
	body := models.NewOauthTokenBody()
	clientId := "hc2lExstZNUD4tDdTN237H845tVJShgR"
	secretId := "o5s8p5nOu7alYSFm_m6gTA3aYOWdssOaW8pOz9eGNeJ3OTiFTQbbLxf-fMsdHzoA"
	body.SetClientId(&clientId)
	body.SetClientSecret(&secretId)
	authClient := NewOrumClient(adapter)
	authRes, err := authClient.Oauth().Token().Post(context.TODO(), body, &oauth.TokenRequestBuilderPostRequestConfiguration{})
	if err != nil {
		log.Panicf("Error obtaining access token: %v\n", err)
	}

	at := authRes.GetAccessToken()
	authTokenProvider, err := authentication.NewApiKeyAuthenticationProvider(
		fmt.Sprintf("Bearer %s", *at),
		"Authorization",
		authentication.HEADER_KEYLOCATION,
	)
	if err != nil {
		log.Panicf("Error creating api key auth provider: %v\n", err)
	}

	authAdapter, err := bundle.NewDefaultRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		authTokenProvider,
		nil,
		nil,
		hclient,
	)
	if err != nil {
		log.Panicf("Error creating authenticated request adapter: %v\n", err)
	}
	authAdapter.SetBaseUrl("https://api-sandbox.orum.io")
	authedClient := NewOrumClient(authAdapter)
	res, err := authedClient.Webhooks().Secret().Get(context.TODO(), &webhooks.SecretRequestBuilderGetRequestConfiguration{})
	//res, err := authedClient.Webhooks().Secret().Post(context.TODO(), &webhooks.SecretRequestBuilderPostRequestConfiguration{})
	if err != nil {
		log.Panicf("Error creating authenticated request adapter: %v\n", err)
	}
	log.Printf("Key: %v\n", *res.GetEnterpriseKeypair().GetPublicKey())
}
