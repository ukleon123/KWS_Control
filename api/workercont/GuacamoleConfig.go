package WorkerCont

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func GuacamoleConfig(UUID string, Ip string, PrivateKey string) {
	// MySQL 연결 정보 수정 (내부 IP 사용)
	//db, err := sql.Open("mysql", "root:3c0GBUDtUBhyqHd@tcp(10.5.12.123:3306)/guacamole_db")
	// 외부에서 접근 시:
	println("1111")
	db, err := sql.Open("mysql", "root:3c0GBUDtUBhyqHd@tcp(223.194.20.119:25118)/guacamole_db")

	if err != nil {
		println("2222")
		log.Fatal("DB 연결 실패:", err)
	}
	defer db.Close()
	println("3333")
	// 사용자 생성
	userID := UUID
	userPass := rand.Intn(10000000) + 1
	salt := "20250307"
	println("4444")
	passwordHash := fmt.Sprintf("%x", sha256Sum(strconv.Itoa(userPass)+salt))
	println("5555")
	_, err = db.Exec(`
      INSERT INTO guacamole_user (username, password_hash, password_salt, password_date)
      VALUES (?, UNHEX(?), ?, ?)
   `, userID, passwordHash, salt, time.Now())
	println("6666")
	if err != nil {
		println("7777")
		log.Fatal("User 생성 실패:", err)
	}
	fmt.Println("User 생성 완료")
	println("8888")
	// SSH Connection 생성
	res, err := db.Exec(`
      INSERT INTO guacamole_connection (connection_name, protocol)
      VALUES (?, 'ssh')
   `, "My SSH Connection")
	println("9999")
	if err != nil {

		log.Fatal("Connection 생성 실패:", err)
	}
	connID, _ := res.LastInsertId()

	fmt.Println("Connection 생성 완료, ID:", connID)

	// Connection Parameter 설정
	params := map[string]string{
		"hostname": Ip, // SSH 접속할 VM의 IP
		"port":     "22",
		"username": userID,
		"private-key": fmt.Sprintf(`-----BEGIN OPENSSH PRIVATE KEY-----
	%s
	-----END OPENSSH PRIVATE KEY-----`, PrivateKey),
	}

	for name, value := range params {
		_, err = db.Exec(`
         INSERT INTO guacamole_connection_parameter (connection_id, parameter_name, parameter_value)
         VALUES (?, ?, ?)
      `, connID, name, value)
		if err != nil {
			println("1111")
			log.Fatalf("Parameter %s 설정 실패: %v", name, err)
		}
	}
	fmt.Println("Connection Parameter 설정 완료")

	// 사용자 권한 부여
	_, err = db.Exec(`
      INSERT INTO guacamole_connection_permission (user_id, connection_id, permission)
      VALUES ((SELECT user_id FROM guacamole_user WHERE username = ?), ?, 'READ')
   `, userID, connID)
	if err != nil {
		log.Fatal("Permission 부여 실패:", err)
	}

	fmt.Println(" Permission 부여 완료")
}

func sha256Sum(input string) []byte {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hash.Sum(nil)
}
