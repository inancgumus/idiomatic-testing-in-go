# Listing 2.6: Returning the mutated value

## [oop](https://github.com/inancgumus/gobyexample/blob/3329e8423c4af636bc4728f35aff356d1b10dbbe/oop) / [value-receivers](https://github.com/inancgumus/gobyexample/blob/3329e8423c4af636bc4728f35aff356d1b10dbbe/oop/value-receivers) / [main.go](https://github.com/inancgumus/gobyexample/blob/3329e8423c4af636bc4728f35aff356d1b10dbbe/oop/value-receivers/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "fmt"

type usage int

func (u usage) high() bool       { return u >= 95 }
func (u usage) set(to int) usage { return usage(to) }

func main() {
	var cpu usage = 99 // cpu is 99
	cpu = cpu.set(70)  // cpu is 70
	fmt.Printf("cpu: %d%% high:%t\n", cpu, cpu.high())
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -4,11 +4,11 @@ import "fmt"
 
 type usage int
 
-func (u usage) high() bool { return u >= 95 }
-func (u usage) set(to int) { u = usage(to) } //nolint:staticcheck,ineffassign
+func (u usage) high() bool       { return u >= 95 }
+func (u usage) set(to int) usage { return usage(to) }
 
 func main() {
 	var cpu usage = 99 // cpu is 99
-	cpu.set(70)        // cpu is still 99
+	cpu = cpu.set(70)  // cpu is 70
 	fmt.Printf("cpu: %d%% high:%t\n", cpu, cpu.high())
 }
```

