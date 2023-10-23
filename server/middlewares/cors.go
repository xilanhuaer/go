package middlewares

// func Cors() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		context.Writer.Header().Set("Access-Control-Allow-Origin", "http://111.230.89.67:80")
// 		context.Writer.Header().Set("Access-Control-Max-Age", "86400")
// 		context.Writer.Header().Set("Access-Control-Allow-Methods", "*")
// 		context.Writer.Header().Set("Access-Control-Allow-Headers", "*")
// 		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		if context.Request.Method == "OPTIONS" {
// 			context.AbortWithStatus(http.StatusOK)
// 		}
// 		context.Next()
// 	}
// }
