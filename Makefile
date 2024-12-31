dev/setup:
	docker build . -t gorgon-eye 

gofmt: 
	docker run -it -v $(PWD):/app gorgon-eye /bin/sh -c "gofmt -w ."

dev/run:
	docker run -it -v $(PWD):/app gorgon-eye /bin/sh 
