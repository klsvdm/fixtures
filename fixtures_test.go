package fixtures

import (
	"os"
	"slices"
	"testing"

	"gopkg.in/yaml.v3"
)

type __user struct {
	Name  string `yaml:"name"`
	Age   int    `yaml:"age"`
	Email string `yaml:"email"`
}

type __product struct {
	Name  string `yaml:"name"`
	Price int    `yaml:"price"`
}

func Test_Fixtures(t *testing.T) {
	fixture, err := Load("./fixtures")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("list fixture", func(t *testing.T) {
		usersData, err := os.ReadFile("./fixtures/users.yaml")
		if err != nil {
			t.Fatal(err)
		}

		expectedUsers := make([]__user, 0)
		_ = yaml.Unmarshal(usersData, &expectedUsers)

		users := GetList[__user](t, fixture, "users")

		if !slices.Equal(expectedUsers, users) {
			t.Errorf("expected %v, got %v", expectedUsers, users)
		}
	})

	t.Run("map fixture", func(t *testing.T) {
		productsData, err := os.ReadFile("./fixtures/products.yaml")
		if err != nil {
			t.Fatal(err)
		}

		expectedProducts := make(map[string][]__product)
		_ = yaml.Unmarshal(productsData, &expectedProducts)

		products := GetMap[[]__product](t, fixture, "products")

		if len(expectedProducts) != len(products) {
			t.Fatalf("expected %v, got %v", expectedProducts, products)
		}

		for key, value := range expectedProducts {
			actualValue, ok := products[key]
			if !ok {
				t.Errorf("expected key '%v' not found", key)
				continue
			}

			if !slices.Equal(value, actualValue) {
				t.Errorf("expected %v, got %v", value, actualValue)
			}
		}
	})
}