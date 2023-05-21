# site-accessibility
Checks if certain site is reachable and gets its time to connect in milliseconds.

This project was written using Golang and [Gin web framework](https://github.com/gin-gonic/gin).

Inspired by [NestJs](https://github.com/nestjs/nest).

## How to run the code(Windows)
* Download the source code

    `$ git clone https://github.com/svturkin/site-accessibility.git`

* Go to dir project `site-accessibility`

    `$ cd .\site-accessibility\`

* Run the project

    `$ go run .`

* Open `http://localhost:8080`

### API Docs

- GET /url/info?url=google.com
  - URL params: url
  - returns data in milliseconds(if 0 - site is down or unreachable)

- GET /url/fastest
  - returns the fastest reachable site name that is available

- GET /url/slowest
  - returns the slowest reachable site name that is available

- GET /admin/counter
  - returns JSON object that provides requests statistics for every available route