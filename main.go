package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type GitignoreRule struct {
	pattern  string
	basePath string // 이 규칙이 적용되는 기준 디렉토리
}

func loadAllGitignorePatterns(projectDir string) ([]GitignoreRule, error) {
	var rules []GitignoreRule

	err := filepath.WalkDir(projectDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.Name() == ".gitignore" {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			basePath := filepath.Dir(path)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				pattern := strings.TrimSpace(scanner.Text())
				if pattern != "" && !strings.HasPrefix(pattern, "#") {
					rules = append(rules, GitignoreRule{
						pattern:  pattern,
						basePath: basePath,
					})
				}
			}
		}
		return nil
	})

	return rules, err
}

func isIgnored(path string, rules []GitignoreRule) bool {
	if strings.HasPrefix(filepath.Base(path), ".") {
		return true
	}

	for _, rule := range rules {
		// 현재 파일이 규칙의 기준 디렉토리 아래에 있는지 확인
		relPath, err := filepath.Rel(rule.basePath, path)
		if err != nil || strings.HasPrefix(relPath, "..") {
			continue // 이 규칙의 범위 밖
		}

		matched, err := filepath.Match(rule.pattern, filepath.Base(path))
		if err == nil && matched {
			return true
		}

		// 디렉토리 패턴 처리 (예: node_modules/)
		if strings.HasSuffix(rule.pattern, "/") {
			dirPattern := strings.TrimSuffix(rule.pattern, "/")
			if strings.Contains(relPath, dirPattern) {
				return true
			}
		}

		// 전체 경로 패턴 처리
		if strings.Contains(rule.pattern, "/") {
			matched, err := filepath.Match(rule.pattern, relPath)
			if err == nil && matched {
				return true
			}
		}
	}
	return false
}

func writeRules(structure *strings.Builder) {
	structure.WriteString("# Default Rules\n")
	structure.WriteString("- Always write in 한국어\n\n")
	structure.WriteString("# Commit Message Creation Rules\n")
	structure.WriteString("- Write in English\n")
	structure.WriteString("- Write simply in one line\n")
	structure.WriteString("- Use fix, chore, docs, refactor, style, test\n")
	structure.WriteString("- feat(api)\n")
	structure.WriteString("- feat(web)\n")
	structure.WriteString("- feat(apps)\n")
	structure.WriteString("- feat(lib)\n\n")
	structure.WriteString("# Text & Message Rules\n")
	structure.WriteString("- All text must be written in i18n.\n")
	structure.WriteString("- The key is set to a lowercase and hyphen (-) combination.\n\n")
	structure.WriteString("# Project Structure\n")
}

type FileEntry struct {
	path   string
	name   string
	isDir  bool
	parent string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("프로젝트 디렉토리 경로를 입력하세요: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	projectDir := strings.Trim(strings.TrimSpace(input), "'")

	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		log.Fatal("디렉토리가 존재하지 않습니다:", projectDir)
	}

	// 모든 .gitignore 파일의 규칙을 로드
	ignoreRules, err := loadAllGitignorePatterns(projectDir)
	if err != nil {
		log.Fatal("gitignore 규칙을 로드하는 중 오류 발생:", err)
	}

	dirContents := make(map[string][]FileEntry)

	err = filepath.WalkDir(projectDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if isIgnored(path, ignoreRules) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if path == projectDir {
			return nil
		}

		parent := filepath.Dir(path)
		dirContents[parent] = append(dirContents[parent], FileEntry{
			path:   path,
			name:   d.Name(),
			isDir:  d.IsDir(),
			parent: parent,
		})

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// 각 디렉토리 내에서 정렬
	for dir := range dirContents {
		sort.Slice(dirContents[dir], func(i, j int) bool {
			// 같은 디렉토리 내에서 폴더 먼저, 그 다음 파일
			if dirContents[dir][i].isDir != dirContents[dir][j].isDir {
				return dirContents[dir][i].isDir
			}
			// 같은 타입이면 이름 순
			return dirContents[dir][i].name < dirContents[dir][j].name
		})
	}

	var structure strings.Builder

	// 규칙들 먼저 작성
	writeRules(&structure)

	// 프로젝트 이름 (하이픈 없이)
	structure.WriteString(filepath.Base(projectDir) + "/\n")

	var printDir func(dir string, depth int)
	printDir = func(dir string, depth int) {
		entries := dirContents[dir]
		prefix := strings.Repeat("-", depth) + " "

		for _, entry := range entries {
			if entry.isDir {
				structure.WriteString(prefix + entry.name + "/\n")
				printDir(entry.path, depth+1)
			} else {
				structure.WriteString(prefix + entry.name + "\n")
			}
		}
	}

	printDir(projectDir, 1)

	cursorrules := filepath.Join(projectDir, ".cursorrules")
	err = os.WriteFile(cursorrules, []byte(structure.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n프로젝트 구조가 다음 파일에 저장되었습니다:", cursorrules)
	fmt.Println("\n프로젝트 구조:")
	fmt.Println(structure.String())
}
