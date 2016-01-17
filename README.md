# Introduction

`file_finder` is a simple file finder that (can recursively) find files you're looking for based on a simple string.

## Command
`$ file_finder -file README -recurse -dir ~/go/src/github.com/cwramsey`

`-recurse` is optional.

### Example

```
$ file_finder -file arrow.png -recurse -dir ~/Projects

Match: /Users/cwramsey/Projects/resume/arrow.png
Match: /Users/cwramsey/Projects/twitterbot/images/comment-arrow.png
```

### ToDo

* Learn more about the `flag` module
* Add glob searches
