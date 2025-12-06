package db

import (
	"context"
	"log"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func ConnectArango() driver.Database {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"}, // 도커 아랑고 주소
	})
	if err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}

	// 클라이언트 생성
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "1234"), // 계정명, 비밀번호
	})
	if err != nil {
		log.Fatalf("Failed to create ArangoDB client: %v", err)
	}

	// dev라는 DB 불러오기
	db, err := client.Database(nil, "dev")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	return db
}

// 컬렉션 자동 생성 함수
func EnsureCollection(db driver.Database, name string) driver.Collection {
	ctx := context.Background()

	// 1) 컬렉션 존재 여부 체크
	exists, err := db.CollectionExists(ctx, name)
	if err != nil {
		log.Fatalf("Failed to check collection %s: %v", name, err)
	}

	// 2) 없으면 생성
	if !exists {
		_, err := db.CreateCollection(ctx, name, &driver.CreateCollectionOptions{})
		if err != nil {
			log.Fatalf("Failed to create collection %s: %v", name, err)
		}
		log.Println("Collection created:", name)
	}

	// 3) 존재했든 새로 만들었든 컬렉션 핸들 가져와 반환
	col, err := db.Collection(ctx, name)
	if err != nil {
		log.Fatalf("Failed to get collection %s: %v", name, err)
	}

	return col
}

// Wire 에서 컬렉션 이름을 DI할 수 있게 만드는 Provider
func ProvideCollection(name string) func(driver.Database) driver.Collection {
	return func(db driver.Database) driver.Collection {
		return EnsureCollection(db, name)
	}
}
