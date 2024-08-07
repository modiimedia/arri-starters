import { defineConfig, servers } from "arri";

export default defineConfig({
    server: servers.tsServer({
        entry: "app.ts",
        port: 3000,
    }),
    generators: [],
});
