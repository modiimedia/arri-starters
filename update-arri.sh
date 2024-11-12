(cd app-starter-go ; echo "\nUpdating \"app-starter-go\"" ; pnpm i arri@latest ; go get github.com/modiimedia/arri@latest ; go mod tidy )
(cd app-starter-ts ; echo "\nUpdating \"app-starter-ts\"" ; pnpm i arri@latest @arrirpc/server@latest @arrirpc/schema@latest)
(cd app-starter-ts-with-eslint ; echo "\nUpdating \"app-starter-ts-with-eslint\"" ; pnpm i arri@latest @arrirpc/server@latest @arrirpc/schema@latest ; pnpm i --save-dev @arrirpc/eslint-plugin@latest)
(cd generator-starter ; echo "\nUpdating \"generator-starter\"" ; pnpm i @arrirpc/codegen-utils@latest)