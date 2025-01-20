# Rule Maker
프로젝트 디렉토리 구조를 .cursorrules 파일로 생성하는 도구입니다.

## 기능
- 프로젝트의 디렉토리 구조를 tree 형태로 표시
- .gitignore 규칙 적용 (모든 하위 디렉토리의 .gitignore 포함)
- 결과를 .cursorrules 파일로 저장

## 실행 방법
1. 직접 실행:
go run main.go

2. 빌드 후 실행:
go build -o maker main.go
./maker

## 사용 방법
1. 프로그램 실행
2. 프로젝트 디렉토리 경로 입력 (드래그 앤 드롭 가능)
3. .cursorrules 파일이 해당 디렉토리에 생성됨