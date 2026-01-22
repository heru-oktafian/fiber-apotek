package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// validate adalah instance global dari validator.
// Diinisialisasi sekali saja untuk menghindari membuat instance baru di setiap request,
// hal ini menghemat memori dan CPU.
var validate = validator.New()

// ValidateStruct membantu memvalidasi struct berdasarkan tag 'validate' yang didefinisikan.
// Jika validasi gagal, fungsi ini mengembalikan error dengan pesan yang lebih mudah dibaca
// yang menjelaskan field mana yang gagal dan aturan validasi apa yang dilanggar.
//
// Parameter:
//
//	s interface{} : Struct yang ingin divalidasi.
//
// Return:
//
//	error : nil jika validasi berhasil, atau error jika validasi gagal.
func ValidateStruct(s interface{}) error {
	// Jalankan validasi pada struct
	if err := validate.Struct(s); err != nil {
		// Jika ada error validasi, kami akan memproses mereka untuk membuat pesan yang lebih informatif.
		var errorMessages []string
		// Konversikan error ke tipe validator.ValidationErrors untuk iterasi.
		for _, err := range err.(validator.ValidationErrors) {
			// Buat pesan error dasar. Contoh: "Field 'MemberId' is required"
			errorMessage := fmt.Sprintf("Field '%s' is %s", err.Field(), err.Tag())

			// Tambahkan parameter validasi jika tersedia (misal: `min`, `max`, `len`, dll)
			if err.Param() != "" {
				errorMessage = fmt.Sprintf("Field '%s' %s %s", err.Field(), err.Tag(), err.Param())
			}

			errorMessages = append(errorMessages, errorMessage)
		}
		// Gabungkan semua pesan error menjadi satu string yang dipisahkan dengan koma.
		return fmt.Errorf("validation failed: %s", strings.Join(errorMessages, ", "))
	}
	// Jika tidak ada error, validasi berhasil dilakukan.
	return nil
}

/*
// Anda bisa menambahkan fungsi validasi kustom di sini jika diperlukan.
// Contoh cara mendaftarkan validasi kustom:
import "reflect" // Diperlukan untuk validasi kustom saat mengecek tipe field

func init() {
	// Contoh: mendaftarkan aturan validasi 'isAdult'
	// yang mengecek apakah nilai integer >= 18.
	validate.RegisterValidation("isAdult", func(fl validator.FieldLevel) bool {
		// Pastikan field adalah tipe integer sebelum membandingkan.
		if fl.Field().Kind() == reflect.Int {
			return fl.Field().Int() >= 18
		}
		return false // Atau tangani tipe data lain sesuai kebutuhan
	})

	// Untuk menggunakan validasi kustom ini, tambahkan tag `validate:"isAdult"` pada field struct Anda.
	// Contoh: Age int `json:"age" validate:"required,isAdult"`
}
*/
