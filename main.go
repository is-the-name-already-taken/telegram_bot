package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("테스트 서비스가 시작되었습니다!")

	ticker := time.NewTicker(1 * time.Second)
	counter := 0

	for {
		select {
		case t := <-ticker.C:
			counter++
			fmt.Printf("[%s] 서비스가 잘 작동 중입니다. (카운트: %d)\n", t.Format("15:04:05"), counter)

			// 10초(10회)마다 의도적으로 죽여서 재시작 테스트
			if counter >= 10 {
				log.Fatal("의도적인 종료: 시스템이 5초 뒤에 다시 살리는지 확인합니다.")
			}
		}
	}
}
