//go:build kafka
// +build kafka

package main

import (
	"testing"
)

// TestGenerateOrder verifies that GenerateOrder creates a valid order.
func TestGenerateOrder(t *testing.T) {
	config := NewProducerConfig()
	producer := NewOrderProducer(config)

	template := OrderTemplate{
		User:     "test-user",
		Item:     "test-item",
		Quantity: 3,
		Price:    10.00,
	}

	order := producer.GenerateOrder(template, 1)

	// Verify basic fields
	if order.OrderID == "" {
		t.Error("Expected OrderID to be set")
	}
	if order.Sequence != 1 {
		t.Errorf("Expected Sequence to be 1, got %d", order.Sequence)
	}
	if order.Status != "pending" {
		t.Errorf("Expected Status to be 'pending', got %s", order.Status)
	}

	// Verify items
	if len(order.Items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(order.Items))
	}
	if order.Items[0].ItemName != "test-item" {
		t.Errorf("Expected ItemName to be 'test-item', got %s", order.Items[0].ItemName)
	}
	if order.Items[0].Quantity != 3 {
		t.Errorf("Expected Quantity to be 3, got %d", order.Items[0].Quantity)
	}

	// Verify financial calculations
	expectedSubTotal := 30.00 // 3 * 10.00
	if order.SubTotal != expectedSubTotal {
		t.Errorf("Expected SubTotal to be %.2f, got %.2f", expectedSubTotal, order.SubTotal)
	}

	expectedTax := expectedSubTotal * config.TaxRate
	if order.Tax != expectedTax {
		t.Errorf("Expected Tax to be %.2f, got %.2f", expectedTax, order.Tax)
	}

	expectedTotal := expectedSubTotal + expectedTax + config.ShippingFee
	if order.Total != expectedTotal {
		t.Errorf("Expected Total to be %.2f, got %.2f", expectedTotal, order.Total)
	}

	// Verify customer info
	if order.CustomerInfo.CustomerID != "test-user" {
		t.Errorf("Expected CustomerID to be 'test-user', got %s", order.CustomerInfo.CustomerID)
	}

	// Verify metadata
	if order.Metadata.EventType != "order.created" {
		t.Errorf("Expected EventType to be 'order.created', got %s", order.Metadata.EventType)
	}
	if order.Metadata.Source != "producer-service" {
		t.Errorf("Expected Source to be 'producer-service', got %s", order.Metadata.Source)
	}
}

// TestNewProducerConfig verifies that the default configuration is correctly created.
func TestNewProducerConfig(t *testing.T) {
	config := NewProducerConfig()

	if config.KafkaBroker != DefaultKafkaBroker {
		t.Errorf("Expected KafkaBroker to be %s, got %s", DefaultKafkaBroker, config.KafkaBroker)
	}
	if config.Topic != DefaultTopic {
		t.Errorf("Expected Topic to be %s, got %s", DefaultTopic, config.Topic)
	}
	if config.TaxRate != ProducerDefaultTaxRate {
		t.Errorf("Expected TaxRate to be %f, got %f", ProducerDefaultTaxRate, config.TaxRate)
	}
	if config.ShippingFee != ProducerDefaultShippingFee {
		t.Errorf("Expected ShippingFee to be %f, got %f", ProducerDefaultShippingFee, config.ShippingFee)
	}
	if config.Currency != ProducerDefaultCurrency {
		t.Errorf("Expected Currency to be %s, got %s", ProducerDefaultCurrency, config.Currency)
	}
}

// TestDefaultOrderTemplates verifies that the default templates are defined.
func TestDefaultOrderTemplates(t *testing.T) {
	if len(DefaultOrderTemplates) == 0 {
		t.Error("Expected DefaultOrderTemplates to have at least one template")
	}

	// Verify all templates have required fields
	for i, template := range DefaultOrderTemplates {
		if template.User == "" {
			t.Errorf("Template %d: Expected User to be set", i)
		}
		if template.Item == "" {
			t.Errorf("Template %d: Expected Item to be set", i)
		}
		if template.Quantity <= 0 {
			t.Errorf("Template %d: Expected Quantity to be positive, got %d", i, template.Quantity)
		}
		if template.Price <= 0 {
			t.Errorf("Template %d: Expected Price to be positive, got %f", i, template.Price)
		}
	}
}

// TestNewOrderProducer verifies that a new OrderProducer is correctly created.
func TestNewOrderProducer(t *testing.T) {
	config := NewProducerConfig()
	producer := NewOrderProducer(config)

	if producer.config != config {
		t.Error("Expected config to be set")
	}
	if producer.sequence != 1 {
		t.Errorf("Expected sequence to start at 1, got %d", producer.sequence)
	}
	if len(producer.templates) == 0 {
		t.Error("Expected templates to be set")
	}
}
