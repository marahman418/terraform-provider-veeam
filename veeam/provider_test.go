package veeam

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"veeam": testAccProvider,
	}
}
func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("VEEAM_SERVER_IP"); v == "" {
		t.Fatal("VEEAM_SERVER_IP must be set for acceptance tests")
	}

	if v := os.Getenv("VEEAM_SERVER_PORT"); v == "" {
		t.Fatal("VEEAM_SERVER_PORT must be set for acceptance tests")
	}

	if v := os.Getenv("VEEAM_SERVER_USERNAME"); v == "" {
		t.Fatal("VEEAM_SERVER_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("VEEAM_SERVER_PASSWORD"); v == "" {
		t.Fatal("VEEAM_SERVER_PASSWORD must be set for acceptance tests")
	}
}