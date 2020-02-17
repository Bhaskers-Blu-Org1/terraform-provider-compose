package compose

import (
	"fmt"
	"testing"

	composeapi "github.com/compose/gocomposeapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccComposeResourceDeploymentPostgres(t *testing.T) {

	var deployment composeapi.Deployment
	resourceName := "compose_deployment.postgresql"

	deploymentName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { TestAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDeploymentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentConfigPostgres(deploymentName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentResourceExists(resourceName, &deployment),
				),
			},
		},
	})
}

func TestAccComposeResourceDeploymentRedis(t *testing.T) {

	var deployment composeapi.Deployment
	resourceName := "compose_deployment.redis"

	deploymentName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { TestAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDeploymentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentConfigRedis(deploymentName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentResourceExists(resourceName, &deployment),
				),
			},
		},
	})
}

// func TestAccComposeResourceDeploymentElasticSearch(t *testing.T) {

// var deployment composeapi.Deployment
// resourceName := "compose_deployment.elastic_search"

// deploymentName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

// resource.ParallelTest(t, resource.TestCase{
// PreCheck:     func() { TestAccPreCheck(t) },
// Providers:    testAccProviders,
// CheckDestroy: testAccCheckDeploymentDestroy,
// Steps: []resource.TestStep{
// {
// Config: testAccDeploymentConfigElasticSearch(deploymentName),
// Check: resource.ComposeTestCheckFunc(
// testAccCheckDeploymentResourceExists(resourceName, &deployment),
// ),
// },
// },
// })
// }

// func TestAccComposeResourceDeploymentEtcd(t *testing.T) {

// var deployment composeapi.Deployment
// resourceName := "compose_deployment.etcd"

// deploymentName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

// resource.ParallelTest(t, resource.TestCase{
// PreCheck:     func() { TestAccPreCheck(t) },
// Providers:    testAccProviders,
// CheckDestroy: testAccCheckDeploymentDestroy,
// Steps: []resource.TestStep{
// {
// Config: testAccDeploymentConfigEtcd(deploymentName),
// Check: resource.ComposeTestCheckFunc(
// testAccCheckDeploymentResourceExists(resourceName, &deployment),
// ),
// },
// },
// })
// }

// func TestAccComposeResourceDeploymentScylla(t *testing.T) {

// var deployment composeapi.Deployment
// resourceName := "compose_deployment.scylla"

// deploymentName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

// resource.ParallelTest(t, resource.TestCase{
// PreCheck:     func() { TestAccPreCheck(t) },
// Providers:    testAccProviders,
// CheckDestroy: testAccCheckDeploymentDestroy,
// Steps: []resource.TestStep{
// {
// Config: testAccDeploymentConfigScylla(deploymentName),
// Check: resource.ComposeTestCheckFunc(
// testAccCheckDeploymentResourceExists(resourceName, &deployment),
// ),
// },
// },
// })
// }

func TestAccComposeResourceDeploymentMysql(t *testing.T) {

	var deployment composeapi.Deployment
	resourceName := "compose_deployment.mysql"

	deploymentName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { TestAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDeploymentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentConfigMysql(deploymentName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentResourceExists(resourceName, &deployment),
				),
			},
		},
	})
}

func TestAccComposeResourceDeploymentDisque(t *testing.T) {

	var deployment composeapi.Deployment
	resourceName := "compose_deployment.disque"

	deploymentName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { TestAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDeploymentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentConfigDisque(deploymentName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentResourceExists(resourceName, &deployment),
				),
			},
		},
	})
}

func testAccCheckDeploymentResourceExists(n string, deployment *composeapi.Deployment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := testAccProvider.Meta().(*composeapi.Client)
		_, err := conn.GetDeployment(rs.Primary.ID)
		if err != nil {
			return err[0]
		}

		return nil

	}
}

func testAccDeploymentConfigPostgres(name string) string {
	return fmt.Sprintf(`
data "compose_account" "account" {}

resource "compose_deployment" "postgresql" {
  name = "terraform-test-postgresql-%[1]s"
  account_id = "${data.compose_account.account.id}"
  datacenter = "aws:us-east-1"
  type = "postgresql"
  units = 1
}
  `, name)
}

func testAccDeploymentConfigRedis(name string) string {
	return fmt.Sprintf(`
data "compose_account" "account" {}

resource "compose_deployment" "redis" {
  name = "terraform-test-redis-%[1]s"
  account_id = "${data.compose_account.account.id}"
  datacenter = "aws:us-east-1"
  type = "redis"
  cache_mode = true
  units = 1
}
  `, name)
}

// func testAccDeploymentConfigElasticSearch(name string) string {
// return fmt.Sprintf(`
// data "compose_account" "account" {}

// resource "compose_deployment" "elastic_search" {
// name = "terraform-test-elastic-search-%[1]s"
// account_id = "${data.compose_account.account.id}"
// datacenter = "aws:us-east-1"
// type = "elastic_search"
// units = 1
// }
// `, name)
// }

// func testAccDeploymentConfigEtcd(name string) string {
// return fmt.Sprintf(`
// data "compose_account" "account" {}

// resource "compose_deployment" "etcd" {
// name = "terraform-test-etcd-%[1]s"
// account_id = "${data.compose_account.account.id}"
// datacenter = "aws:us-east-1"
// type = "etcd"
// units = 1
// }
// `, name)
// }

// func testAccDeploymentConfigScylla(name string) string {
// return fmt.Sprintf(`
// data "compose_account" "account" {}

// resource "compose_deployment" "scylla" {
// name = "terraform-test-scylla-%[1]s"
// account_id = "${data.compose_account.account.id}"
// datacenter = "aws:us-east-1"
// type = "scylla"
// units = 1
// }
// `, name)
// }

func testAccDeploymentConfigMysql(name string) string {
	return fmt.Sprintf(`
data "compose_account" "account" {}

resource "compose_deployment" "mysql" {
  name = "terraform-test-mysql-%[1]s"
  account_id = "${data.compose_account.account.id}"
  datacenter = "aws:us-east-1"
  type = "mysql"
  units = 1
}
  `, name)
}

func testAccDeploymentConfigDisque(name string) string {
	return fmt.Sprintf(`
data "compose_account" "account" {}

resource "compose_deployment" "disque" {
  name = "terraform-test-disque-%[1]s"
  account_id = "${data.compose_account.account.id}"
  datacenter = "aws:us-east-1"
  type = "disque"
  units = 1
}
  `, name)
}

func testAccCheckDeploymentDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*composeapi.Client)

	// loop through the resources in state, verifying each deployment
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "compose_deployment" {
			continue
		}

		deployment, err := conn.GetDeployment(rs.Primary.ID)
		if err != nil && deployment.ID == rs.Primary.ID {
			return fmt.Errorf("deployment (%s) still exists", rs.Primary.ID)
		}
	}
	return nil
}
