(cd app-starter-go ; echo "testing app-starter-go" ; pnpm arri build ;)
(cd app-starter-ts ; echo "testing app-starter-ts" ; pnpm arri build ; pnpm typecheck ;)
(cd app-starter-ts-with-eslint ; echo "testing app-starter-ts-with-eslint" ; pnpm arri build ; pnpm lint ; pnpm typecheck ; )
(cd generator-starter ; echo "testing generator-starter" ; pnpm build ; pnpm test ; pnpm typecheck ;)