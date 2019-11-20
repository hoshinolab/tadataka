package db

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/gomodule/redigo/redis"
)

var conn redis.Conn
var redisErr error

func RedisCSVLoader() {
	redisInit()

	u, _ := user.Current()
	homeDir := u.HomeDir
	tadatakaDir := filepath.Join(homeDir, ".tadataka")
	isjDir := filepath.Join(tadatakaDir, "isj")
	isjCSVPath := filepath.Join(isjDir, "isj-concat.csv")
	jukyoJushoDir := filepath.Join(tadatakaDir, "jukyoJusho")
	jukyoJushoCSVPath := filepath.Join(jukyoJushoDir, "jukyo-jusho-concat.csv")

	// load ISJ CSV to Redis DB 10
	fmt.Println("isj(位置参照情報) CSV")
	conn.Send("SELECT", 10)
	isjCSV, err := os.Open(isjCSVPath)
	if err != nil {
		panic(err)
	}

	ibs := bufio.NewScanner(isjCSV)
	for ibs.Scan() {
		sl := strings.Split(ibs.Text(), ",")
		fullGrid := sl[7]
		grid := fullGrid[:9] //shorten like 8RM327PF+
		redisValue := strings.Replace(ibs.Text(), ",", ":", -1)

		_, err := conn.Do("SADD", grid, redisValue)
		if err != nil {
			panic(err)
		}
	}

	//JukyoJuksho CSV
	fmt.Println("JukyoJusho CSV")

	conn.Send("SELECT", 11)
	jjCSV, err := os.Open(jukyoJushoCSVPath)
	if err != nil {
		panic(err)
	}

	jbs := bufio.NewScanner(jjCSV)
	for jbs.Scan() {
		sl := strings.Split(jbs.Text(), ",")
		fullGrid := sl[7]
		grid := fullGrid[:9] //shorten like 8RM327PF+
		redisValue := strings.Replace(jbs.Text(), ",", ":", -1)

		_, err := conn.Do("SADD", grid, redisValue)
		if err != nil {
			panic(err)
		}
	}

	defer conn.Close()
}

func redisInit() {
	conn, redisErr = redis.Dial("tcp", "localhost:6379")
	if redisErr != nil {
		panic(redisErr)
	}
}

func GetMembersFromList(key, targetDB string) []string {
	if conn == nil {
		redisInit()
	}

	if targetDB == "ISJ" {
		conn.Send("SELECT", 10)
	} else if targetDB == "JukyoJusho" {
		conn.Send("SELECT", 11)
	} else {
		panic("taregt DB is not designated")
	}

	s, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		panic(err)
	}
	return s
}
