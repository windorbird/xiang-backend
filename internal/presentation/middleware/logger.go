package middleware

import (
    "bytes"
    "io/ioutil"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/windorbird/xiang-backend/log"
    "go.uber.org/zap"
)

// 是否打印请求体，生产环境建议关闭
const PrintRequestBody = true
const PrintResponseBody = true

type bodyWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
    w.body.Write(b) // 记录响应
    return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        traceID := uuid.NewString()
        c.Set("trace_id", traceID)

        // 读取请求 body
        var reqBody []byte
        if PrintRequestBody {
            if c.Request.Body != nil {
                reqBody, _ = ioutil.ReadAll(c.Request.Body)
                c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // 还原 body，否则后续读不到
            }
        }

        // 包装响应 writer
        w := &bodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
        c.Writer = w

        c.Next()

        latency := time.Since(start)

        fields := []zap.Field{
            zap.String("trace_id", traceID),
            zap.String("method", c.Request.Method),
            zap.String("path", c.Request.URL.Path),
            zap.String("ip", c.ClientIP()),
            zap.Int("status", c.Writer.Status()),
            zap.Int64("cost_ms", latency.Milliseconds()),
        }

        if PrintRequestBody {
            fields = append(fields, zap.String("req_body", string(reqBody)))
        }
        if PrintResponseBody {
            fields = append(fields, zap.String("resp_body", w.body.String()))
        }

        if len(c.Errors) > 0 {
            fields = append(fields, zap.String("error", c.Errors[0].Error()))
        }

        log.Logger.Info("http access", fields...)
    }
}
