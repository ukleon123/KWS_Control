package WorkerCont

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
)

func removePEMBoundaries(pemStr string) string {
	lines := strings.Split(pemStr, "\n")
	var filteredLines []string

	for _, line := range lines {
		if !strings.HasPrefix(line, "-----") { // 헤더/푸터 제외
			filteredLines = append(filteredLines, line)
		}
	}

	return strings.Join(filteredLines, "")
}

func SshKeygen() (privateKeyPEM, publicKeyPEM string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", fmt.Errorf("키 생성 실패: %v", err)
	}

	privKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	pubKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", fmt.Errorf("퍼블릭 키 변환 실패: %v", err)
	}
	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	})

	// 헤더와 푸터 제거
	cleanedPrivKey := removePEMBoundaries(string(privKeyPEM))
	cleanedPubKey := removePEMBoundaries(string(pubKeyPEM))

	return cleanedPrivKey, cleanedPubKey, nil
}
