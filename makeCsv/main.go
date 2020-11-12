package main
import (
    "bufio"
    "encoding/csv"
    "os"
    "flag"
    "fmt"
)

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

    for i := 0; i < *line; i++ {
      // csv 내용 쓰기
      wr.Write([]string{"A", "0.25"})
      wr.Write([]string{"B", "55.70"})
    }
    wr.Flush()
}
