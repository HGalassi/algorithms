package main

import (
	"fmt"

	ctt "github.com/HGalassi/algorithms/go/internal/big_o/constant_time"
	cbt "github.com/HGalassi/algorithms/go/internal/big_o/cubic_time"
	ext "github.com/HGalassi/algorithms/go/internal/big_o/exponential_time"
	lt "github.com/HGalassi/algorithms/go/internal/big_o/linear_time"
	llt "github.com/HGalassi/algorithms/go/internal/big_o/log_linear_time"
	lgt "github.com/HGalassi/algorithms/go/internal/big_o/logarithmic_time"
	qdt "github.com/HGalassi/algorithms/go/internal/big_o/quadratic_time"
	"github.com/HGalassi/algorithms/go/internal/timer"
)

func main() {

	//Linear Time O(n)
	var customers = []string{"Alice", "Bob", "Charlie", "David", "Eve", "Elis", "Frank", "Grace", "Heidi", "Ivan", "Judy", "Mallory",
		"Niaj", "Olivia", "Peggy", "Quentin", "Rupert", "Sybil", "Trent", "Uma"}

	// Constant Time O(1)
	ctt.ConstantTime()
	// Logarithmic Time O(log n)
	lgt.Logarithmic_time(customers, "Eve")
	// Linear Time O(n)
	lt.LinearTime(customers)
	// Logarithmic Time O(log n) increase
	lgt.Append_customers_for_logarithmic_time_increase(&customers, []string{"Frank", "Grace", "Heidi", "Ivan", "Judy", "Mallory", "Niaj",
		"Olivia", "Peggy", "Quentin", "Rupert", "Sybil", "Trent", "Uma", "Victor",
		"Wendy", "Xander", "Yvonne", "Zara", "Aaron", "Betty", "Cathy", "Derek", "Ellen",
		"Franklin", "Gina", "Hank", "Irene", "Jack", "Kathy", "Liam", "Mona", "Nate", "Olga",
		"Paul", "Queen", "Ralph", "Sandy", "Tom", "Ursula", "Vince", "Willa", "Xenia", "Yosef", "Zane",
		"Uma", "Victor", "Wendy", "Xander", "Yvonne", "Zara", "Aaron", "Betty", "Cathy", "Derek", "Ellen",
		"Franklin", "Gina", "Hank", "Irene", "Jack", "Kathy", "Liam", "Mona", "Nate", "Olga",
		"Paul", "Queen", "Ralph", "Sandy", "Tom", "Ursula", "Vince", "Willa", "Xenia", "Yosef", "Zane", "Tina", "Lucas",
		"Mariana", "Roberto", "Fernanda", "Gabriel", "Isabela", "Juliana", "Leonardo", "Camila", "Eduardo", "Bianca", "Rafael", "Larissa",
		"Matheus", "Sofia", "Thiago", "Amanda", "Bruno", "Carolina", "Diego", "Elena", "Felipe", "Giovanna", "Henrique", "Júlia",
		"Kevin", "Letícia", "Murilo", "Natália", "Otávio", "Priscila", "Quésia", "Renato", "Samara", "Tatiana", "Vitor",
		"Wellington", "Ximena", "Yara", "Zeca", "Adriana", "Bruno", "Cecília", "Daniel", "Elisa", "Fábio", "Glória", "Helena", "Igor", "Janaína", "Kléber", "Lúcia", "Marcos", "Nina", "Otto", "Paula", "Quirino", "Rosa", "Samuel", "Tânia",
		"Ulisses", "Valéria", "Wilson", "Xuxa", "Yuri", "Zilda", "Ana", "Bruno", "Clara", "Diego", "Eva", "Félix", "Gustavo",
		"Hugo", "Isis", "João", "Kátia", "Leandro", "Marta", "Nicolas", "Olívia", "Pedro", "Queila", "Renan", "Sílvia",
		"Tiago", "Ubirajara", "Vanessa", "Wagner", "Xênia", "Yasmin", "Zuleica",
	})
	lgt.Logarithmic_time(customers, "Zuleica")

	//Log Linear Time O(n log n)
	llt.Log_linear_time()

	// Quadratic Time O(n^2)
	qdt.Quadratic_time()

	// Cubic Time O(n^3)
	cbt.Cubic_time()

	// Exponential Time O(2^n)
	ext.Exponential_time()

	timer.PrintExecutions()

	response := Fatorial(5)
	fmt.Printf("Fatorial de 5 é: %d\n", response)

	Challenge(1)
	Linear_search_example()
	Binary_search_example()
	fmt.Print(Challenge_binary_search())
	arr := []int{1, 3, 2, 6}
	Bubble_sort(&arr)
	fmt.Printf("%v\n", arr)
}
