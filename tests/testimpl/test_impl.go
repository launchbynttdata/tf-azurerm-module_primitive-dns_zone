package common

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/dns/mgmt/dns"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/lib/azure/configure"
	internalDns "github.com/launchbynttdata/lcaf-component-terratest/lib/azure/dns"
	"github.com/launchbynttdata/lcaf-component-terratest/lib/azure/login"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

const terraformDir string = "../../examples/public_dns_zone"

func TestDnsZone(t *testing.T, ctx types.TestContext) {

	envVarMap := login.GetEnvironmentVariables()
	clientID := envVarMap["clientID"]
	clientSecret := envVarMap["clientSecret"]
	tenantID := envVarMap["tenantID"]
	subscriptionID := envVarMap["subscriptionID"]

	spt, err := login.GetServicePrincipalToken(clientID, clientSecret, tenantID)
	if err != nil {
		t.Fatalf("Error getting Service Principal Token: %v", err)
	}

	dnsZonesClient := internalDns.GetDNSZonesClient(spt, subscriptionID)

	var varFiles = []string{"../../examples/public_dns_zone/test.tfvars"}

	terraformOptions := configure.ConfigureTerraform(terraformDir, varFiles)

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	dnsZoneIds := terraform.OutputList(t, ctx.TerratestTerraformOptions(), "ids")
	for range dnsZoneIds {
		t.Run("doesDnsZoneExist", func(t *testing.T) {
			checkDNSZoneExistence(t, dnsZonesClient, terraformOptions, ctx)
		})
	}
}

func checkDNSZoneExistence(t *testing.T, dnsZonesClient dns.ZonesClient, terraformOptions *terraform.Options, ctx types.TestContext) {
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	expectedDomainNames := make(map[string]bool)
	inputDomainNames := ctx.TestConfig().(*ThisTFModuleConfig).DomainNames

	// Set all expected domain names to false
	for _, domain := range inputDomainNames {
		expectedDomainNames[domain] = false
	}

	// If expected domain name matches with actual domain name, set the map value to true
	for domain := range expectedDomainNames {
		dnsZone, err := dnsZonesClient.Get(context.Background(), resourceGroupName, domain)
		if err != nil {
			t.Fatalf("Error getting DNS Zone: %v", err)
		}
		if dnsZone.Name == nil {
			t.Fatalf("DNS Zone does not exist")
		}
		if *dnsZone.Name == domain {
			expectedDomainNames[domain] = true
		}
	}

	// Check if all domain names are true. If not, test fails
	for domain, found := range expectedDomainNames {
		assert.True(t, found, "DNS Zone does not exist for domain: "+domain)
	}

}
