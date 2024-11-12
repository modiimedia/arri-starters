(cd app-starter-go ; echo "\nRunning \"pnpm update\" in \"app-starter-go\"" ; pnpm update --save ; echo "\nRunning \"go get -u\" in \"app-starter-go\"" ; go get -u ; echo "\nRunning \"go mod tidy\" in \"app-starter-go\"" ; go mod tidy)
(cd app-starter-ts ; echo "\nRunning \"pnpm update\" in \"app-starter-ts\"" ; pnpm update --save)
(cd app-starter-ts-with-eslint ; echo "\nRunning \"pnpm update\" in \"app-starter-ts-with-eslint\"" ; pnpm update --save)
(cd generator-starter ; echo "\nRunning \"pnpm update\" in  \"generator-starter\"" ; pnpm update --save)