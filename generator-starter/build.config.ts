import { defineBuildConfig } from "unbuild";

export default defineBuildConfig({
    entries: ["./src/index.ts"],
    declaration: true,
    rollup: {
        emitCJS: true,
        dts: {
            respectExternal: true,
        },
    },
    outDir: "dist",
});
