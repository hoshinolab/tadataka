package status

import (
	"fmt"
	"tadataka/db"
)

func RedisCheck() {
	s1 := db.GetMembersFromList("8Q6MJWGC+", "ISJ")
	s2 := db.GetMembersFromList("8Q6MJWGC+", "JukyoJusho")
	fmt.Println(s1)
	fmt.Println("--")
	fmt.Println(s2)
}
