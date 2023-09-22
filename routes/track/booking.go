package track

import (
	"net/http"
	"track/config"
	"github.com/gin-gonic/gin"
)

type BookingRequest struct {
	UserName     string `json:"userName"`
	FromLocation int    `json:"fromLocation"`
	ToLocation   int    `json:"toLocation"`
}

func BookNow(c *gin.Context) {
	var request BookingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.Database

	// Insert data into the database
	_, err := db.Exec("INSERT INTO bookings (user_name, from_location, to_location) VALUES (?, ?, ?)",
		request.UserName, request.FromLocation, request.ToLocation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking saved successfully"})
}
