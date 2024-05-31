generate=bash generate.sh

install:
	@npm config set legacy-peer-deps true

	@npm install google-protobuf
	@npm install grpc-web

	@npm install svelte @sveltestrap/sveltestrap

	@npm install

run:
	@rm proto/*.js 2>/dev/null || true
	@${generate}
	@HOST=0.0.0.0 npm run dev

clean:
	@rm -rf node_modules public/build package-lock.json proto/*.js 2>/dev/null || true
