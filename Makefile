prepare:
	npx tailwindcss -c tailwind.config.js -o ./static/output.css && cd svelte-components && yarn build

run:
	air

tailwind:
	npx tailwindcss -c tailwind.config.js -o ./static/output.css --watch

	#npx tailwindcss -i ./src/input.css -o ./static/output.css --watch