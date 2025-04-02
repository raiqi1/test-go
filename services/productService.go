package services

import (
	"encoding/json"
	"product-management/config"
	"product-management/models"
	"strconv"
	"time"

	"golang.org/x/net/context"
)

var ctx = context.Background()

func GetAllProducts() ([]models.Product, string, error) {
	var products []models.Product
	source := "database"

	val, err := config.RedisClient.Get(ctx, "products").Result()
	if err == nil {
		err = json.Unmarshal([]byte(val), &products)
		if err == nil {
			source = "redis"
			return products, source, nil
		}
	}

	result := config.DB.Find(&products)
	if result.Error != nil {
		return nil, source, result.Error
	}

	data, err := json.Marshal(products)
	if err == nil {
		config.RedisClient.Set(ctx, "products", data, 30*time.Minute)
	}

	return products, source, nil
}

func GetProductByID(id uint) (models.Product, string, error) {
	var product models.Product
	source := "database"

	key := "product:" + strconv.FormatUint(uint64(id), 10)
	val, err := config.RedisClient.Get(ctx, key).Result()

	if err == nil {
		// Data found in Redis
		err = json.Unmarshal([]byte(val), &product)
		if err == nil {
			source = "redis"
			return product, source, nil
		}
	}

	result := config.DB.First(&product, id)
	if result.Error != nil {
		return product, source, result.Error
	}

	data, err := json.Marshal(product)
	if err == nil {
		config.RedisClient.Set(ctx, key, data, 30*time.Minute)
	}

	return product, source, nil
}

func CreateProduct(product models.Product) (models.Product, error) {
	result := config.DB.Create(&product)

	// Invalidate cache
	config.RedisClient.Del(ctx, "products")

	return product, result.Error
}

func UpdateProduct(id uint, updatedProduct models.Product) (models.Product, error) {
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return product, err
	}

	product.Name = updatedProduct.Name
	product.Description = updatedProduct.Description
	product.Price = updatedProduct.Price
	product.Stock = updatedProduct.Stock

	result := config.DB.Save(&product)

	key := "product:" + strconv.FormatUint(uint64(id), 10)
	config.RedisClient.Del(ctx, key)
	config.RedisClient.Del(ctx, "products")

	return product, result.Error
}

func DeleteProduct(id uint) error {
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return err
	}

	key := "product:" + strconv.FormatUint(uint64(id), 10)
	config.RedisClient.Del(ctx, key)
	config.RedisClient.Del(ctx, "products")

	return config.DB.Delete(&product).Error
}
