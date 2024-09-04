/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/xunholy/provider-github/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal github credentials as JSON"
	errInvalidConfig        = "invalid provider configuration"

	// provider config keys
	keyBaseURL         = "base_url"
	keyOwner           = "owner"
	keyToken           = "token"
	keyAppAuth         = "app_auth"
	keyWriteDelayMs    = "write_delay_ms"
	keyReadDelayMs     = "read_delay_ms"
	keyRetryDelayMs    = "retry_delay_ms"
	keyMaxRetries      = "max_retries"
	keyRetryableErrors = "retryable_errors"
)

type appAuth struct {
	ID             string `json:"id"`
	InstallationID string `json:"installation_id"`
	PemFile        string `json:"pem_file"`
}

type githubConfig struct {
	BaseURL         *string    `json:"base_url,omitempty"`
	Owner           *string    `json:"owner,omitempty"`
	Token           *string    `json:"token,omitempty"`
	AppAuth         *[]appAuth `json:"app_auth,omitempty"`
	WriteDelayMs    *int       `json:"write_delay_ms,omitempty"`
	ReadDelayMs     *int       `json:"read_delay_ms,omitempty"`
	RetryDelayMs    *int       `json:"retry_delay_ms,omitempty"`
	MaxRetries      *int       `json:"max_retries,omitempty"`
	RetryableErrors []int      `json:"retryable_errors,omitempty"`
}

func (c *githubConfig) UnmarshalJSON(data []byte) error {
	type Alias githubConfig
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// Set default values here if needed
	if c.MaxRetries == nil {
		defaultMaxRetries := 3
		c.MaxRetries = &defaultMaxRetries
	}
	return nil
}

func (c githubConfig) validate() error {
	if c.Token == nil && c.AppAuth == nil {
		return errors.New("either token or app_auth must be provided")
	}
	return nil
}

func (c githubConfig) getStringValue(key string) string {
	switch key {
	case keyBaseURL:
		if c.BaseURL != nil {
			return *c.BaseURL
		}
	case keyOwner:
		if c.Owner != nil {
			return *c.Owner
		}
	case keyToken:
		if c.Token != nil {
			return *c.Token
		}
	}
	return ""
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}

		config := &githubConfig{}
		if err := json.Unmarshal(data, config); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		if err := config.validate(); err != nil {
			return ps, errors.Wrap(err, errInvalidConfig)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{}
		if v := config.getStringValue(keyBaseURL); v != "" {
			ps.Configuration[keyBaseURL] = v
		}
		if v := config.getStringValue(keyOwner); v != "" {
			ps.Configuration[keyOwner] = v
		}
		if v := config.getStringValue(keyToken); v != "" {
			ps.Configuration[keyToken] = v
		}
		if config.AppAuth != nil {
			ps.Configuration[keyAppAuth] = *config.AppAuth
		}
		if config.WriteDelayMs != nil {
			ps.Configuration[keyWriteDelayMs] = *config.WriteDelayMs
		}
		if config.ReadDelayMs != nil {
			ps.Configuration[keyReadDelayMs] = *config.ReadDelayMs
		}
		if config.RetryDelayMs != nil {
			ps.Configuration[keyRetryDelayMs] = *config.RetryDelayMs
		}
		if config.MaxRetries != nil {
			ps.Configuration[keyMaxRetries] = *config.MaxRetries
		}
		if len(config.RetryableErrors) > 0 {
			ps.Configuration[keyRetryableErrors] = config.RetryableErrors
		}

		return ps, nil
	}
}
