import { defineConfig, servers } from "arri";

export default defineConfig({
  server: servers.goServer(),
  // register client generators here
  generators: [],
});
