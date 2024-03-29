package common

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestDnsZone(t *testing.T, ctx types.TestContext) {

	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		t.Fatal("ARM_SUBSCRIPTION_ID is not set in the environment variables ")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("Unable to get credentials: %e\n", err)
	}

	clientFactory, err := armdns.NewClientFactory(subscriptionID, cred, nil)
	if err != nil {
		t.Fatalf("Unable to get clientFactory: %e\n", err)
	}

	dnsZonesClient := clientFactory.NewZonesClient()

	dnsZoneIds := terraform.OutputList(t, ctx.TerratestTerraformOptions(), "ids")
	for range dnsZoneIds {
		t.Run("doesDnsZoneExist", func(t *testing.T) {
			checkDNSZoneExistence(t, dnsZonesClient, ctx)
		})
	}
}

func checkDNSZoneExistence(t *testing.T, dnsZonesClient *armdns.ZonesClient, ctx types.TestContext) {
	resourceGroupName := terraform.Output(t, ctx.TerratestTerraformOptions(), "resource_group_name")
	expectedDomainNames := make(map[string]bool)
	inputDomainNames := ctx.TestConfig().(*ThisTFModuleConfig).DomainNames

	// Set all expected domain names to false
	for _, domain := range inputDomainNames {
		expectedDomainNames[domain] = false
	}

	// If expected domain name matches with actual domain name, set the map value to true
	for domain := range expectedDomainNames {
		dnsZone, err := dnsZonesClient.Get(context.Background(), resourceGroupName, domain, nil)
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
