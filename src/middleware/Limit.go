package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/vidorg/vid_backend/src/controller/exception"
	"github.com/vidorg/vid_backend/src/model/dto/common"
	"io/ioutil"
	"net/http"
)

func LimitMiddleware(maxSize int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)
		rawBuf, err := c.GetRawData()
		if err != nil {
			conn, bufRw, err := c.Writer.Hijack()
			if err == nil {
				_ = bufRw.Flush()
				_ = conn.Close()
			}
			c.JSON(http.StatusRequestEntityTooLarge, common.Result{}.Error(http.StatusRequestEntityTooLarge).SetMessage(exception.RequestSizeError.Error()))
			c.Abort()
			return
		}

		// [GIN-debug] error on parse multipart form array: multipart: NextPart: EOF
		buf := bytes.NewBuffer(rawBuf)
		c.Request.Body = ioutil.NopCloser(buf)
		c.Next()
	}
}