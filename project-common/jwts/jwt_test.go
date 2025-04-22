package jwts

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestCreateTokenAndParseToken(t *testing.T) {
	token := CreateToken("mikasa", "1234", "1234", 1*time.Second, 1*time.Second)
	log.Println(token.AccessToken)
	time.Sleep(2 * time.Second)
	token2, err := ParseToken(token.AccessToken, "1234")
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(token2)
}
