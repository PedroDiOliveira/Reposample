package main

func register(names []string) []string {

	len := len(names)
	var lista [string]
	mapa:= make(map[string]int)
	var nova_chave string

	for i:= 0; i < len; i++{
		chave := names[i]
		_,ok := mapa[chave]
		//nome repetido
		if ok {
			x := 1
			//loop criar nomes
			for true {
				nova_chave = names[i] + strconv.itoa(x)
				//verificar se o nome ja existe
				_, ok := mapa[nova_chave]
				if ok{
					x++
					continue
				}else{
				mapa[nova_chave] = 0 //chave > nova_chave
				lista = append(lista, nova_chave)
				break
			}
		}
	}else { // novo nome
		   mapa[chave] = 0
		   lista = append(lista,"OK")
	}
}
	return lista
}

codigo correto


//package main
//
//import (
//	"strconv"
//)
//
//func register(names []string) []string {
//	length := len(names)
//	lista := make([]string, 0)
//	mapa := make(map[string]int)
//	var novaChave string
//
//	for i := 0; i < length; i++ {
//		chave := names[i]
//		_, ok := mapa[chave]
//		// Nome repetido
//		if ok {
//			x := 1
//			// Loop para criar nomes
//			for {
//				novaChave = chave + strconv.Itoa(x)
//				// Verificar se o nome já existe
//				_, ok := mapa[novaChave]
//				if ok {
//					x++
//					continue
//				} else {
//					mapa[novaChave] = 0 // chave > novaChave
//					lista = append(lista, novaChave)
//					break
//				}
//			}
//		} else { // Novo nome
//			mapa[chave] = 0
//			lista = append(lista, "OK")
//		}
//	}
//
//	return lista
//}