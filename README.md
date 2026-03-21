# vm-ssh-collector

## 한 줄 소개
SSH로 여러 VM에 접속해 파일시스템 정보를 수집하고, 주기적으로 전송하는 Go 기반 수집기입니다.

## 저장소 성격
- 분류: 인프라 / 수집기
- 목적: VM 시스템 정보 수집, 원격 명령 실행, 주기 수집 자동화
- 핵심 기술: Go, SSH, Gin, GORM, MySQL

## 현재 구현 범위
- DB에서 수집 대상 VM과 접속 정보를 읽어옵니다.
- 각 VM에 SSH로 접속해 파일시스템 정보를 수집합니다.
- 수집 결과를 Logstash 형태의 payload로 전송합니다.
- 주기적으로 동일한 수집 작업을 반복합니다.
- HTTP API로 단일 VM 또는 여러 VM에 명령을 실행할 수 있습니다.

현재 기본 수집 로직은 디스크 파일시스템 정보 기준으로 구현되어 있습니다.

## 실행 방법

### 1. 저장소 클론
```bash
git clone https://github.com/swlee3306/vm-ssh-collector.git
cd vm-ssh-collector
```

### 2. 설정 파일 작성
기본 설정 파일 이름은 `setting.yml` 입니다.

```yaml
database:
  dsn: "root:1234@tcp(localhost:3306)/baton?parseTime=true"

collector:
  collectorTime: 1
```

### 3. 환경 변수로 덮어쓰기
필요하면 아래 환경 변수로 설정을 덮어쓸 수 있습니다.

```bash
export BATON_SETTING_FILENAME=setting.yml
export BATON_DB_DSN='root:1234@tcp(localhost:3306)/baton?parseTime=true'
export BATON_DATABASE_MAX_IDLE_CONNS=10
export BATON_DATABASE_MAX_LIFETIME_HOUR=1
export BATON_COLLECT_TIME=1
export OKE_DEBUG=false
```

### 4. 실행
```bash
go run .
```

## HTTP API

### 단일 VM 명령 실행
```http
POST /api/v1/sshCommend
Content-Type: application/json

{
  "user": "ubuntu",
  "pwd": "password",
  "host_ip": "10.0.0.10",
  "port": 22,
  "cmd": "df -h"
}
```

### 등록된 VM 전체 명령 실행
```http
POST /api/v1/sshCommendAll
Content-Type: application/json

{
  "cmd": "hostname"
}
```

## 주요 흐름
1. `main.go`에서 설정 파일과 환경 변수를 읽습니다.
2. DB 연결 후 수집 대상 VM 목록을 로드합니다.
3. 초기 수집을 1회 수행합니다.
4. `collectorTime` 기준으로 주기 수집을 반복합니다.
5. Gin 서버에서 SSH 명령 실행 API를 제공합니다.

## 주요 파일
```text
vm-ssh-collector/
├── main.go
├── main_LoadEnv.go
├── main_LoadYml.go
├── collector/
│   └── collector.go
├── internal/
│   ├── dblinker/
│   ├── lslinker/
│   ├── ssh/
│   ├── sysdef/
│   └── sysenv/
├── utils/
│   └── router/
└── setting.yml
```

## 검증
```bash
go test ./...
```

## 핀 저장소용 설명 문구
SSH로 VM 파일시스템 메트릭을 수집하고 주기적으로 전송하는 Go collector
