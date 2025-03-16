import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

// https://vite.dev/config/
export default defineConfig({
	plugins: [react()],
	build: {
		outDir: "web/",
		rollupOptions: {
			output: {
				manualChunks: {
					react: ["react", "react-dom"],
					vendor: ["react-markdown"]
				}
			}
		}
	}
});
