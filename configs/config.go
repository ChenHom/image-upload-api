package configs

type Config struct {
	Port         string `json:"port"`
	StoragePath  string `json:"storage_path"`
	CloudEnabled bool   `json:"cloud_enabled"`
	AIEndpoint   string `json:"ai_endpoint"`
	APIKey       string `json:"api_key"`
}

func LoadConfig(filePath string) (*Config, error) {
	// 這裡可以實作讀取和解析配置檔案的邏輯
	return &Config{}, nil
}
