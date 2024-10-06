package main

import (
	"testing"
	"time"
)

// Helper function to create a product for testing
func createProduct(sku string, date string, slot int) *Product {
	expDate, _ := time.Parse("2006-01-02", date)
	return &Product{
		SKU:     sku,
		ExpDate: expDate,
		Slot:    slot,
	}
}

func TestNewWarehouse(t *testing.T) {
	warehouse := NewWarehouse(5)
	if warehouse.Capacity != 5 {
		t.Errorf("expected capacity 5, got %d", warehouse.Capacity)
	}
	if len(warehouse.Slots) != 5 {
		t.Errorf("expected 5 slots, got %d", len(warehouse.Slots))
	}
}

func TestRackIn(t *testing.T) {
	warehouse := NewWarehouse(3)
	response := warehouse.RackIn("SKU123", "2024-12-31")
	if response != "Allocated slot number: 1" {
		t.Errorf("unexpected response: %s", response)
	}
}

func TestRackIn_Full(t *testing.T) {
	warehouse := NewWarehouse(1)
	warehouse.RackIn("SKU123", "2024-12-31")
	response := warehouse.RackIn("SKU456", "2024-12-31")
	if response != "Sorry, rack is full" {
		t.Errorf("expected 'Sorry, rack is full', got %s", response)
	}
}

func TestRackOut(t *testing.T) {
	warehouse := NewWarehouse(3)
	warehouse.RackIn("SKU123", "2024-12-31")
	response := warehouse.RackOut(1)
	if response != "Slot number 1 is free" {
		t.Errorf("unexpected response: %s", response)
	}
}

func TestRackOut_InvalidSlot(t *testing.T) {
	warehouse := NewWarehouse(3)
	response := warehouse.RackOut(1)
	if response != "Invalid slot" {
		t.Errorf("expected 'Invalid slot', got %s", response)
	}
}

func TestSKUForExpDate(t *testing.T) {
	warehouse := NewWarehouse(3)
	warehouse.Slots[0] = createProduct("SKU123", "2024-12-31", 1)
	warehouse.Slots[1] = createProduct("SKU456", "2024-12-31", 2)

	response := warehouse.SKUForExpDate("2024-12-31")
	if response != "SKU123, SKU456" {
		t.Errorf("unexpected response: %s", response)
	}
}

func TestSlotForExpDate(t *testing.T) {
	warehouse := NewWarehouse(3)
	warehouse.Slots[0] = createProduct("SKU123", "2024-12-31", 1)
	warehouse.Slots[1] = createProduct("SKU456", "2024-12-31", 2)

	response := warehouse.SlotForExpDate("2024-12-31")
	if response != "1, 2" {
		t.Errorf("unexpected response: %s", response)
	}
}

func TestSlotForSKU(t *testing.T) {
	warehouse := NewWarehouse(3)
	warehouse.Slots[0] = createProduct("SKU123", "2024-12-31", 1)

	response := warehouse.SlotForSKU("SKU123")
	if response != "1" {
		t.Errorf("unexpected response: %s", response)
	}
}

func TestSlotForSKU_NotFound(t *testing.T) {
	warehouse := NewWarehouse(3)

	response := warehouse.SlotForSKU("SKU999")
	if response != "Not found" {
		t.Errorf("expected 'Not found', got %s", response)
	}
}
