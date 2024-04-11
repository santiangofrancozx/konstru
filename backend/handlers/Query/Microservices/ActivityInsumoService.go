package Microservices

//func ActivityInsumoHandler(DSN string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var requestData struct {
//			ID string `json:"id"`
//		}
//
//		// Bind JSON a la estructura requestData
//		if err := c.BindJSON(&requestData); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		handleActividadQuery(c, DSN, requestData.ID)
//		//handleInsumoQuery(c, DSN, requestData.ID)
//
//	}
//}
