package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"strings"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required,min=3"`
	Category    string  `json:"category"`    // "bread", "cake", "pastry"
	Price       float64 `json:"price"`
	Stock       int     `json:"stock" binding:"gte=0"`
	Description string  `json:"description"`
}

type Order struct {
	ID          int       `json:"id"`
	CustomerName string   `json:"customer_name"`
	ProductID   int       `json:"product_id"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total_price"`
	Status      string    `json:"status"`      // "pending", "completed", "cancelled"
	OrderDate   time.Time `json:"order_date"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`            // "customer", "staff", "admin"
}

var products = []Product{
	{ID: 1, Name: "Sourdough Bread", Category: "bread", Price: 3.50, Stock: 20, Description: "A classic sourdough bread with a crispy crust."},
	{ID: 2, Name: "Chocolate Cake", Category: "cake", Price: 15.00, Stock: 10, Description: "Rich and moist chocolate cake with chocolate frosting."},
	{ID: 3, Name: "Croissant", Category: "pastry", Price: 2.50, Stock: 30, Description: "Flaky and buttery croissant perfect for breakfast."},
}

var orders = []Order{
	{ID: 1, CustomerName: "Alice", ProductID: 1, Quantity: 2, TotalPrice: 7.00, Status: "completed", OrderDate: time.Now().AddDate(0, 0, -1)},
	{ID: 2, CustomerName: "Bob", ProductID: 2, Quantity: 1, TotalPrice: 15.00, Status: "pending", OrderDate: time.Now()},
}

var users = []User{
	{ID: 1, Username: "caelus", Role: "admin"},
	{ID: 2, Username: "march", Role: "admin"},
	{ID: 3, Username: "staff789", Role: "staff"},
	{ID: 4, Username: "franky", Role: "customer"},
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func main() {
	router := gin.New()

	router.Use(LoggingMiddleware())
	router.Use(AuthMiddleware())
	router.Use(RateLimitMiddleware())

	router.GET("/", func(c *gin.Context) {
		SendSuccessResponse(c, "Welcome to Caelus and March Bakery!", nil)
	})

	router.GET("/products", func(c *gin.Context) {
		SendSuccessResponse(c, "Here are all our delicious products!", gin.H{
			"products": products,
			"total":    len(products),
		})
	})

	router.GET("/products/search", func(c *gin.Context) {
		search := c.Query("q")
		category := c.Query("category")

		var results []Product
		for _, product := range products {
			nameMatch := search == "" || strings.Contains(strings.ToLower(product.Name), strings.ToLower(search))
			descMatch := search == "" || strings.Contains(strings.ToLower(product.Description), strings.ToLower(search))
			categoryMatch := category == "" || product.Category == category
			
			if (nameMatch || descMatch) && categoryMatch {
				results = append(results, product)
			}
		}

		SendSuccessResponse(c, "Search results retrieved successfully!", gin.H{
			"products": results,
			"total":    len(results),
		})
	})

	router.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, product := range products {
			if fmt.Sprintf("%d", product.ID) == id {
				SendSuccessResponse(c, "Product details retrieved successfully!", product)
				return
			}
		}
		
		SendErrorResponse(c, "Product not found", http.StatusNotFound)
	})

	router.POST("/orders", CustomerOnly(), func(c *gin.Context) {
		var newOrder Order
		if err := c.ShouldBindJSON(&newOrder); err != nil {
			SendErrorResponse(c, "Invalid order data", http.StatusBadRequest)
			return
		}

		var product *Product
		for i := range products {
			if products[i].ID == newOrder.ProductID {
				product = &products[i]
				break
			}
		}

		if product == nil {
			SendErrorResponse(c, "Product not found", http.StatusNotFound)
			return
		}

		if product.Stock < newOrder.Quantity {
			SendErrorResponse(c, fmt.Sprintf("Insufficient stock. Available: %d, Requested: %d", product.Stock, newOrder.Quantity), http.StatusBadRequest)
			return
		}

		// Kurangi stok produk
		product.Stock -= newOrder.Quantity

		newOrder.ID = len(orders) + 1
		newOrder.OrderDate = time.Now()
		newOrder.TotalPrice = calculateTotalPrice(newOrder.ProductID, newOrder.Quantity)
		newOrder.Status = "pending"

		orders = append(orders, newOrder)
		SendSuccessResponse(c, "Order created successfully!", newOrder)
	})

	router.GET("/orders/my", CustomerOnly(), func(c *gin.Context) {
		customerName := c.Query("customer_name")
		var customerOrders []Order
		for _, order := range orders {
			if order.CustomerName == customerName {
				customerOrders = append(customerOrders, order)
			}
		}
		SendSuccessResponse(c, "Your orders retrieved successfully!", gin.H{
			"orders":  customerOrders,
			"total":   len(customerOrders),
		})
	})
	
	router.GET("/orders", StaffOnly(), func(c *gin.Context) {
		SendSuccessResponse(c, "All orders retrieved successfully!", gin.H{
			"orders":  orders,
			"total":   len(orders),
		})
	})

	router.PATCH("/orders/:id/status", StaffOnly(), func(c *gin.Context) {
		id := c.Param("id")
		newStatus := c.Query("status")
		
		for i, order := range orders {
			if fmt.Sprintf("%d", order.ID) == id {
				orders[i].Status = newStatus
				SendSuccessResponse(c, "Order status updated successfully!", orders[i])
				return
			}
		}
		SendErrorResponse(c, "Order not found", http.StatusNotFound)
	})

	router.GET("/products/stock", StaffOnly(), func(c *gin.Context) {
		var stock []Product
		for _, product := range products {
			if product.Stock < 10 {
				stock = append(stock, product)
			}
		}
		SendSuccessResponse(c, "Stock information retrieved successfully!", gin.H{
			"products": stock,
			"total":    len(stock),
		})
	})

	router.POST("/products", AdminOnly(), func(c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			SendErrorResponse(c, "Invalid product data", http.StatusBadRequest)
			return
		}

		if len(c.Query("name")) < 3 {
			SendErrorResponse(c, "Product name must be at least 3 characters long", http.StatusBadRequest)
			return
		}

		if c.Query("category") == "" || (c.Query("category") != "bread" && c.Query("category") != "cake" && c.Query("category") != "pastry") {
			SendErrorResponse(c, "Product category must be one of: bread, cake, pastry", http.StatusBadRequest)
			return
		}

		if c.Query("price") <= "0" {
			SendErrorResponse(c, "Product price must be greater than 0", http.StatusBadRequest)
			return
		}

		if c.Query("stock") < "0" {
			SendErrorResponse(c, "Product stock cannot be negative", http.StatusBadRequest)
			return
		}

		newProduct.ID = len(products) + 1
		products = append(products, newProduct)
		SendSuccessResponse(c, "Product created successfully!", newProduct)
	})

	router.PUT("/products/:id", AdminOnly(), func(c *gin.Context) {
		id := c.Param("id")
		var updatedProduct Product
		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			SendErrorResponse(c, "Invalid product data", http.StatusBadRequest)
			return
		}
		
		for i, product := range products {
			if fmt.Sprintf("%d", product.ID) == id {
				if updatedProduct.Stock < 0 {
					SendErrorResponse(c, "Product stock cannot be negative", http.StatusBadRequest)
					return
				}

				if updatedProduct.Status != "completed" && updatedProduct.Status != "cancelled" {
					SendErrorResponse(c, "Product status must be either 'completed' or 'cancelled'", http.StatusBadRequest)
					return
				}
				
				products[i].Name = updatedProduct.Name
				products[i].Category = updatedProduct.Category
				products[i].Price = updatedProduct.Price
				products[i].Stock = updatedProduct.Stock
				products[i].Description = updatedProduct.Description
				products[i].Status = updatedProduct.Status
				
				SendSuccessResponse(c, "Product updated successfully!", products[i])
				return
			}
		}
		SendErrorResponse(c, "Product not found", http.StatusNotFound)
	})

	router.PATCH("/products/:id/stock", AdminOnly(), func(c *gin.Context) {
		id := c.Param("id")
		newStock := c.Query("stock")
		
		stock, err := strconv.Atoi(newStock)
		if err != nil {
			SendErrorResponse(c, "Invalid stock value", http.StatusBadRequest)
			return
		}
		
		for i, product := range products {
			if fmt.Sprintf("%d", product.ID) == id {
				if stock < 0 {
					SendErrorResponse(c, "Product stock cannot be negative", http.StatusBadRequest)
					return
				}
				products[i].Stock = stock
				SendSuccessResponse(c, "Product stock updated successfully!", products[i])
				return
			}
		}
		SendErrorResponse(c, "Product not found", http.StatusNotFound)
	})

	router.DELETE("/products/:id", AdminOnly(), func(c *gin.Context) {
		id := c.Param("id")
		for i, product := range products {
			if fmt.Sprintf("%d", product.ID) == id {
				products = append(products[:i], products[i+1:]...)
				SendSuccessResponse(c, "Product deleted successfully!", nil)
				return
			}
		}
		SendErrorResponse(c, "Product not found", http.StatusNotFound)
	})
	
	router.DELETE("/orders/:id", AdminOnly(), func(c *gin.Context) {
		id := c.Param("id")
		for i, order := range orders {
			if fmt.Sprintf("%d", order.ID) == id {
				orders = append(orders[:i], orders[i+1:]...)
				SendSuccessResponse(c, "Order deleted successfully!", nil)
				return
			}
		}
		SendErrorResponse(c, "Order not found", http.StatusNotFound)
	})

	router.POST("/products/:id/image", AdminOnly(), func(c *gin.Context) {
		image := c.FormFile("image")
		id := c.Param("id")
		if image == nil {
			SendErrorResponse(c, "No image uploaded", http.StatusBadRequest)
			return
		}
		if image.Header.Get("Content-Type") != "image/jpeg" && image.Header.Get("Content-Type") != "image/png" {
			SendErrorResponse(c, "Invalid image format. Only .jpg and .png are allowed.", http.StatusBadRequest)
			return
		}
		for i, product := range products {
			if fmt.Sprintf("%d", product.ID) == id {
				products[i].Image = image
				SendSuccessResponse(c, "Product image uploaded successfully!", products[i])
				return
			}
		}
		SendErrorResponse(c, "Product not found", http.StatusNotFound)
	})

	router.Run(":8080")
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		
		c.Next()
		
		timestamp := time.Now().Format(time.RFC3339)
		method := c.Request.Method
		path := c.Request.URL.Path
		ipAddress := c.ClientIP()
		statusCode := c.Writer.Status()
		duration := time.Since(startTime)

		fmt.Printf("[%s] %s %s IP: %s Status: %d Duration: %v\n", timestamp, method, path, ipAddress, statusCode, duration)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Set("userRole", "")
		} else if authHeader == "Bearer customer-token-123" {
			c.Set("userRole", "customer")
		} else if authHeader == "Bearer staff-token-456" {
			c.Set("userRole", "staff")
		} else if authHeader == "Bearer admin-token-789" {
			c.Set("userRole", "admin")
		} else {
			SendErrorResponse(c, "Unauthorized: Invalid token", http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		
		// Cek apakah ada error yang terjadi
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			
			// Tentukan status code berdasarkan jenis error
			statusCode := http.StatusInternalServerError
			if c.Writer.Status() != http.StatusOK {
				statusCode = c.Writer.Status()
			}
			
			// Kirim response error dalam format standar
			SendErrorResponse(c, err.Error(), statusCode)
		}
	}
}

func CustomerOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists || (userRole != "" &&userRole != "customer" && userRole != "staff" && userRole != "admin") {
			SendErrorResponse(c, "Forbidden: Customers only", http.StatusForbidden)
			c.Abort()
			return
		}
		c.Next()
	}
}

func StaffOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists || userRole == "" {
			SendErrorResponse(c, "Unauthorized: Staff only", http.StatusUnauthorized)
			c.Abort()
			return
		}
		if !exists || (userRole != "staff" && userRole != "admin") {
			SendErrorResponse(c, "Forbidden: Staff only", http.StatusForbidden)
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists || userRole == "" {
			SendErrorResponse(c, "Unauthorized: Admins only", http.StatusUnauthorized)
			c.Abort()
			return
		}
		if userRole != "admin" {
			SendErrorResponse(c, "Forbidden: Admins only", http.StatusForbidden)
			c.Abort()
			return
		}
		c.Next()
	}
}

func RateLimitMiddleware() gin.HandlerFunc {
	// Map untuk menyimpan request count per IP
	requestCounts := make(map[string][]time.Time)
	
	return func(c *gin.Context) {
		// Hanya terapkan rate limit untuk endpoint /products
		if c.Request.URL.Path != "/products" {
			c.Next()
			return
		}
		
		ip := c.Client IP()
		now := time.Now()
		
		// Hapus request yang sudah lebih dari 1 menit
		if timestamps, exists := requestCounts[ip]; exists {
			var validTimestamps []time.Time
			for _, timestamp := range timestamps {
				if now.Sub(timestamp) < time.Minute {
					validTimestamps = append(validTimestamps, timestamp)
				}
			}
			requestCounts[ip] = validTimestamps
		}
		
		// Cek apakah sudah mencapai limit
		if len(requestCounts[ip]) >= 20 {
			SendErrorResponse(c, "Rate limit exceeded: Maximum 20 requests per minute for /products endpoint", http.StatusTooManyRequests)
			c.Abort()
			return
		}
		
		// Tambahkan request saat ini
		requestCounts[ip] = append(requestCounts[ip], now)
		
		c.Next()
	}
}

func SendSuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SendErrorResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
	})
}

func calculateTotalPrice(productID int, quantity int) float64 {
	for _, product := range products {
		if product.ID == productID {
			return product.Price * float64(quantity)
		}
	}
	return 0.0
}

func containsIgnoreCase(str, substr string) bool {
	return len(str) >= len(substr) && (str == substr || containsIgnoreCase(str[1:], substr))
}