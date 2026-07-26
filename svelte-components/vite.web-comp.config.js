// import { defineConfig } from "vite";
// import { svelte } from "@sveltejs/vite-plugin-svelte";
// import sveltePreprocess from "svelte-preprocess";
//
// export default defineConfig({
// 	plugins: [
// 		svelte({
// 			preprocess: sveltePreprocess({}),
// 			exclude: /\.comp\.svelte$/,
// 			emitCss: false,
// 		}),
// 		svelte({
// 			preprocess: sveltePreprocess(),
// 			include: /\.comp\.svelte$/,
// 			compilerOptions: {
// 				customElement: true,
// 			},
// 			emitCss: false,
// 		}),
// 	],
// 	build: {
// 		sourcemap: true,
// 		target: "modules",
// 		lib: {
// 			entry: [
// 				"src/lib/Details.comp.svelte",
// 				"src/lib/List.comp.svelte",
// 			],
// 		},
// 	},
// });

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
		target: "es2015",
		lib: {
			entry: [
				"src/lib/Details.comp.svelte",
				"src/lib/List.comp.svelte",
			],
		},
		rollupOptions: {
			output: {
				// Force any shared/common code split out from the two entries
				// into a chunk explicitly named "index" instead of letting
				// Rollup auto-derive a name (which produced "disclose-version-*").
				manualChunks(id) {
					if (
						!id.includes("Details.comp.svelte") &&
						!id.includes("List.comp.svelte")
					) {
						return "index";
					}
				},
				chunkFileNames: "[name]-[hash].js",
			},
		},
	},
});
