package responses

// Fungsi untuk memberikan respon ketika controller gagal dijalankan
func StatusFailed(message string) map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "failed",
		"message": message,
	}
	return result
}

// Fungsi untuk memberikan respon ketika controller service error dijalankan
func StatusFailedInternal(message string, data interface{}) map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "Unathorized failed",
		"message": message,
		"data":    data,
	}
	return result
}

// Fungsi untuk memberikan respon ketika controller berhasil dijalankan
func StatusSuccess(message string) map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "success",
		"message": message,
	}
	return result
}

// Fungsi untuk memberikan respon ketika controller berhasil dijalankan dan menerima masukan data
func StatusSuccessRegister(data interface{}) map[string]interface{} {
	var result = map[string]interface{}{
		"status": "success",
		"data":   data,
	}
	return result
}
