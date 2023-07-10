package main

import "fmt"

func ArrayMerge(ArrayA, ArrayB []string) []string {

	ArrayBaru := make([]string, 0)

	for _, elemen := range ArrayA {
		if !cek(ArrayBaru, elemen) {
			ArrayBaru = append(ArrayBaru, elemen)
		}
	}

	for _, elemen := range ArrayB {
		if !cek(ArrayBaru, elemen) {
			ArrayBaru = append(ArrayBaru, elemen)
		}
	}

	return ArrayBaru
}

func cek(slice []string, elemen string) bool {
	for _, item := range slice {
		if item == elemen {
			return true
		}
	}
	return false
}

func main() {

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))

}
