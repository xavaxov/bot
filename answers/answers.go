//Package answers - пакет для выбора вариантов ответа
package answers

// Answer - Функция для выбора варианта ответов
func Answer(a string) string {
	switch a {
	case "❤️":
		return "Love you too, "
	default:
		return "Send me love ❤️, "
	}
}
