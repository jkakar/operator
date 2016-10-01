package openstack

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/networking/v2/subnets"
)

func resourceNetworkingSubnetV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkingSubnetV2Create,
		Read:   resourceNetworkingSubnetV2Read,
		Update: resourceNetworkingSubnetV2Update,
		Delete: resourceNetworkingSubnetV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_REGION_NAME", ""),
			},
			"network_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"allocation_pools": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start": {
							Type:     schema.TypeString,
							Required: true,
						},
						"end": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"gateway_ip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Computed: true,
			},
			"no_gateway": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
			},
			"ip_version": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
				ForceNew: true,
			},
			"enable_dhcp": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  true,
			},
			"dns_nameservers": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: false,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"host_routes": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_cidr": {
							Type:     schema.TypeString,
							Required: true,
						},
						"next_hop": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"value_specs": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

// SubnetCreateOpts represents the attributes used when creating a new subnet.
type SubnetCreateOpts struct {
	// Required
	NetworkID string
	CIDR      string
	// Optional
	Name            string
	TenantID        string
	AllocationPools []subnets.AllocationPool
	GatewayIP       string
	NoGateway       bool
	IPVersion       int
	EnableDHCP      *bool
	DNSNameservers  []string
	HostRoutes      []subnets.HostRoute
	ValueSpecs      map[string]string
}

// ToSubnetCreateMap casts a CreateOpts struct to a map.
func (opts SubnetCreateOpts) ToSubnetCreateMap() (map[string]interface{}, error) {
	s := make(map[string]interface{})

	if opts.NetworkID == "" {
		return nil, fmt.Errorf("A network ID is required")
	}
	if opts.CIDR == "" {
		return nil, fmt.Errorf("A valid CIDR is required")
	}
	if opts.IPVersion != 0 && opts.IPVersion != subnets.IPv4 && opts.IPVersion != subnets.IPv6 {
		return nil, fmt.Errorf("An IP type must either be 4 or 6")
	}

	// Both GatewayIP and NoGateway should not be set
	if opts.GatewayIP != "" && opts.NoGateway {
		return nil, fmt.Errorf("Both disabling the gateway and specifying a gateway is not allowed")
	}

	s["network_id"] = opts.NetworkID
	s["cidr"] = opts.CIDR

	if opts.EnableDHCP != nil {
		s["enable_dhcp"] = &opts.EnableDHCP
	}
	if opts.Name != "" {
		s["name"] = opts.Name
	}
	if opts.GatewayIP != "" {
		s["gateway_ip"] = opts.GatewayIP
	} else if opts.NoGateway {
		s["gateway_ip"] = nil
	}
	if opts.TenantID != "" {
		s["tenant_id"] = opts.TenantID
	}
	if opts.IPVersion != 0 {
		s["ip_version"] = opts.IPVersion
	}
	if len(opts.AllocationPools) != 0 {
		s["allocation_pools"] = opts.AllocationPools
	}
	if len(opts.DNSNameservers) != 0 {
		s["dns_nameservers"] = opts.DNSNameservers
	}
	if len(opts.HostRoutes) != 0 {
		s["host_routes"] = opts.HostRoutes
	}

	if opts.ValueSpecs != nil {
		for k, v := range opts.ValueSpecs {
			s[k] = v
		}
	}

	return map[string]interface{}{"subnet": s}, nil
}

func resourceNetworkingSubnetV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(d.Get("region").(string))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	if _, ok := d.GetOk("gateway_ip"); ok {
		if _, ok2 := d.GetOk("no_gateway"); ok2 {
			return fmt.Errorf("Both gateway_ip and no_gateway cannot be set.")
		}
	}

	enableDHCP := d.Get("enable_dhcp").(bool)

	createOpts := SubnetCreateOpts{
		NetworkID:       d.Get("network_id").(string),
		CIDR:            d.Get("cidr").(string),
		Name:            d.Get("name").(string),
		TenantID:        d.Get("tenant_id").(string),
		AllocationPools: resourceSubnetAllocationPoolsV2(d),
		GatewayIP:       d.Get("gateway_ip").(string),
		NoGateway:       d.Get("no_gateway").(bool),
		IPVersion:       d.Get("ip_version").(int),
		DNSNameservers:  resourceSubnetDNSNameserversV2(d),
		HostRoutes:      resourceSubnetHostRoutesV2(d),
		EnableDHCP:      &enableDHCP,
		ValueSpecs:      subnetValueSpecs(d),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	s, err := subnets.Create(networkingClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating OpenStack Neutron subnet: %s", err)
	}
	log.Printf("[INFO] Subnet ID: %s", s.ID)

	log.Printf("[DEBUG] Waiting for Subnet (%s) to become available", s.ID)
	stateConf := &resource.StateChangeConf{
		Target:     []string{"ACTIVE"},
		Refresh:    waitForSubnetActive(networkingClient, s.ID),
		Timeout:    10 * time.Minute,
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()

	d.SetId(s.ID)

	return resourceNetworkingSubnetV2Read(d, meta)
}

func resourceNetworkingSubnetV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(d.Get("region").(string))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	s, err := subnets.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "subnet")
	}

	log.Printf("[DEBUG] Retreived Subnet %s: %+v", d.Id(), s)

	d.Set("newtork_id", s.NetworkID)
	d.Set("cidr", s.CIDR)
	d.Set("ip_version", s.IPVersion)
	d.Set("name", s.Name)
	d.Set("tenant_id", s.TenantID)
	d.Set("allocation_pools", s.AllocationPools)
	d.Set("gateway_ip", s.GatewayIP)
	d.Set("dns_nameservers", s.DNSNameservers)
	d.Set("host_routes", s.HostRoutes)
	d.Set("enable_dhcp", s.EnableDHCP)
	d.Set("network_id", s.NetworkID)

	return nil
}

func resourceNetworkingSubnetV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(d.Get("region").(string))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	// Check if both gateway_ip and no_gateway are set
	if _, ok := d.GetOk("gateway_ip"); ok {
		if _, ok2 := d.GetOk("no_gateway"); ok2 {
			return fmt.Errorf("Both gateway_ip and no_gateway cannot be set.")
		}
	}

	var updateOpts subnets.UpdateOpts

	if d.HasChange("name") {
		updateOpts.Name = d.Get("name").(string)
	}

	if d.HasChange("gateway_ip") {
		updateOpts.GatewayIP = d.Get("gateway_ip").(string)
	}

	if d.HasChange("no_gateway") {
		updateOpts.NoGateway = d.Get("no_gateway").(bool)
	}

	if d.HasChange("dns_nameservers") {
		updateOpts.DNSNameservers = resourceSubnetDNSNameserversV2(d)
	}

	if d.HasChange("host_routes") {
		updateOpts.HostRoutes = resourceSubnetHostRoutesV2(d)
	}

	if d.HasChange("enable_dhcp") {
		v := d.Get("enable_dhcp").(bool)
		updateOpts.EnableDHCP = &v
	}

	log.Printf("[DEBUG] Updating Subnet %s with options: %+v", d.Id(), updateOpts)

	_, err = subnets.Update(networkingClient, d.Id(), updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating OpenStack Neutron Subnet: %s", err)
	}

	return resourceNetworkingSubnetV2Read(d, meta)
}

func resourceNetworkingSubnetV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(d.Get("region").(string))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE"},
		Target:     []string{"DELETED"},
		Refresh:    waitForSubnetDelete(networkingClient, d.Id()),
		Timeout:    10 * time.Minute,
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error deleting OpenStack Neutron Subnet: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceSubnetAllocationPoolsV2(d *schema.ResourceData) []subnets.AllocationPool {
	rawAPs := d.Get("allocation_pools").([]interface{})
	aps := make([]subnets.AllocationPool, len(rawAPs))
	for i, raw := range rawAPs {
		rawMap := raw.(map[string]interface{})
		aps[i] = subnets.AllocationPool{
			Start: rawMap["start"].(string),
			End:   rawMap["end"].(string),
		}
	}
	return aps
}

func resourceSubnetDNSNameserversV2(d *schema.ResourceData) []string {
	rawDNSN := d.Get("dns_nameservers").(*schema.Set)
	dnsn := make([]string, rawDNSN.Len())
	for i, raw := range rawDNSN.List() {
		dnsn[i] = raw.(string)
	}
	return dnsn
}

func resourceSubnetHostRoutesV2(d *schema.ResourceData) []subnets.HostRoute {
	rawHR := d.Get("host_routes").([]interface{})
	hr := make([]subnets.HostRoute, len(rawHR))
	for i, raw := range rawHR {
		rawMap := raw.(map[string]interface{})
		hr[i] = subnets.HostRoute{
			DestinationCIDR: rawMap["destination_cidr"].(string),
			NextHop:         rawMap["next_hop"].(string),
		}
	}
	return hr
}

func waitForSubnetActive(networkingClient *gophercloud.ServiceClient, subnetId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		s, err := subnets.Get(networkingClient, subnetId).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] OpenStack Neutron Subnet: %+v", s)
		return s, "ACTIVE", nil
	}
}

func waitForSubnetDelete(networkingClient *gophercloud.ServiceClient, subnetId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete OpenStack Subnet %s.\n", subnetId)

		s, err := subnets.Get(networkingClient, subnetId).Extract()
		if err != nil {
			errCode, ok := err.(*gophercloud.UnexpectedResponseCodeError)
			if !ok {
				return s, "ACTIVE", err
			}
			if errCode.Actual == 404 {
				log.Printf("[DEBUG] Successfully deleted OpenStack Subnet %s", subnetId)
				return s, "DELETED", nil
			}
		}

		err = subnets.Delete(networkingClient, subnetId).ExtractErr()
		if err != nil {
			errCode, ok := err.(*gophercloud.UnexpectedResponseCodeError)
			if !ok {
				return s, "ACTIVE", err
			}
			if errCode.Actual == 404 {
				log.Printf("[DEBUG] Successfully deleted OpenStack Subnet %s", subnetId)
				return s, "DELETED", nil
			}
		}

		log.Printf("[DEBUG] OpenStack Subnet %s still active.\n", subnetId)
		return s, "ACTIVE", nil
	}
}

func subnetValueSpecs(d *schema.ResourceData) map[string]string {
	m := make(map[string]string)
	for key, val := range d.Get("value_specs").(map[string]interface{}) {
		m[key] = val.(string)
	}
	return m
}
