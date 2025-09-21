package main

import (
	cbt "github.com/HGalassi/algorithms/go/internal/cubic_time"
	ext "github.com/HGalassi/algorithms/go/internal/exponential_time"
	lt "github.com/HGalassi/algorithms/go/internal/linear_time"
	llt "github.com/HGalassi/algorithms/go/internal/log_linear_time"
	lgt "github.com/HGalassi/algorithms/go/internal/logarithmic_time"
	qdt "github.com/HGalassi/algorithms/go/internal/quadratic_time"
	"github.com/HGalassi/algorithms/go/internal/timer"
)

func main() {

	//Linear Time O(n)
	var customers = []string{"Alice", "Bob", "Charlie", "David", "Eve", "Elis"}
	lt.LinearTime(customers)
	// Logarithmic Time O(log n)
	lgt.Logarithmic_time(customers, "Eve")
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

}
