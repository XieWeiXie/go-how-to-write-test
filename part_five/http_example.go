package part_five

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	_ "github.com/mkevac/debugcharts"
)

func StartServer() {
	http.HandleFunc("/simple", func(writer http.ResponseWriter, request *http.Request) {
		//timer := time.Tick(time.Second)
		//
		//go func() {
		//	for {
		//		select {
		//		case now := <-timer:
		//			nowUnix := now.Unix()
		//			fmt.Println(nowUnix)
		//		}
		//	}
		//}()

		writer.Write([]byte("hello simple"))
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello index"))
	})
	if err := http.ListenAndServe(":9876", nil); err != nil {
		log.Fatalf("HTTP Server Failed: %v", err)
	}
}
