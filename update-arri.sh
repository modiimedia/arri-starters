(cd app-starter-go ; echo "\nUpdating \"app-starter-go\"" ; pnpm arri use latest ; pnpm i ; GONOSUMDB=github.com/modiimedia/arri go get ; go mod tidy )
(cd app-starter-ts ; echo "\nUpdating \"app-starter-ts\"" ; pnpm arri use latest; pnpm i)
(cd app-starter-ts-with-eslint ; echo "\nUpdating \"app-starter-ts-with-eslint\"" ; pnpm arri use latest; pnpm i;)
(cd generator-starter ; echo "\nUpdating \"generator-starter\"" ; pnpm i @arrirpc/codegen-utils@latest)