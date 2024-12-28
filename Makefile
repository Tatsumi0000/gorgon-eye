dev/setup:
	docker build . -t gorgon-eye 

gofmt: 
	docker run -it --rm -v $(PWD):/app gorgon-eye /bin/sh -c "gofmt -w ."

dev/run:
	docker run -it --rm -v $(PWD):/app gorgon-eye /bin/sh 
