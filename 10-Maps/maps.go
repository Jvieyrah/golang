package main

import "fmt"

func main() {
	fmt.Println("Maps")
	usuarios := map[string]string{
		"nome":      "joao",
		"profissao": "Dev",
	}

	fmt.Println(usuarios)
	fmt.Println(usuarios["nome"])

	usuario2 := map[string]map[string]string{
		"nome": map[string]string{
			"primeiro": "joao",
			"ultimo":   "filho",
		},
		"profissao": map[string]string{
			"primeiro": "dev",
			"ultimo":   "pleno",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["primeiro"])
	fmt.Println(usuario2["profissao"]["ultimo"])

	delete(usuario2, "profissao")
	fmt.Println(usuario2)

	usuario2["profissao"] = map[string]string{
		"primeiro": "dev",
		"ultimo":   "pleno",
	}
	fmt.Println(usuario2)

}
