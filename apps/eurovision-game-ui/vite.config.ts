import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import viteTsConfigPaths from "vite-tsconfig-paths";

export default defineConfig({
	server: {
		port: 3000,
		host: "localhost",
	},
	plugins: [
		react(),
		viteTsConfigPaths({
			root: "../../",
		}),
	],
});
