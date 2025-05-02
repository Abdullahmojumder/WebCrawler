*WebCrawler*
A concurrent web crawler written in Go that crawls a website, extracts internal links, and generates a sorted report of pages visited.

*Features*
- Normalizes URLs to ensure consistent tracking.
- Extracts URLs from HTML <a> tags, converting relative URLs to absolute.
- Fetches HTML content with HTTP requests, validating content type and status.
- Crawls recursively with configurable concurrency and page limits.
- Produces a sorted report of pages by link count and URL.

*Installation*
- Clone the repository:
[bash]
git clone https://github.com/Abdullahmojumder/WebCrawler.git
cd WebCrawler

- Install dependencies:
[bash]
go mod tidy

- Build and run:
[bash]
go build -o crawler
./crawler https://example.com 3 10

*Requirements*
- Go 1.21 or later
- Internet connection for crawling

*License*
MIT License


