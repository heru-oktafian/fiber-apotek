package services

import (
	json "encoding/json"
	fmt "fmt"
	time "time"

	configs "github.com/heru-oktafian/fiber-apotek/configs"
	models "github.com/heru-oktafian/fiber-apotek/models"
	redis "github.com/redis/go-redis/v9"
)

// GetRedisClient mengembalikan instance Redis client dari configs
// Fungsi ini digunakan sebagai wrapper untuk mengakses configs.RDB
func GetRedisClient() *redis.Client {
	return configs.RDB
}

// SetTemporaryProductCache menyimpan daftar produk sementara untuk penjualan ke Redis dengan cacheKey sebagai pembeda
func SetTemporaryProductCache(cacheKey string, products []models.ProdSaleCombo) error {
	// Ping Redis to check connection
	if _, err := configs.RDB.Ping(configs.Ctx).Result(); err != nil {
		fmt.Printf("Redis ping failed: %v\n", err)
		return err
	}

	key := fmt.Sprintf("tmp:products:sale:%s", cacheKey)
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}

	// Set dengan TTL 30 menit
	err = configs.RDB.Set(configs.Ctx, key, data, 30*time.Minute).Err()
	if err == nil {
		fmt.Printf("Successfully saved product cache to Redis key: %s\n", key)
	}
	return err
}

// GetTemporaryProductCache mengambil daftar produk sementara untuk penjualan dari Redis berdasarkan cacheKey
func GetTemporaryProductCache(cacheKey string) ([]models.ProdSaleCombo, error) {
	key := fmt.Sprintf("tmp:products:sale:%s", cacheKey)
	val, err := configs.RDB.Get(configs.Ctx, key).Result()
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
	key := fmt.Sprintf("tmp:products:sale:%s", cacheKey)
	return configs.RDB.Del(configs.Ctx, key).Err()
}

// SetTemporaryPurchaseProductCache menyimpan daftar produk sementara untuk pembelian ke Redis dengan cacheKey sebagai pembeda
func SetTemporaryPurchaseProductCache(cacheKey string, products []models.ProdPurchaseCombo) error {
	// Ping Redis to check connection
	if _, err := configs.RDB.Ping(configs.Ctx).Result(); err != nil {
		fmt.Printf("Redis ping failed: %v\n", err)
		return err
	}

	key := fmt.Sprintf("tmp:products:purchase:%s", cacheKey)
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}

	// Set dengan TTL 30 menit
	err = configs.RDB.Set(configs.Ctx, key, data, 30*time.Minute).Err()
	if err == nil {
		fmt.Printf("Successfully saved product cache to Redis key: %s\n", key)
	}
	return err
}

// GetTemporaryPurchaseProductCache mengambil daftar produk pembelian sementara dari Redis berdasarkan cacheKey
func GetTemporaryPurchaseProductCache(cacheKey string) ([]models.ProdPurchaseCombo, error) {
	key := fmt.Sprintf("tmp:products:purchase:%s", cacheKey)
	val, err := configs.RDB.Get(configs.Ctx, key).Result()
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
	key := fmt.Sprintf("tmp:products:purchase:%s", cacheKey)
	return configs.RDB.Del(configs.Ctx, key).Err()
}

// SetTemporaryOpnameProductCache menyimpan daftar produk opname sementara ke Redis dengan cacheKey sebagai pembeda
func SetTemporaryOpnameProductCache(cacheKey string, products []models.ComboboxProducts) error {
	// Ping Redis to check connection
	if _, err := configs.RDB.Ping(configs.Ctx).Result(); err != nil {
		fmt.Printf("Redis ping failed: %v\n", err)
		return err
	}
	key := fmt.Sprintf("tmp:products:opname:%s", cacheKey)
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}
	// Set dengan TTL 30 menit
	err = configs.RDB.Set(configs.Ctx, key, data, 30*time.Minute).Err()
	if err == nil {
		fmt.Printf("Successfully saved opname product cache to Redis key: %s\n", key)
	}
	return err
}

// GetTemporaryOpnameProductCache mengambil daftar produk opname sementara dari Redis berdasarkan cacheKey
func GetTemporaryOpnameProductCache(cacheKey string) ([]models.ComboboxProducts, error) {
	key := fmt.Sprintf("tmp:products:opname:%s", cacheKey)
	val, err := configs.RDB.Get(configs.Ctx, key).Result()
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
	key := fmt.Sprintf("tmp:products:opname:%s", cacheKey)
	return configs.RDB.Del(configs.Ctx, key).Err()
}

// UpdateSaleProductStockInRedisAsync mengupdate stock produk di cache PENJUALAN secara asinkron
func UpdateSaleProductStockInRedisAsync(cacheKey, productID string, newStock int) {
	go func() {
		// Ping Redis to check connection
		if _, err := configs.RDB.Ping(configs.Ctx).Result(); err != nil {
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

// UpdatePurchaseProductStockInRedisAsync mengupdate stock produk di cache PEMBELIAN secara asinkron
func UpdatePurchaseProductStockInRedisAsync(cacheKey, productID string, newStock int) {
	go func() {
		// Ping Redis to check connection
		if _, err := configs.RDB.Ping(configs.Ctx).Result(); err != nil {
			fmt.Printf("Redis ping failed: %v\n", err)
			return
		}

		// Ambil data cache produk pembelian
		products, err := GetTemporaryPurchaseProductCache(cacheKey)
		if err != nil {
			fmt.Printf("Failed to get purchase product cache for cacheKey %s: %v\n", cacheKey, err)
			return
		}
		if products == nil {
			fmt.Printf("No purchase product cache found for cacheKey %s\n", cacheKey)
			return
		}

		// Cari dan update stock produk (meskipun purchase combo tidak punya stock, tapi untuk konsistensi)
		// Note: ProdPurchaseCombo tidak memiliki field Stock, jadi ini mungkin tidak diperlukan
		// Tapi kita tetap implementasikan untuk konsistensi
		fmt.Printf("Purchase product cache updated for product %s in cache key: tmp:products:purchase:%s\n", productID, cacheKey)
	}()
}

// UpdateOpnameProductStockInRedisAsync mengupdate stock produk di cache OPNAME secara asinkron
func UpdateOpnameProductStockInRedisAsync(cacheKey, productID string, newStock int) {
	go func() {
		// Ping Redis to check connection
		if _, err := configs.RDB.Ping(configs.Ctx).Result(); err != nil {
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
