module Chatty-GoGo

go 1.15

replace mygithub.com/Logic => /Logic

replace mygithub.com/Routes => /Routes

require (
	github.com/gin-gonic/gin v1.6.3 // indirect
	mygithub.com/Logic v0.0.0-00010101000000-000000000000
	mygithub.com/Routes v0.0.0-00010101000000-000000000000
)
