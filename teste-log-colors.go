package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main2() {
	// Crie um novo logger
	logger := logrus.New()

	// Defina o formato de saída para texto
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true, // Forçar o uso de cores mesmo se não estiver em um terminal
	})

	// Defina a saída para o console
	logger.SetOutput(os.Stdout)

	// Exemplo de mensagens de log com diferentes níveis e cores
	logger.Info("Mensagem de informação (verde)")
	logger.Warn("Mensagem de aviso (amarelo)")
	logger.Error("Mensagem de erro (vermelho)")

	// Você também pode personalizar as cores dos níveis de log
	logrus.AddHook(NewCustomColorsHook())

	logger.Info("Mensagem de informação personalizada")
	logger.Warn("Mensagem de aviso personalizada")
	logger.Error("Mensagem de erro personalizada")
}

// Defina cores personalizadas para os níveis de log
type CustomColorsHook struct{}

func NewCustomColorsHook() *CustomColorsHook {
	return &CustomColorsHook{}
}

func (hook *CustomColorsHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel}
}

func (hook *CustomColorsHook) Fire(entry *logrus.Entry) error {
	switch entry.Level {
	case logrus.InfoLevel:
		entry.Message = "\033[32m" + entry.Message + "\033[0m" // Verde
	case logrus.WarnLevel:
		entry.Message = "\033[33m" + entry.Message + "\033[0m" // Amarelo
	case logrus.ErrorLevel:
		entry.Message = "\033[31m" + entry.Message + "\033[0m" // Vermelho
	}
	return nil
}
