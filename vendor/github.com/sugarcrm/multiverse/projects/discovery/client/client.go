package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	v1 "github.com/sugarcrm/multiverse/apis/discovery/v1"
)

const (
	URL_PRODUCTION = "https://discovery.services.sugarcrm.com"
	VERSION        = "v1"
)

type Client struct {
	// The base discovery URL http://discovery/v1
	url string

	// HTTP client object
	hc *http.Client

	// Discovery objects
	Kinds    *v1.Kinds
	Regions  *v1.Regions
	Services *v1.Services
}

// Create new instance
func New(url string) *Client {
	return &Client{
		url:      url + "/" + VERSION,
		Kinds:    &v1.Kinds{},
		Regions:  &v1.Regions{},
		Services: &v1.Services{},
	}
}

// Get active partition
func (c *Client) GetPartition() string {
	return c.Kinds.Partition
}

// Get all regions
func (c *Client) GetRegions() []*v1.Region {
	return c.Regions.Regions
}

// Get all services
func (c *Client) GetServices() []*v1.Service {
	return c.Services.Services
}

// Get endpoint for given service and region.
func (c *Client) GetEndpoint(name, region string) (*v1.ServiceEndpoint, error) {
	for _, s := range c.GetServices() {
		if s.Name == name {
			for _, e := range s.Endpoints {
				if e.Region == region {
					return e, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("Service %s for region %s not found", name, region)
}

// Load discovery objects from file
func (c *Client) LoadFromFile(file string) error {
	// TODO
	return nil
}

// Enable auto refresh. This function returns the ticker so
// auto refresh can be stopped by calling ticker.Stop().
func (c *Client) AutoRefresh(interval time.Duration) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for _ = range ticker.C {
			err := c.Refresh(context.Background())
			if err != nil {
				log.Warnf("Discovery auto refresh error: %v", err)
			} else {
				log.Infof("Discovery auto refresh success")
			}
		}
	}()
	return ticker
}

// Refresh all objects
func (c *Client) Refresh(ctx context.Context) error {
	if err := c.RefreshKinds(ctx); err != nil {
		return fmt.Errorf("Error refreshing kinds: %v", err)
	}
	if err := c.RefreshRegions(ctx); err != nil {
		return fmt.Errorf("Error refreshing regions: %v", err)
	}
	if err := c.RefreshServices(ctx); err != nil {
		return fmt.Errorf("Error refreshing services: %v", err)
	}
	return nil
}

// Refresh kinds
func (c *Client) RefreshKinds(ctx context.Context) error {
	return c.refresh(ctx, "", c.Kinds)
}

// Refresh regions
func (c *Client) RefreshRegions(ctx context.Context) error {
	return c.refresh(ctx, "/regions", c.Regions)
}

// Refresh services
func (c *Client) RefreshServices(ctx context.Context) error {
	return c.refresh(ctx, "/services", c.Services)
}

// Low level HTTP request wrapper
func (c *Client) refresh(ctx context.Context, url string, object interface{}) error {

	// Check for discovery URL
	if c.url == "" {
		return fmt.Errorf("No discovery URL set")
	}

	// Create request object
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.url, url), nil)
	if err != nil {
		return fmt.Errorf("Cannot create request: %v", err)
	}

	// Lazy load http client
	if c.hc == nil {
		c.hc = &http.Client{}
	}

	// Perform request using context
	req.WithContext(ctx)
	res, err := c.hc.Do(req)
	if err != nil {
		return err
	}

	// Unmarshal result
	body, _ := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(body, object); err != nil {
		return err
	}

	return nil
}
