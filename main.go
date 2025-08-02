package main

import (
  "fmt"
  "log"
  "net/http"
  "sync"

  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool {
    return true
  },
}

// Player 구조체: 연결 정보와 선택 저장
type Player struct {
  conn    *websocket.Conn
  choice  string
	matchChan chan *Player
}

var (
  waitingPlayer *Player      // 매칭 대기 중인 플레이어
  mu            sync.Mutex   // 동시 접근 막기 위한 뮤텍스
)

// 가위바위보 승패 판단
func determineResult(p1, p2 *Player) (string, string) {
  c1, c2 := p1.choice, p2.choice

  if c1 == c2 {
    return "무승부", "무승부"
  }

  winMap := map[string]string{
    "가위": "보",
    "바위": "가위",
    "보":  "바위",
  }

  if winMap[c1] == c2 {
    return "승", "패"
  }
  return "패", "승"
}

// 두 플레이어가 매칭되면 결과 판단 후 전송
func matchPlayers(p1, p2 *Player) {
  result1, result2 := determineResult(p1, p2)

  p1.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("상대: %s / 결과: %s", p2.choice, result1)))
  p2.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("상대: %s / 결과: %s", p1.choice, result2)))

  p1.conn.Close()
  p2.conn.Close()
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Println("업그레이드 실패:", err)
    return
  }

  defer func() {
    if r := recover(); r != nil {
      conn.Close()
    }
  }()

  _, msg, err := conn.ReadMessage()
  if err != nil {
    log.Println("메시지 읽기 실패:", err)
    conn.Close()
    return
  }

  choice := string(msg)
  if choice != "가위" && choice != "바위" && choice != "보" {
    conn.WriteMessage(websocket.TextMessage, []byte("가위/바위/보 중 하나를 보내주세요."))
    conn.Close()
    return
  }

  player := &Player{conn: conn, choice: choice}

  mu.Lock()
  if waitingPlayer == nil {
    player.matchChan = make(chan *Player)
		waitingPlayer = player
		mu.Unlock()

		log.Println("매칭 대기 중...")

		opponent := <-player.matchChan
		matchPlayers(player, opponent)
  } else {
		opponent := waitingPlayer
		waitingPlayer = nil
		mu.Unlock()

		opponent.matchChan <- player
	}
}

func main() {
  http.HandleFunc("/ws", wsHandler)

  port := 8000
  log.Printf("웹소켓 서버 실행 중 : http://localhost:%d/ws", port)
  err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
  if err != nil {
    log.Fatal("서버 시작 실패:", err)
  }
}