package compose

import (
	"errors"
	"github.com/compose/gocomposeapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"net/url"
	"strings"
	"time"
)

func resourceDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentCreate,
		Read:   resourceDeploymentRead,
		Update: resourceDeploymentUpdate,
		Delete: resourceDeploymentDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_id": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"datacenter"},
			},
			"datacenter": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"cluster_id"},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"units": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"cache_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"wired_tiger": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"customer_billing_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			// "provisioning_tags": &schema.Schema{
			// Type:     schema.TypeList,
			// Elem:     &schema.Schema{Type: schema.TypeString},
			// Optional: true,
			// ForceNew: true,
			// },
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ca_certificate_base64": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"connection_details": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"admin_username": {
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
						},
						"admin_password": {
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
						},
						"sslmode": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"connection_strings": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health": {
							Type:      schema.TypeList,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Computed:  true,
							Optional:  true,
							Sensitive: true,
						},
						"ssh": {
							Type:      schema.TypeList,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Computed:  true,
							Optional:  true,
							Sensitive: true,
						},
						"admin": {
							Type:      schema.TypeList,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Computed:  true,
							Optional:  true,
							Sensitive: true,
						},
						"ssh_admin": {
							Type:      schema.TypeList,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Computed:  true,
							Optional:  true,
							Sensitive: true,
						},
						"cli": {
							Type:      schema.TypeList,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Computed:  true,
							Optional:  true,
							Sensitive: true,
						},
						"direct": {
							Type:      schema.TypeList,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Computed:  true,
							Optional:  true,
							Sensitive: true,
						},
					},
				},
			},
		},
	}
}

func resourceDeploymentCreate(d *schema.ResourceData, m interface{}) error {
	// return resourceDeploymentRead(d, m)
	c := m.(*composeapi.Client)
	deploymentParams := composeapi.DeploymentParams{
		Name:                d.Get("name").(string),
		AccountID:           d.Get("account_id").(string),
		Datacenter:          d.Get("datacenter").(string),
		DatabaseType:        d.Get("type").(string),
		Units:               d.Get("units").(int),
		Notes:               d.Get("notes").(string),
		CustomerBillingCode: d.Get("customer_billing_code").(string),
		Version:             d.Get("version").(string),
		// ProvisioningTags:    d.Get("provisioning_tags").([]string),
	}

	deployment, errs := c.CreateDeployment(deploymentParams)
	if errs != nil {
		return errs[0]
	}

	if err := waitForRecipe(c, deployment.ProvisionRecipeID); err != nil {
		return err
	}

	d.SetId(deployment.ID)

	return resourceDeploymentRead(d, m)
}

func resourceDeploymentRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*composeapi.Client)
	deployment, errs := c.GetDeployment(d.Id())
	if errs != nil {
		return errs[0]
	}

	if err := d.Set("name", deployment.Name); err != nil {
		return err
	}
	if err := d.Set("type", deployment.Type); err != nil {
		return err
	}
	if err := d.Set("created_at", deployment.CreatedAt.String()); err != nil {
		return err
	}
	if err := d.Set("ca_certificate_base64", deployment.CACertificateBase64); err != nil {
		return err
	}
	if err := d.Set("notes", deployment.Notes); err != nil {
		return err
	}
	if err := d.Set("customer_billing_code", deployment.CustomerBillingCode); err != nil {
		return err
	}
	if err := d.Set("version", deployment.Version); err != nil {
		return err
	}
	if err := d.Set("cluster_id", deployment.ClusterID); err != nil {
		return err
	}

	err := d.Set("connection_strings", []interface{}{
		map[string]interface{}{
			"health":    deployment.Connection.Health,
			"ssh":       deployment.Connection.SSH,
			"admin":     deployment.Connection.Admin,
			"ssh_admin": deployment.Connection.SSHAdmin,
			"cli":       deployment.Connection.CLI,
			"direct":    deployment.Connection.Direct,
		},
	})
	if err != nil {
		return err
	}

	connections := make([]interface{}, len(deployment.Connection.Direct))
	for i, connectionString := range deployment.Connection.Direct {
		// d.Set("")
		u, err := url.Parse(connectionString)
		if err != nil {
			return err
		}

		password, passwordExists := u.User.Password()
		if !passwordExists {
			return errors.New("No password was returned by compose.io")
		}
		connections[i] = map[string]interface{}{
			"admin_username": u.User.Username(),
			"admin_password": password,
			"host":           u.Hostname(),
			"port":           u.Port(),
			"database":       strings.TrimPrefix(u.Path, "/"),
			"sslmode":        "required",
		}
	}
	if err := d.Set("connection_details", connections); err != nil {
		return err
	}

	scalings, errs := c.GetScalings(d.Id())
	if errs != nil {
		return errs[0]
	}
	if err := d.Set("units", scalings.AllocatedUnits); err != nil {
		return err
	}

	// TODO: links
	return nil
}

func resourceDeploymentUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	c := m.(*composeapi.Client)

	if d.HasChange("version") {
		recipe, errs := c.UpdateVersion(d.Id(), d.Get("version").(string))
		if errs != nil {
			return errs[0]
		}
		if err := waitForRecipe(c, recipe.ID); err != nil {
			return err
		}
		d.SetPartial("version")
	}

	if d.HasChange("units") {
		scalingsParams := composeapi.ScalingsParams{
			DeploymentID: d.Id(),
			Units:        d.Get("units").(int),
		}
		recipe, errs := c.SetScalings(scalingsParams)
		if errs != nil {
			return errs[0]
		}
		if err := waitForRecipe(c, recipe.ID); err != nil {
			return err
		}
		d.SetPartial("units")
	}

	if d.HasChange("notes") || d.HasChange("customer_billing_code") {
		patchDeploymentParams := composeapi.PatchDeploymentParams{
			DeploymentID:        d.Id(),
			Notes:               d.Get("notes").(string),
			CustomerBillingCode: d.Get("customer_billing_code").(string),
		}
		_, errs := c.PatchDeployment(patchDeploymentParams)
		if errs != nil {
			return errs[0]
		}
		d.SetPartial("notes")
		d.SetPartial("customer_billing_code")
	}

	return resourceDeploymentRead(d, m)
}

func resourceDeploymentDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*composeapi.Client)

	deprovisionRecipe, errs := c.DeprovisionDeployment(d.Id())
	if errs != nil {
		return errs[0]
	}

	if err := waitForRecipe(c, deprovisionRecipe.ID); err != nil {
		return err
	}

	return nil
}

func waitForRecipe(c *composeapi.Client, recipeID string, timeoutOptional ...time.Duration) error {
	timeout := 30 * time.Minute
	if len(timeoutOptional) > 0 {
		timeout = timeoutOptional[0]
	}
	start := time.Now()

	for {
		now := time.Now()
		elapsed := now.Sub(start)
		if elapsed > timeout {
			return errors.New("Timeout error")
		}
		time.Sleep(5 * time.Second)

		recipe, errs := c.GetRecipe(recipeID)
		if errs != nil {
			return errs[0]
		}

		switch status := recipe.Status; status {
		case "complete":
			return nil
		case "failed":
			return errors.New(recipe.StatusDetail)
		default:
		}
	}
}
