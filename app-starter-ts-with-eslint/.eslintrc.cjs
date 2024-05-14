module.exports = {
    root: true,
    extends: [
        "eslint:recommended",
        "plugin:@typescript-eslint/recommended",
        "plugin:@arrirpc/recommended",
        "prettier",
    ],
    rules: {
        "no-undef": 0, // ts and @typescript-eslint handles this
        "no-shadow": 0, // ts and @typescript-eslint handles this
    },
};
