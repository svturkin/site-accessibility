package helpers

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	requestsCountMap   = make(map[string]uint)
	requestsCountMutex sync.Mutex
)

func CountRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		routeName := c.FullPath()

		requestsCountMutex.Lock()
		defer requestsCountMutex.Unlock()

		if routeName != "" {
			requestsCountMap[routeName]++
		}

		c.Next()
	}
}

func GetRequestsCount() map[string]uint {
	return requestsCountMap
}
