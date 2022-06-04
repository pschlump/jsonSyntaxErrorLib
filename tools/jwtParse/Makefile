
all:
	go build

test: all
	jwtParse --token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3Rva2VuIjoiZDc5ZTg0ZDMtZjcwYy00YjAyLWE2NGQtNTgyNGY0MDY5NWIyIn0.pPxwn5-a3n4F3ObSctp1tlAZOUiiIdbCwNQcLXAABiA | tail -1 >,a
	diff ,a ref/output.1
	@echo 
	@echo "PASS"



