### leap
cli app to leap from one thought to another


### features
- bookmarking
- context switching through projects
    - has bookmarks
    - annotations
    - updates
    - todo lists?

### libraries
- spf13/cobra
- go-sqlite3
- text/template
- charmbracelet/log
- charmbracelet/lipgloss

### tap, install and update
```
brew tap ajilk/leap http://github.com/ajilk/leap
brew install leap
brew update
```

### build && run
```
go build && ./leap
```

### deploy
```
git tag v0.0.0
git push --tags
goreleaser release --clean
```
