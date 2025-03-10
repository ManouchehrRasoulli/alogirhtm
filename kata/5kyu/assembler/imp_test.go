package kata

import "testing"

func TestIncrement(t *testing.T) {
	variables := map[string]int{
		"a": 1,
	}

	incrementer := inc{
		variable: "b",
	}

	pc := 0
	npc := incrementer.execute(variables, pc)

	t.Log(variables, npc)
}

func TestSimpleAssembler(t *testing.T) {
	t.Logf("%+v", SimpleAssembler([]string{"mov a 5", "inc a", "dec a", "dec a", "jnz a -1", "inc a"})) // map[a:1]
	t.Logf("%+v", SimpleAssembler([]string{"mov a -10", "mov b a", "inc a", "dec b", "jnz a -2"}))      // map[a:0 b:-20]
	t.Logf("%+v", SimpleAssembler([]string{
		"mov a 1", "mov b 1",
		"mov c 0", "mov d 26",
		"jnz c 2", "jnz 1 5",
		"mov c 7", "inc d",
		"dec c", "jnz c -2",
		"mov c a", "inc a",
		"dec b", "jnz b -2",
		"mov b c", "dec d",
		"jnz d -6", "mov c 18",
		"mov d 11", "inc a",
		"dec d", "jnz d -2",
		"dec c", "jnz c -5"})) // map[a:318009 b:196418 c:0 d:0]
}

func TestSimpleAssemblerInvalidInstruction(t *testing.T) {
	t.Logf("%+v", SimpleAssembler([]string{
		"mov k 165", "mov h 3",
		"jnz h 2", "jnz k 6",
		"jnz dec k", "jnz inc h",
		"jnz dec k", "jnz dec h",
		"jnz dec h", "jnz dec k",
		"jnz inc h"})) // map[h:3 k:165]
}
