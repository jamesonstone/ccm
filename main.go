package main

import "fmt"

func main() {
    fmt.Println(`| Commit Type | Title                    | Description                                                                                                 | Emoji  | Emoji Code            |
| ----------- | ------------------------ | ----------------------------------------------------------------------------------------------------------- |:------:|:---------------------:|
| 'feat'      | Features                 | A new feature                                                                                               | ✨     | :sparkles:            |
| 'fix'       | Bug Fixes                | A bug Fix                                                                                                   | 🐛     | :bug:                 |
| 'docs'      | Documentation            | Documentation only changes                                                                                  | 📓     | :notebook:            |
| 'style'     | Styles                   | Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)      | 🎨     | :art:                 |
| 'refactor'  | Code Refactoring         | A code change that neither fixes a bug nor adds a feature                                                   | ♻️      | :recycle:             |
| 'perf'      | Performance Improvements | A code change that improves performance                                                                     | ⚡     | :zap:                 |
| 'test'      | Tests                    | Adding missing tests or correcting existing tests                                                           | 🧪     | :test_tube:           |
| 'build'     | Builds                   | Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)         | 🛠      | :hammer_and_wrench:   |
| 'ci'        | Continuous Integrations  | Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs) | 👷     | :construction_worker: |
| 'chore'     | Chores                   | Other changes that don't modify src or test files                                                           | 🔧     | :wrench:              |
| 'revert'    | Reverts                  | Reverts a previous commit                                                                                   | 🗑      | :wastebasket:         |
`)
}
