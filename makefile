
run:
	templ generate && go build -o ./tmp/main ./cmd && ./tmp/main


tailrun:
	npx tailwindcss -i ./src/input.css -o ./public/output.css --watch