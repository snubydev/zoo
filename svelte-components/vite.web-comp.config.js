import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import sveltePreprocess from "svelte-preprocess";

export default defineConfig({
	plugins: [
		svelte({
			preprocess: sveltePreprocess({}),
			exclude: /\.comp\.svelte$/,
			emitCss: false,
		}),
		svelte({
			preprocess: sveltePreprocess(),
			include: /\.comp\.svelte$/,
			compilerOptions: {
				customElement: true,
			},
			emitCss: false,
		}),
	],
	build: {
		sourcemap: true,
		target: "modules",
		lib: {
			entry: [
				"src/lib/Details.comp.svelte",
				"src/lib/List.comp.svelte",
			],
		},
	},
});
