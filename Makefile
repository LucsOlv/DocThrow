run:
	@templ generate
	@air

tailwind:
	@bunx tailwindcss -i ./web/css/input.css -o ./assets/css/output.css --watch

templ:
	@templ generate -watch -proxy=http://localhost:3000
