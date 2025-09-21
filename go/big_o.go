package big_o

import (
	"fmt"
	"time"
)

func main() {

	// Example of Linear Time O(n)
	var customers = []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	linear_time(customers)

	logarithmic_time(customers, "Eve")
	append_customers_for_logarithmic_time_increase(customers, []string{"Frank", "Grace", "Heidi", "Ivan", "Judy", "Mallory", "Niaj",
		"Olivia", "Peggy", "Quentin", "Rupert", "Sybil", "Trent", "Uma", "Victor",
		"Wendy", "Xander", "Yvonne", "Zara", "Aaron", "Betty", "Cathy", "Derek", "Ellen",
		"Franklin", "Gina", "Hank", "Irene", "Jack", "Kathy", "Liam", "Mona", "Nate", "Olga",
		"Paul", "Queen", "Ralph", "Sandy", "Tom", "Ursula", "Vince", "Willa", "Xenia", "Yosef", "Zane"})
	logarithmic_time(customers, "Zane")

}

func linear_time(customers []string) {
	// T(n) = 1
	start := time.Now()
	free_books := customers[0]
	fmt.Printf("Free book goes to: %s\n", free_books)
	elapsed := time.Since(start)
	fmt.Printf("Linear Time O(n) took %s\n", elapsed)
}

func logarithmic_time(customers []string, target string) {
	start := time.Now()
	// Example of Logarithmic Time O(log n)
	for i := 0; i < len(customers); i++ {
		if customers[i] == target {
			fmt.Printf("Found customer: %s\n", customers[i])
			break
		}
	}
	logarithmic_time(customers, target)
	elapsed := time.Since(start)
	fmt.Printf("Logarithmic Time O(log n) took %s\n", elapsed)
}

func append_customers_for_logarithmic_time_increase(customers []string, new_customers []string) []string {
	customers = append(customers, new_customers...)
	return customers
}
