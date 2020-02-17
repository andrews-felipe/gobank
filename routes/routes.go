import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

)

func main() {
    router := mux.NewRouter()
    // Create
    router.HandleFunc("/transfers", getTransfers).Methods("POST")
    // Read
    router.HandleFunc("/orders/{orderId}", getOrder).Methods("GET")
    // Read-all
    router.HandleFunc("/orders", getOrders).Methods("GET")
    // Update
    router.HandleFunc("/orders/{orderId}", updateOrder).Methods("PUT")
    // Delete
    router.HandleFunc("/orders/{orderId}", deleteOrder).Methods("DELETE")
    // Initialize db connection
    initDB()

    log.Fatal(http.ListenAndServe(":8080", router))
}