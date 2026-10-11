// NOTE: goto jumps forward to shared cleanup. It can never skip a declaration
// nor dive into a block, which keeps it honest. Prefer labeled break or an
// early return; save goto for the rare case where both read worse.
package main

import "fmt"

func main() {
	speed := 12
	for speed < 100 {
		if speed%5 == 0 {
			goto done // both exits below need the same wind-down
		}
		speed = speed*2 + 1
	}
	fmt.Println("loop finished on its own")

done:
	fmt.Println("wind down at speed", speed)

	// BUG: both of these jumps are rejected by the compiler, so they stay
	// commented out. A goto may never skip a declaration:
	//   speed := 1
	//   goto skip
	//   extra := 2  // rejects: jumps over declaration of extra
	// skip:
	//   _ = extra
	// Nor may it dive into a block:
	//   if speed > 0 {
	//       goto inner // rejects: jumps into block
	//   }
	//   if speed < 100 {
	//   inner:
	//       fmt.Println("unreachable legally")
	//   }
}
