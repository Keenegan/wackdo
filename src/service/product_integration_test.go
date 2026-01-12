package service

import (
	"testing"
	"wackdo/src/initializers"
	"wackdo/src/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) {
	var err error
	initializers.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err, "failed to connect to test database")

	err = initializers.DB.AutoMigrate(&models.Product{}, &models.Menu{})
	require.NoError(t, err, "failed to migrate test database")
}

func TestProductCRUD(t *testing.T) {
	setupTestDB(t)

	// CREATE
	product := models.Product{
		Name:        "Test Burger",
		BasePrice:   9.99,
		Description: "A delicious test burger",
		Image:       "burger.png",
		Category:    models.CategoryFood,
		Available:   true,
	}

	createdProduct, err := CreateProduct(product)
	require.NoError(t, err, "CreateProduct should not return an error")
	assert.NotZero(t, createdProduct.ID, "created product should have an ID")
	assert.Equal(t, "Test Burger", createdProduct.Name)
	assert.Equal(t, float32(9.99), createdProduct.BasePrice)
	assert.Equal(t, models.CategoryFood, createdProduct.Category)

	// GET by ID
	fetchedProduct, err := GetProductById(int(createdProduct.ID))
	require.NoError(t, err, "GetProductById should not return an error")
	assert.Equal(t, createdProduct.ID, fetchedProduct.ID)
	assert.Equal(t, "Test Burger", fetchedProduct.Name)

	// GET by name
	productsByName, err := GetProductByName("Burger")
	require.NoError(t, err, "GetProductByName should not return an error")
	assert.Len(t, productsByName, 1)
	assert.Equal(t, "Test Burger", productsByName[0].Name)

	// GET all (paginated)
	allProducts, err := GetProducts(0, 10)
	require.NoError(t, err, "GetProducts should not return an error")
	assert.Len(t, allProducts, 1)

	// ProductExists
	exists, err := ProductExists("Test Burger")
	require.NoError(t, err, "ProductExists should not return an error")
	assert.True(t, exists)

	notExists, err := ProductExists("Non-existent")
	require.NoError(t, err, "ProductExists should not return an error")
	assert.False(t, notExists)

	// UPDATE
	fetchedProduct.Name = "Updated Burger"
	fetchedProduct.BasePrice = 12.99
	updatedProduct, err := UpdateProduct(fetchedProduct)
	require.NoError(t, err, "UpdateProduct should not return an error")
	assert.Equal(t, "Updated Burger", updatedProduct.Name)
	assert.Equal(t, float32(12.99), updatedProduct.BasePrice)

	// Verify update persisted
	refetchedProduct, err := GetProductById(int(createdProduct.ID))
	require.NoError(t, err)
	assert.Equal(t, "Updated Burger", refetchedProduct.Name)
	assert.Equal(t, float32(12.99), refetchedProduct.BasePrice)

	// DELETE
	err = DeleteProductById(int(createdProduct.ID))
	require.NoError(t, err, "DeleteProductById should not return an error")

	// Verify deletion
	_, err = GetProductById(int(createdProduct.ID))
	assert.Error(t, err, "GetProductById should return an error after deletion")

	allProductsAfterDelete, err := GetProducts(0, 10)
	require.NoError(t, err)
	assert.Len(t, allProductsAfterDelete, 0, "no products should remain after deletion")
}
