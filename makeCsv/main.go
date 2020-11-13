package main
import (
    "bufio"
    "encoding/csv"
    "os"
    "flag"
    "fmt"
    "github.com/goombaio/namegenerator"
    "math/rand"
    "strings"
    "time"
    "strconv"
)

func randomString() (result string) {
  rand.Seed(time.Now().Unix())

  charSet := "abcdedfghijklmnopqrstQWERTYUIOPASDFGHJKLZXCVBNM"
  var output strings.Builder
  length := rand.Intn(200 - 50) + 50 // 최소 50, 최대 200 문자

  for i := 0; i < length; i++ {
    random := rand.Intn(len(charSet))
    randomChar := charSet[random]
    output.WriteString(string(randomChar))
  }
  result = output.String()
  output.Reset()

  return result
}

func main() {
    // 라인 수 입력받기
    fileName := flag.String("f", "output.csv", "파일명")
    line := flag.Int("l", 0, "파일 라인 수")

    flag.Parse() // 명령줄 옵션의 내용을 각 자료형별로 분석

    if flag.NFlag() == 0 { // 명령줄 옵션의 개수가 0개이면
      flag.Usage()   // 명령줄 옵션 기본 사용법 출력
      return
    }

    path := fmt.Sprintf("./%s", *fileName)

    fmt.Printf( "파일 생성 중... (%s)\n", path)

    // 파일 생성
    file, err := os.Create(path)
    if err != nil {
        panic(err)
    }

    // csv writer 생성
    wr := csv.NewWriter(bufio.NewWriter(file))

    // header
    wr.Write([]string{"name", "age", "sub"})

    seed := time.Now().UTC().UnixNano()
    nameGenerator := namegenerator.NewNameGenerator(seed)

    for i := 0; i < *line; i++ {
      name := nameGenerator.Generate()
      age := rand.Intn(100)
      sub := randomString()
      
      wr.Write([]string{name, strconv.Itoa(age), sub})
    }
    wr.Flush()
}
