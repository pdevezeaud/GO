package main

import (
	"fmt"
	"strings"
)

// func MyFunc() (type1, type2)  {
//     return val1,val2
//  }

// on passe en parametre le nom que l'on veut mettre en majuscule
// puis le retour  (string, int)

func TolowerStr(name string) (string, int) {

	return strings.ToLower(name), len(name)

}

func main() {
	nomcapitale, longueur := TolowerStr("ALICE")
	fmt.Printf("%s (lens=%d) \n", nomcapitale, longueur)

	//dans le cas ou l'on veut ignorer un retour dans ce cas la mise en capitale
	_, longueur = TolowerStr("Bob")

	fmt.Printf("%s longueur=%d\n", "bob", longueur)
}
