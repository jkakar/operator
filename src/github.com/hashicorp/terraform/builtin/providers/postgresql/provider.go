package postgresql

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("POSTGRESQL_HOST", nil),
				Description: "The postgresql server address",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     5432,
				Description: "The postgresql server port",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("POSTGRESQL_USERNAME", nil),
				Description: "Username for postgresql server connection",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("POSTGRESQL_PASSWORD", nil),
				Description: "Password for postgresql server connection",
			},
			"ssl_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "prefer",
				Description: "Connection mode for postgresql server",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"postgresql_database": resourcePostgresqlDatabase(),
			"postgresql_role":     resourcePostgresqlRole(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Host:     d.Get("host").(string),
		Port:     d.Get("port").(int),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		SslMode:  d.Get("ssl_mode").(string),
	}

	client, err := config.NewClient()
	if err != nil {
		return nil, fmt.Errorf("Error initializing Postgresql client: %s", err)
	}

	return client, nil
}
