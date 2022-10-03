package practices

import (
	"regexp"
	"strings"
	"sync"
)

const content = `
  <html>
    <head>
      <title>Test Title</title>
    </head>
    <body>
      <h1>Title 1</h1><p>Lorem ipsum 1</p>
      <h1>Title 2</h1>
      <p>Lorem ipsum 2</p>
      <div class="top">
          <div class="mid">
              <div class="bot">
                  <p>Lorem ipsum 3</p>
              </div>
          </div>
      </div>
    </body>
  </html>
`

func FindTags(s string) []string {
	lines := strings.Split(s, "\n")
	tags := make(chan string)

	var wg sync.WaitGroup
	for _, line := range lines {
		line := line
		wg.Add(1)
		go func() {
			defer wg.Done()
			findTags(line, tags)
		}()
	}

	go func() {
		wg.Wait()
		close(tags)
	}()

	var res []string
	for tag := range tags {
		res = append(res, tag)
	}

	return res
}

var tagName = regexp.MustCompile(`<([a-zA-Z0-9_\-]+)\s?.*>`)

func findTags(s string, tags chan string) {
	matches := tagName.FindAllStringSubmatch(s, -1)
	for _, match := range matches {
		if len(match) > 0 {
			tags <- match[1]
		}
	}
}
