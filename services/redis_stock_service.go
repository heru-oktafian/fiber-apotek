package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/heru-oktafian/fiber-apotek/models"
	"github.com/redis/go-redis/v9"
)

// Inisialisasi RedisClient ini adalah instance dari redis client yang akan digunnakan sebagai media pemrosesan data di fungsi-fungsi yang membutuhkan redis
var RedisClient *redis.Client = func() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	var host, port string
	if addr != "" {
		// Parse REDIS_ADDR as host:port
		parts := strings.Split(addr, ":")
		if len(parts) == 2 {
			host = parts[0]
			port = parts[1]
		} else {
			host = "localhost"
			port = "6379"
		}
	} else {
		host = os.Getenv("REDIS_HOST")
		if host == "" {
			host = "localhost"
		}
		port = os.Getenv("REDIS_PORT")
		if port == "" {
			port = "6379"
		}
	}
	password := os.Getenv("REDIS_PASS")
	dbStr := os.Getenv("REDIS_DB")
	db := 0
	if dbStr != "" {
		if parsed, err := strconv.Atoi(dbStr); err == nil {
			db = parsed
		}
	}
	return redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})
}()

// SetTemporaryProductCache menyimpan daftar produk sementara untuk penjualan ke Redis dengan cacheKey sebagai pembeda
func SetTemporaryProductCache(cacheKey string, products []models.ProdSaleCombo) error {
	ctx := context.Background()

	// Ping Redis to check connection
	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		fmt.Printf("Redis ping failed: %v\n", err)
		return err
	}

	key := fmt.Sprintf("tmp:products:sale:%s", cacheKey)
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}

	// Set dengan TTL 30 menit
	err = RedisClient.Set(ctx, key, data, 30*time.Minute).Err()
	if err == nil {
		fmt.Printf("Successfully saved product cache to Redis key: %s\n", key)
	}
	return err
}

// GetTemporaryProductCache mengambil daftar produk sementara untuk penjualan dari Redis berdasarkan cacheKey
func GetTemporaryProductCache(cacheKey string) ([]models.ProdSaleCombo, error) {
	ctx := context.Background()
	key := fmt.Sprintf("tmp:products:sale:%s", cacheKey)
	val, err := RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // Tidak ada data cache
	}
	if err != nil {
		return nil, err
	}
	var products []models.ProdSaleCombo
	if err := json.Unmarshal([]byte(val), &products); err != nil {
		return nil, err
	}
	return products, nil
}

// DeleteTemporaryProductCache menghapus cache produk sementara untuk penjualan dari Redis berdasarkan cacheKey
func DeleteTemporaryProductCache(cacheKey string) error {
	ctx := context.Background()
	key := fmt.Sprintf("tmp:products:sale:%s", cacheKey)
	return RedisClient.Del(ctx, key).Err()
}

// GetTemporaryPurchaseProductCache mengambil daftar produk pembelian sementara dari Redis berdasarkan cacheKey
func GetTemporaryPurchaseProductCache(cacheKey string) ([]models.ProdPurchaseCombo, error) {
	ctx := context.Background()
	key := fmt.Sprintf("tmp:products:purchase:%s", cacheKey)
	val, err := RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // Tidak ada data cache
	}
	if err != nil {
		return nil, err
	}
	var products []models.ProdPurchaseCombo
	if err := json.Unmarshal([]byte(val), &products); err != nil {
		return nil, err
	}
	return products, nil
}

// DeleteTemporaryPurchaseProductCache menghapus cache produk pembelian sementara dari Redis berdasarkan cacheKey
func DeleteTemporaryPurchaseProductCache(cacheKey string) error {
	ctx := context.Background()
	key := fmt.Sprintf("tmp:products:purchase:%s", cacheKey)
	return RedisClient.Del(ctx, key).Err()
}

// SetTemporaryOpnameProductCache menyimpan daftar produk opname sementara ke Redis dengan cacheKey sebagai pembeda
func SetTemporaryOpnameProductCache(cacheKey string, products []models.ComboboxProducts) error {
	ctx := context.Background()
	// Ping Redis to check connection
	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		fmt.Printf("Redis ping failed: %v\n", err)
		return err
	}
	key := fmt.Sprintf("tmp:products:opname:%s", cacheKey)
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}
	// Set dengan TTL 30 menit
	err = RedisClient.Set(ctx, key, data, 30*time.Minute).Err()
	if err == nil {
		fmt.Printf("Successfully saved opname product cache to Redis key: %s\n", key)
	}
	return err
}

// GetTemporaryOpnameProductCache mengambil daftar produk opname sementara dari Redis berdasarkan cacheKey
func GetTemporaryOpnameProductCache(cacheKey string) ([]models.ComboboxProducts, error) {
	ctx := context.Background()
	key := fmt.Sprintf("tmp:products:opname:%s", cacheKey)
	val, err := RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // Tidak ada data cache
	}
	if err != nil {
		return nil, err
	}
	var products []models.ComboboxProducts
	if err := json.Unmarshal([]byte(val), &products); err != nil {
		return nil, err
	}
	return products, nil
}

// DeleteTemporaryOpnameProductCache menghapus cache produk opname sementara dari Redis berdasarkan cacheKey
func DeleteTemporaryOpnameProductCache(cacheKey string) error {
	ctx := context.Background()
	key := fmt.Sprintf("tmp:products:opname:%s", cacheKey)
	return RedisClient.Del(ctx, key).Err()
}

// UpdateSaleProductStockInRedisAsync mengupdate stock produk di cache PENJUALAN secara asinkron
func UpdateSaleProductStockInRedisAsync(cacheKey, productID string, newStock int) {
	go func() {
		ctx := context.Background()
		// Ping Redis to check connection
		if _, err := RedisClient.Ping(ctx).Result(); err != nil {
			fmt.Printf("Redis ping failed: %v\n", err)
			return
		}

		// Ambil data cache produk penjualan
		products, err := GetTemporaryProductCache(cacheKey)
		if err != nil {
			fmt.Printf("Failed to get product sale cache for cacheKey %s: %v\n", cacheKey, err)
			return
		}
		if products == nil {
			// Cache tidak ditemukan (mungkin belum diset), normal.
			return
		}

		// Cari dan update stock produk
		found := false
		for i := range products {
			if products[i].ProductId == productID {
				products[i].Stock = newStock
				found = true
				break
			}
		}

		if found {
			// Simpan kembali ke Redis hanya jika produk ditemukan di cache
			if err := SetTemporaryProductCache(cacheKey, products); err != nil {
				fmt.Printf("Failed to update product sale cache for cacheKey %s: %v\n", cacheKey, err)
			} else {
				fmt.Printf("Successfully updated stock for product %s in Sale cache key: tmp:products:sale:%s\n", productID, cacheKey)
			}
		}
	}()
}

// UpdateOpnameProductStockInRedisAsync mengupdate stock produk di cache OPNAME secara asinkron
func UpdateOpnameProductStockInRedisAsync(cacheKey, productID string, newStock int) {
	go func() {
		ctx := context.Background()
		// Ping Redis to check connection
		if _, err := RedisClient.Ping(ctx).Result(); err != nil {
			fmt.Printf("Redis ping failed: %v\n", err)
			return
		}

		// Ambil data cache produk opname
		products, err := GetTemporaryOpnameProductCache(cacheKey)
		if err != nil {
			fmt.Printf("Failed to get product opname cache for cacheKey %s: %v\n", cacheKey, err)
			return
		}
		if products == nil {
			// Cache tidak ditemukan, normal.
			return
		}

		// Cari dan update stock produk
		found := false
		for i := range products {
			if products[i].ProID == productID { // struct Opname pakai ProID
				products[i].Stock = newStock
				found = true
				break
			}
		}

		if found {
			// Simpan kembali ke Redis
			if err := SetTemporaryOpnameProductCache(cacheKey, products); err != nil {
				fmt.Printf("Failed to update product opname cache for cacheKey %s: %v\n", cacheKey, err)
			} else {
				fmt.Printf("Successfully updated stock for product %s in Opname cache key: tmp:products:opname:%s\n", productID, cacheKey)
			}
		}
	}()
}
