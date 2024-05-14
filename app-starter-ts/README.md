# Arri Server

This project was created by `arri init`. View details at [https://github.com/modiimedia/arri](https://github.com/modiimedia/arri)

## Usage

```bash
arri dev
arri build
```

### Viewing All Procedures

You can view all server rpcs by visiting the `/__definition` endpoint. This can be customized in the `ArriApp` constructor options using the `rpcRoutePrefix` and `rpcDefinitionPath` options.

### Adding Procedures

Create new RPCs in the `./src/procedures` directory by adding `.rpc.ts` files. This behavior can be configured in the `arri.config.ts` using the `procedureDir` and `procedureGlobPatterns` options.

**Server Code:**

```ts
// ./src/procedures/sayHello.rpc.ts
import { defineRpc } from "@arrirpc/server";
import { a } from "@arrirpc/schema";

export default defineRpc({
    params: a.object({
        name: a.string(),
    }),
    response: a.object({
        message: a.string(),
    }),
    handler({ params }) {
        return {
            message: `Hello ${params.name}`,
        };
    },
});
```

```ts
// ./src/procedures/users/createUser.rpc.ts
import { defineRpc } from "@arrirpc/server";
import { a } from "@arrirpc/schema";

export default defineRpc({
    params: a.object({
        name: a.string(),
        email: a.string(),
    }),
    response: a.object({
        id: a.string(),
        name: a.string(),
        email: a.string(),
    }),
    handler({ params }) {
        return {
            id: getRandomId(),
            name: params.name,
            email: params.email,
        };
    },
});
```

**Generated Client Code:**

```ts
await client.sayHello({
    name: "John",
});
await client.users.createUser({
    name: "John Doe",
    email: "johndoe@example.com",
});
```

### Manually Adding Procedures

If you don't want to use the file-based router you can also manual add RPCs like so:

#### Using the `ArriApp` instance

```ts
import { ArriApp } from "@arrirpc/server";
import { a } from "@arrirpc/schema";

const app = new ArriApp();

app.rpc("sayHello", {
    params: a.object({
        name: a.string(),
    }),
    response: a.object({
        message: a.string(),
    }),
    handler({ params }) {
        return {
            message: `Hello ${params.name}`,
        };
    },
});

export default app;
```

#### Using a router

```ts
import { ArriApp, ArriRouter } from "@arrirpc/server";
import { a } from "@arrirpc/schema";

const app = new ArriApp();
const router = new ArriRouter();

router.rpc("sayHello", {
    params: a.object({
        name: a.string(),
    }),
    response: a.object({
        message: a.string(),
    }),
    handler({ params }) {
        return {
            message: `Hello ${params.name}`,
        };
    },
});

app.use(router);

export default app;
```

### Adding Non-RPC Routes

For miscellaneous API endpoints you can do the following:

```ts
// using the app instance
app.route({
    path: "/status",
    method: "get",
    handler(event) {
        return "ok";
    },
});

// using a router instance
router.route({
    path: "/status",
    method: "get",
    handler(event) {
        return "ok";
    },
});
```

These routes will not be listed in the `__definition.json` file and will be ignored by code generators.
