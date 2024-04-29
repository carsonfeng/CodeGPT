package ollama

type Model string

const (
	LLaMA3_70b Model = "llama3:70b"
)

func (m Model) String() string {
	return string(m)
}

func (m Model) GetModel() string {
	return GetModel(m)
}

func (m Model) IsVaild() bool {
	switch m {
	case LLaMA3_70b:
		return true
	default:
		return false
	}
}

var model = map[Model]string{
	LLaMA3_70b: string(LLaMA3_70b),
}

// GetModel returns the model name.
func GetModel(modelName Model) string {
	if _, ok := model[modelName]; !ok {
		return model[LLaMA3_70b]
	}
	return model[modelName]
}
