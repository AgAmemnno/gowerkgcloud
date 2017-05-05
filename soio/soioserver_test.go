package soio

import (
	//"fmt"
	"log"
	"net/http"
        "testing"
	"github.com/googollee/go-socket.io"
        "time"
)

func TestSoio(t *testing.T){
        server, err := Soios()
        
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			/*
                        m := make(map[string]interface{})
			m["a"] = "你好"
			e := so.Emit("cn1111", m["a"])
			//这个没有问题
			fmt.Println(m["a"])

			b := make(map[string]string)
			b["u-a"] = "中文内容" //这个不能是中文
			m["b-c"] = b
			e = so.Emit("cn2222", m)
			log.Println(e)
                         */
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
                        t.Log("chat message.")
		})
		// Socket.io acknowledgement example
		// The return type may vary depending on whether you will return
		// For this example it is "string" type
		so.On("chat message with ack", func(msg string) string {
                        t.Log("chat message with ack.")
			return msg
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

        passmain(server,err)
}


func Soios() (*socketio.Server, error) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
                return nil,err
	}
   return server, err
}


func passmain(s *socketio.Server,e error) {
 // 非同期書き込みされる例
        http.Handle("/socket.io/", s)
	http.Handle("/", http.FileServer(http.Dir("asset")))
        log.Println("Serving at localhost:3641...")

  //go func() {
    time.Sleep(100 * time.Millisecond) // 他にいろんな処理があって書き込みに時間がかかるとする…
    
  //}
       
	
//        return s,e

}
