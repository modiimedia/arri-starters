// This library contains other useful functions to assist in creating a generator
// See https://github.com/modiimedia/arri/tree/master/packages/arri-codegen/utils for details
import { defineClientGeneratorPlugin } from "arri-codegen-utils";

// add or remove fields here for custom generator options
// official client generators have "clientName" and "outputFile" in their options
interface Options {
    clientName: string;
    outputFile: string;
}

export default defineClientGeneratorPlugin((options: Options) => {
    return {
        options,
        generator: async (def) => {
            // generating something using the app definition and the specified options
        },
    };
});
