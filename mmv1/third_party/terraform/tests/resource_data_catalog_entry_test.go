package google_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataCatalogEntry_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": google.RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { TestAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckDataCatalogEntryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogEntry_dataCatalogEntryBasicExample(context),
			},
			{
				ResourceName:      "google_data_catalog_entry.basic_entry",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataCatalogEntry_dataCatalogEntryFullExample(context),
			},
			{
				ResourceName:      "google_data_catalog_entry.basic_entry",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataCatalogEntry_dataCatalogEntryBasicExample(context),
			},
			{
				ResourceName:      "google_data_catalog_entry.basic_entry",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
