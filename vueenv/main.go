package main
import (
    "os"
    "log"
    "fmt"
    "bufio"
    "strings"
)

func readEnv(environments map[string]string) {
	file, err := os.Open(".env")

	if err != nil {
		log.Fatal("failed to open env file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line []string

	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	file.Close()

  // environments := make(map[string]string)
	for _, eachLine := range line {
    data := strings.Split(eachLine, "=")

    if len(data) == 2 && strings.HasPrefix(data[0], "VUE_APP_") {
      environments[data[0]] = data[1]
    }
	}
}

func main() {
    // 사용법 vueenv A3S_PORT 9999
    // => .env 파일에 VUE_APP_A3S_PORT=9999 라고 출력 (덮어쓰기 안됨 아래에 추가)

    if len(os.Args[1:]) < 2 {
      log.Fatal("Usage: vueenv [key] [value]")
    }

    var newKey string = os.Args[1]
    if !strings.HasPrefix(os.Args[1], "VUE_APP_") {
      newKey = fmt.Sprintf("VUE_APP_%s", os.Args[1])
    }
    newValue := os.Args[2]

    environments := make(map[string]string)
    if _, err := os.Stat(".env"); err == nil {
      // fmt.Println(newKey, newValue)
      readEnv(environments)
    }

    // 기존에 있는 값이 있으면 덮어씀
    environments[newKey] = newValue

    // 파일은 overwrite 모드로 오픈
    // file, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    file, err := os.OpenFile(".env", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
      log.Fatal(err)
    }
    defer file.Close()

    fmt.Println("-- env list-----")
    for key, value := range environments {
      line := fmt.Sprintf("%s=%s\n", key, value)
      fmt.Print(line)

      _, err = file.WriteString(line)
      if err != nil {
        log.Fatal(err)
      }
    }
}
