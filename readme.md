# Fiber Project

이 프로젝트는 Go Fiber 프레임워크를 사용하는 웹 애플리케이션입니다.

## 📂 프로젝트 구조 (Project Structure)

```text
fiber/
├── cmd/
│   └── api/                  # 애플리케이션 진입점 (Main entry point)
│       └── main.go           # 서버 실행 파일
├── internal/
│   ├── db/                   # 데이터베이스 연결 및 설정
│   │   └── arango.go         # ArangoDB 연결 설정
│   ├── entity/               # 도메인 모델 (Struct 정의)
│   ├── middleware/           # HTTP 미들웨어 (CORS, Logger 등)
│   │   └── cors.go           # CORS 설정
│   ├── module/               # 비즈니스 로직 모듈 (Domain Modules)
│   │   ├── board/            # 게시판 모듈
│   │   ├── reply/            # 댓글 모듈
│   │   └── user/             # 사용자 모듈
│   │       ├── controller/   # HTTP 요청 처리 핸들러
│   │       ├── service/      # 비즈니스 로직
│   │       ├── repository/   # 데이터 액세스 계층
│   │       └── module.go     # 모듈 설정 (Wire DI 등)
│   └── util/                 # 공통 유틸리티 함수
├── go.mod                    # 모듈 정의 파일
└── go.sum                    # 모듈 체크섬 파일
```

## 🛠 주요 디렉토리 설명

- **cmd/api**: API 서버를 구동하기 위한 `main` 함수가 위치합니다.
- **internal**: 외부에서 임포트할 수 없는 비공개 패키지들이 위치합니다.
  - **db**: DB 연결 초기화 및 설정을 담당합니다. (예: ArangoDB)
  - **entity**: DB 테이블과 매핑되거나 로직에 사용되는 데이터 구조체를 정의합니다.
  - **middleware**: HTTP 요청 처리를 위한 미들웨어들이 위치합니다. (예: CORS)
  - **module**: 각 도메인별 기능이 모듈화되어 있습니다.
    - **board**: 게시글 관련 기능을 담당합니다.
    - **reply**: 댓글 관련 기능을 담당합니다.
    - **user**: 사용자 인증 및 정보 관리를 담당합니다.
    - 각 모듈은 일반적으로 `controller` (요청/응답), `service` (로직), `repository` (DB접근)로 계층화되어 있습니다.
  - **util**: 프로젝트 전반에서 재사용 가능한 헬퍼 함수들을 모아둡니다.

## 🌐 CORS 설정 (CORS Configuration)

이 프로젝트는 `internal/middleware/cors.go`에서 CORS(Cross-Origin Resource Sharing) 정책을 관리합니다.

- **Allowed Origins**:
  - `http://localhost:3000`
  - `http://localhost:5173`
- **Allowed Headers**: `*` (모든 헤더 허용)
- **Allowed Methods**: `*` (모든 메서드 허용)
- **Allow Credentials**: `true` (자격 증명 허용)
