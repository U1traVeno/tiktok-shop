package model

// AllModels 用于统一管理所有模型实例
type AllModels struct {
	Models []interface{}
}

// NewAllModels 初始化所有模型实例
func NewAllModels() *AllModels {
	return &AllModels{
		Models: []interface{}{
			&Auth{},
			&User{},
			&Product{},
			&Cart{},
			&CartItems{},
			&Order{},
		},
	}
}
