package tomba

import (
	"context"
	"errors"
	"os"

	"github.com/tomba-io/go/tomba"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*tomba.Tomba, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "tomba"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*tomba.Tomba), nil
	}
	// Default to the env var settings
	key := os.Getenv("TOMBA_KEY")
	secret := os.Getenv("TOMBA_SECRET")

	// Prefer config settings
	tombaConfig := GetConfig(d.Connection)
	if tombaConfig.Key != nil {
		key = *tombaConfig.Key
	}
	if tombaConfig.Secret != nil {
		secret = *tombaConfig.Secret
	}

	// Error if the minimum config is not set
	if key == "" {
		return nil, errors.New("key must be configured")
	}
	if secret == "" {
		return nil, errors.New("secret must be configured")
	}

	plugin.Logger(ctx).Debug("tomba.connect", "config", tombaConfig)
	conn := tomba.New(key, secret)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func urlString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	q := quals["url"].GetStringValue()
	return q, nil
}
func domainString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	q := quals["domain"].GetStringValue()
	return q, nil
}
