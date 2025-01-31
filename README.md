# Rule Maker
프로젝트 디렉토리 구조를 .cursorrules 파일로 생성하는 도구입니다.

# 커서 규칙 사용 방법
커서(Curer)는 개발자들이 AI를 활용하여 코드를 작성하고 수정하는 데 도움을 주는 강력한 도구로, 특정 지침을 설정할 수 있는 "커서 규칙(Cursor Rules)" 기능을 제공합니다. 이 기능은 개발 환경에 맞게 AI의 행동을 커스터마이징하여 더욱 개인화된 코딩 경험을 가능하게 합니다.

## 커서 규칙이란 무엇인가?
커서 규칙은 커서 AI의 행동을 안내하는 사용자 지정 지침입니다. 이 규칙을 통해 코드를 해석하고, 제안하며, 질의에 응답하는 AI의 동작을 세부적으로 조정할 수 있습니다. 커서 규칙에는 두 가지 주요 유형이 있습니다:
- 글로벌 규칙 (Global Rules): 모든 프로젝트에 적용되는 규칙으로, 커서 설정의 '일반 > AI 규칙'에서 설정할 수 있습니다.
- 프로젝트별 규칙 (Project-Specific Rules): 특정 프로젝트에만 적용되는 규칙으로, 프로젝트 루트 디렉토리에 .cursorrules 파일을 생성하여 설정할 수 있습니다.

## 커서 규칙의 중요성
- 프로젝트 맥락 이해
- AI에게 프로젝트 구조, 사용 기술, 특정 요청사항을 제공하여 프로젝트 맥락을 더 잘 이해하도록 할 수 있습니다.

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

## Cursor rules
https://github.com/PatrickJS/awesome-cursorrules

https://dotcursorrules.com
