package main

import "fmt"

type User struct {
	Id   int64
	Name string
}

func main() {
	var defaultMap map[int64]string
	fmt.Printf("Type: %T  Value:%v\n", defaultMap, defaultMap)

	mapByMake := make(map[int64]string)
	fmt.Printf("Type: %T  Value:%v\n", mapByMake, mapByMake)

	mapByMakeWithCap := make(map[int64]string, 3)
	fmt.Printf("Type: %T  Value:%v\n", mapByMakeWithCap, mapByMakeWithCap)

	mapByLiteral := map[int64]string{18: "Ako", 17: "Era", 3: "Erom"}
	fmt.Printf("Type: %T  Value:%v\n", mapByLiteral, mapByLiteral)

	mapByMake[1] = "Vasya"
	fmt.Printf("Type: %T  Value:%v\n", mapByMake, mapByMake)

	delete(mapByMake, 1)
	fmt.Printf("Type: %T  Value:%v\n", mapByMake, mapByMake)

	mapForIteration := map[string]int{"First": 1, "Second": 2, "Third": 3}
	for key, value := range mapForIteration {
		fmt.Printf("Key:%s Value:%d\n", key, value)
	}

	users := []User{
		{1, "Ako"},
		{2, "Era"},
		{3, "Vasya"},
		{2, "Erom"},
	}
	fmt.Printf("Type: %T  Value:%v\n", users, users)

	uniqueUsers := make(map[int64]struct{}, len(users))
	for _, user := range users {
		if _, ok := uniqueUsers[user.Id]; !ok {
			uniqueUsers[user.Id] = struct{}{}
		}
	}
	fmt.Printf("Type: %T  Value:%v\n", uniqueUsers, uniqueUsers)

	fmt.Println(findInSlice(1, users))
	fmt.Println(findInSlice(2, users))

	//Поиск О(1) вместо О(n)
	fmt.Println(findInMap(18, mapByLiteral))

}

func findInSlice(id int64, users []User) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

// Поиск О(1) вместо О(n)
func findInMap(id int64, users map[int64]string) *string {
	if user, ok := users[id]; ok {
		return &user
	}

	return nil
}
