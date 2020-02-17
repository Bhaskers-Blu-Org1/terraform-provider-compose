package compose

import (
	"errors"
	"fmt"
	"github.com/compose/gocomposeapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccComposeDataSourceAccount(t *testing.T) {

	var account composeapi.Account

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComposeDataSourceAccount(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountDataSourceExists("data.compose_account.account", &account),
				),
			},
		},
	})
}

func testAccCheckAccountDataSourceExists(n string, account *composeapi.Account) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := testAccProvider.Meta().(*composeapi.Client)
		account, err := conn.GetAccount()
		if err != nil {
			return err[0]
		}

		if rs.Primary.ID != account.ID {
			return errors.New("Account ID in terraform state did not match actual account ID")
		}
		return nil

	}
}

func testAccComposeDataSourceAccount() string {
	return "data \"compose_account\" \"account\" {}"
}
