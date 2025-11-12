package douyin

import (
    "fmt"
    "testing"

    "github.com/joho/godotenv"
    "github.com/windorbird/xiang-backend/internal/infrastructure/persistence/redis"
    "github.com/windorbird/xiang-backend/log"
)

func init() {
    err := godotenv.Load("../../../../.env")
    if err != nil {
        log.Logger.Fatal("Error loading .env file")
    }
    log.InitLogger()
    redis.InitRedis()
}
func TestTextCheck(t *testing.T) {
    token := "08011218474862536558416e7a7a542f41376b644e6b536e6a773d3d1883fd9d01"
    content := "我爱你"
    res, err := DetectContent(token, content)
    if err != nil {
        t.Errorf("DetectContent() error = %v", err)
        return
    }

    fmt.Printf("DetectContent(%v) res = %v\n",  content, res)
}