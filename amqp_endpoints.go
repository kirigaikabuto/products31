package products31

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
)

type ProductAmqpEndpoints struct {
	productStore ProductStore
}

func NewProductAmqpEndpoints(p ProductStore) ProductAmqpEndpoints {
	return ProductAmqpEndpoints{productStore: p}
}

func (p *ProductAmqpEndpoints) CreateProductAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		prod := &Product{}
		jsonData := message.Body
		err := json.Unmarshal(jsonData, &prod)
		if err != nil {
			panic(err)
		}
		newProduct, err := p.productStore.Create(prod)
		if err != nil {
			panic(err)
		}
		response, err := json.Marshal(newProduct)
		if err != nil {
			panic(err)
		}
		return &amqp.Message{Body: response}
	}
}

func (p *ProductAmqpEndpoints) ListProductAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		products, err := p.productStore.List()
		response, err := json.Marshal(products)
		if err != nil {
			panic(err)
		}
		return &amqp.Message{Body: response}
	}
}
