// This library contains other useful functions to assist in creating a generator
// See https://github.com/modiimedia/arri/tree/master/packages/arri-codegen/utils for details
import { defineClientGeneratorPlugin } from "arri-codegen-utils";

interface Options {
    // add or remove fields here for custom generator options
    // official client generators typically have "clientName" and "outputFile" in their options
}

export default defineClientGeneratorPlugin((options: Options) => {
    return {
        options,
        generator: async (def) => {
            // generating something using the app definition and the specified options
        },
    };
});
