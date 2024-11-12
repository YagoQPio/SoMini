package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

// Definição de um bloco de memória
type MemoryBlock struct {
	Address int
	Size    int
}

type MemoryManager struct {
	TotalMemory   int
	FreeBlocks    []MemoryBlock
	Allocated     map[int]*Process // Mapa de endereço para processos
	mutex         sync.Mutex
	AllocationAlg string // "best", "worst", "first"
}

// Inicializa o gerenciador de memória
func NewMemoryManager(totalSize int, alg string) *MemoryManager {
	return &MemoryManager{
		TotalMemory:   totalSize,
		FreeBlocks:    []MemoryBlock{{Address: 0, Size: totalSize}},
		Allocated:     make(map[int]*Process),
		AllocationAlg: alg,
	}
}

var memoryManager = NewMemoryManager(1024, "best") // Você pode escolher "best", "worst" ou "first"

type Process struct {
	PID     int
	Command string
	Memory  int
}

func (mm *MemoryManager) allocateMemory(size int) (int, error) {
	mm.mutex.Lock()
	defer mm.mutex.Unlock()

	var selectedBlock *MemoryBlock
	selectedIndex := -1

	switch mm.AllocationAlg {
	case "first":
		for i, block := range mm.FreeBlocks {
			if block.Size >= size {
				selectedBlock = &block
				selectedIndex = i
				break
			}
		}
	case "best":
		minDiff := mm.TotalMemory + 1
		for i, block := range mm.FreeBlocks {
			if block.Size >= size && (block.Size-size) < minDiff {
				selectedBlock = &block
				selectedIndex = i
				minDiff = block.Size - size
			}
		}
	case "worst":
		maxDiff := -1
		for i, block := range mm.FreeBlocks {
			if block.Size >= size && (block.Size-size) > maxDiff {
				selectedBlock = &block
				selectedIndex = i
				maxDiff = block.Size - size
			}
		}
	default:
		return -1, fmt.Errorf("algoritmo de alocação desconhecido")
	}

	if selectedBlock == nil {
		return -1, fmt.Errorf("memória insuficiente")
	}

	address := selectedBlock.Address
	mm.Allocated[address] = &Process{PID: rand.Int(), Memory: size}

	// Atualiza os blocos livres
	if selectedBlock.Size == size {
		// Remove o bloco completamente
		mm.FreeBlocks = append(mm.FreeBlocks[:selectedIndex], mm.FreeBlocks[selectedIndex+1:]...)
	} else {
		// Ajusta o bloco livre
		mm.FreeBlocks[selectedIndex].Address += size
		mm.FreeBlocks[selectedIndex].Size -= size
	}

	return address, nil
}

func (mm *MemoryManager) freeMemory(address int) error {
	mm.mutex.Lock()
	defer mm.mutex.Unlock()

	process, exists := mm.Allocated[address]
	if !exists {
		return fmt.Errorf("endereço de memória não encontrado")
	}

	// Adiciona o bloco de memória de volta aos blocos livres
	mm.FreeBlocks = append(mm.FreeBlocks, MemoryBlock{Address: address, Size: process.Memory})
	delete(mm.Allocated, address)

	// Ordena os blocos livres por endereço para facilitar a fusão
	sort.Slice(mm.FreeBlocks, func(i, j int) bool {
		return mm.FreeBlocks[i].Address < mm.FreeBlocks[j].Address
	})

	// Fusão de blocos adjacentes
	merged := []MemoryBlock{}
	for _, block := range mm.FreeBlocks {
		if len(merged) == 0 {
			merged = append(merged, block)
		} else {
			last := &merged[len(merged)-1]
			if last.Address+last.Size == block.Address {
				last.Size += block.Size
			} else {
				merged = append(merged, block)
			}
		}
	}
	mm.FreeBlocks = merged

	return nil
}

// Função para criar um processo
func createProcess(command string, memorySize int) *Process {
	address, err := memoryManager.allocateMemory(memorySize)
	if err != nil {
		fmt.Println("Erro ao alocar memória:", err)
		return nil
	}
	pid := rand.Int()
	process := &Process{PID: pid, Command: command, Memory: memorySize}
	fmt.Printf("Processo criado: PID=%d, Comando='%s', Memória=%d bytes, Endereço=%d\n", pid, command, memorySize, address)
	return process
}

// Função para encerrar um processo
func endProcess(process *Process) {
	for address, p := range memoryManager.Allocated {
		if p.PID == process.PID {
			err := memoryManager.freeMemory(address)
			if err != nil {
				fmt.Println("Erro ao liberar memória:", err)
			} else {
				fmt.Printf("Memória liberada do processo PID=%d\n", process.PID)
			}
			delete(memoryManager.Allocated, address)
			fmt.Printf("Processo PID=%d encerrado\n", process.PID)
			return
		}
	}
	fmt.Println("Processo não encontrado.")
}
