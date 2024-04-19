import "arri";

declare module "arri" {
    // add fields here to extends the H3EventContext in routes, middlewares, and procedures
    interface ArriEventContext {}

    // add fields here to extend the raw H3Event object
    interface H3Event {}
}
