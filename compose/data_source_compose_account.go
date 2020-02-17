package compose

import (
	"github.com/compose/gocomposeapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAccountRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAccountRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*composeapi.Client)
	account, errs := c.GetAccount()
	if errs != nil {
		return errs[0]
	}

	if err := d.Set("id", account.ID); err != nil {
		return err
	}

	if err := d.Set("slug", account.Slug); err != nil {
		return err
	}

	if err := d.Set("name", account.Name); err != nil {
		return err
	}

	d.SetId(account.ID)

	return nil
}
