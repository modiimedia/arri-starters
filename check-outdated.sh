(cd app-starter-go ; echo "\nChecking outdated in \"app-starter-go\"" ; pnpm outdated ; go list -u -m all)
(cd app-starter-ts ; echo "\nChecking outdated in \"app-starter-ts\"" ; pnpm outdated)
(cd app-starter-ts-with-eslint ; echo "\nChecking outdated in \"app-starter-ts-with-eslint\"" ; pnpm outdated)
(cd generator-starter ; echo "\nChecking outdated in \"generator-starter\"" ; pnpm outdated)