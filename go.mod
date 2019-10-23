module github.com/yashmurty/wealth-park

go 1.12

require (
	github.com/danielkov/gin-helmet v0.0.0-20171108135313-1387e224435e
	github.com/gin-contrib/gzip v0.0.1
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/pkg/errors v0.8.1
	github.com/rs/xid v1.2.1
	github.com/stretchr/testify v1.4.0
	github.com/yashmurty/m-rec v0.0.0-20190814061928-4964b8bd2c05
	go.uber.org/zap v1.11.0
	gopkg.in/go-playground/validator.v9 v9.29.1
)

// Fix the ambiguous import error.
replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
