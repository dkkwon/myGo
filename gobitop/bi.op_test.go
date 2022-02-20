package gobitop

import "fmt"

func ExampleBitOp() {

	const NR_CELL_IDENTITY_LEN uint32 = 36
	globalCellID := uint32(0)

	// TODO: need to check if gnodebId is bigger than 2^gnodebIdLen
	gnodebId := uint32(22577)
	gnodebIdLen := uint32(22)
	fmt.Printf("decimal: %d\n", gnodebId)
	fmt.Printf("binary : %32b\n", gnodebId)

	globalCellID = gnodebId << (NR_CELL_IDENTITY_LEN - gnodebIdLen)
	fmt.Printf("binary : %32b\n", globalCellID)

	// TODO: need to check if cellID is bigger than 2^(NR_CELL_IDENTITY_LEN - gnodebIdLen)
	cellID := uint32(2)
	globalCellID = globalCellID | cellID
	fmt.Printf("binary : %32b\n", globalCellID)

	// Output:
}
