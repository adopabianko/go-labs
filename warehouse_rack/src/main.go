package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	SKU     string
	ExpDate time.Time
	Slot    int
}

type Warehouse struct {
	Capacity int
	Slots    []*Product
}

func NewWarehouse(capacity int) *Warehouse {
	return &Warehouse{
		Capacity: capacity,
		Slots:    make([]*Product, capacity),
	}
}

func main() {
	w := Warehouse{}

	if len(os.Args) > 1 {
		// Run with input file
		fileName := os.Args[1]

		w.ReadFile(fileName)
	} else {
		// Run with command
		w.ReadCommand()
	}
}

func (w *Warehouse) ReadFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	r := bufio.NewScanner(file)
	r.Split(bufio.ScanLines)

	wh := &Warehouse{}

	for r.Scan() {
		readLineText := strings.Split(r.Text(), " ")

		inputText := readLineText[0]

		wh = processInput(inputText, readLineText, wh)
	}
}

func processInput(inputText string, readLineText []string, w *Warehouse) *Warehouse {
	switch expression := inputText; expression {
	case "create_rack", "create_warehouse_rack":
		capacity := atoi(readLineText[1])
		w = NewWarehouse(capacity)
		fmt.Println("Created a warehouse rack with", capacity, "slots")
	case "rack":
		if w != nil {
			fmt.Println(w.RackIn(readLineText[1], readLineText[2]))
		}
	case "rack_out":
		if w != nil {
			slot := atoi(readLineText[1])
			fmt.Println(w.RackOut(slot))
		}
	case "status":
		if w != nil {
			fmt.Print(w.Status())
		}

	case "sku_numbers_for_product_with_exp_date":
		if w != nil {
			fmt.Println(w.SKUForExpDate(readLineText[1]))
		}
	case "slot_numbers_for_product_with_exp_date":
		if w != nil {
			fmt.Println(w.SlotForExpDate(readLineText[1]))
		}
	case "slot_number_for_sku_number":
		if w != nil {
			fmt.Println(w.SlotForSKU(readLineText[1]))
		}
	case "exit":
		os.Exit(1)
	default:
		fmt.Println("Unknown command")
	}

	return w
}

func (w *Warehouse) ReadCommand() {
	wh := Warehouse{}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		readLineText := strings.Split(strings.TrimSuffix(command, "\n"), " ")
		inputText := readLineText[0]

		wh = *processInput(inputText, readLineText, &wh)
	}
}

func (w *Warehouse) RackIn(sku string, expDate string) string {
	for i := 0; i < w.Capacity; i++ {
		if w.Slots[i] == nil {
			expTime, err := time.Parse("2006-01-02", expDate)
			if err != nil {
				return "Invalid date format"
			}
			w.Slots[i] = &Product{SKU: sku, ExpDate: expTime, Slot: i + 1}
			return fmt.Sprintf("Allocated slot number: %d", i+1)
		}
	}
	return "Sorry, rack is full"
}

func (w *Warehouse) RackOut(slot int) string {
	if slot-1 < 0 || slot > w.Capacity || w.Slots[slot-1] == nil {
		return "Invalid slot"
	}

	w.Slots[slot-1] = nil
	return fmt.Sprintf("Slot number %d is free", slot)
}

func (w *Warehouse) Status() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Slot No.\tSKU No.\tExp Date\n")
	for i, product := range w.Slots {
		if product != nil {
			fmt.Fprintf(&sb, "%d\t%s\t%s\n", i+1, product.SKU, product.ExpDate.Format("2006-01-02"))
		}
	}
	return sb.String()
}

func (w *Warehouse) SKUForExpDate(expDate string) string {
	expTime, err := time.Parse("2006-01-02", expDate)
	if err != nil {
		return "Invalid date format"
	}
	var skus []string
	for _, product := range w.Slots {
		if product != nil && product.ExpDate == expTime {
			skus = append(skus, product.SKU)
		}
	}
	return strings.Join(skus, ", ")
}

func (w *Warehouse) SlotForExpDate(expDate string) string {
	expTime, err := time.Parse("2006-01-02", expDate)
	if err != nil {
		return "Invalid date format"
	}
	var slots []int
	for _, product := range w.Slots {
		if product != nil && product.ExpDate == expTime {
			slots = append(slots, product.Slot)
		}
	}
	sort.Ints(slots)
	strSlots := make([]string, len(slots))
	for i, slot := range slots {
		strSlots[i] = fmt.Sprintf("%d", slot)
	}
	return strings.Join(strSlots, ", ")
}

func (w *Warehouse) SlotForSKU(sku string) string {
	for _, product := range w.Slots {
		if product != nil && product.SKU == sku {
			return fmt.Sprintf("%d", product.Slot)
		}
	}
	return "Not found"
}

// convert string to int
func atoi(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}
