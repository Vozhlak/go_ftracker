package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	K1 = 0.035
	K2 = 0.029
)

var (
	Format     = "20060102 15:04:05" // формат даты и времени
	StepLength = 0.65                // длина шага в метрах
	Weight     = 75.0                // вес кг
	Height     = 1.75                // рост м
	Speed      = 1.39                // скорость м/с
)

// parsePackage разбирает входящий пакет в параметре data.
// Возвращаемые значения:
// t — дата и время, указанные в пакете
// steps — количество шагов
// ok — true, если время и шаги указаны корректно, и false — в противном случае
func parsePackage(data string) (t time.Time, steps int, ok bool) {
	// 1. Разделите строку на две части по запятой в слайс ds
	ds := strings.Split(data, ",")

	// 2. Проверьте, чтобы ds состоял из двух элементов
	if len(ds) != 2 {
		return
	}
	var err error
	// получаем время time.Time
	t, err = time.Parse(Format, ds[0])
	if err != nil {
		return
	}
	// получаем количество шагов
	steps, err = strconv.Atoi(ds[1])
	if err != nil || steps < 0 {
		return
	}
	// отмечаем, что данные успешно разобраны
	ok = true
	return
}

// stepsDay перебирает все записи слайса, подсчитывает и возвращает
// общее количество шагов
func stepsDay(storage []string) int {
	// тема оптимизации не затрагивается, поэтому можно
	// использовать parsePackage для каждого элемента списка
	allSteps := 0
	for _, el := range storage {
		_, steps, _ := parsePackage(el)
		allSteps += steps
	}
	return allSteps
}

// calories возвращает количество килокалорий, которые потрачены на
// прохождение указанной дистанции (в метрах) со скоростью 5 км/ч
// Энергозатраты (ккал/мин) = 0,035 * m + (v*v/h) * 0,029 * m
func calories(distance float64) float64 {
	minutes := (distance / Speed) / 60
	return (K1*Weight + (Speed*Speed/Height)*K2*Weight) * minutes
}

// achievement возвращает мотивирующее сообщение в зависимости от
// пройденного расстояния в километрах
func achievement(distance float64) string {
	if distance >= 6.5 {
		return "Отличный результат! Цель достигнута."
	} else if distance >= 3.9 {
		return "Неплохо! День был продуктивный."
	} else if distance >= 2 {
		return "Завтра наверстаем!"
	} else {
		return "Лежать тоже полезно. Главное — участие, а не победа!"
	}
}

// showMessage выводит строку и добавляет два переноса строк
func showMessage(s string) {
	fmt.Printf("%s\n\n", s)
}

// AcceptPackage обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func AcceptPackage(data string, storage []string) []string {
	// 1. Используйте parsePackage для разбора пакета
	t, steps, ok := parsePackage(data)

	//    выведите сообщение в случае ошибки
	if !ok {
		showMessage("ошибочный формат пакета")
		return storage
	}
	//    также проверьте количество шагов на равенство нулю
	if steps == 0 {
		return storage
	}

	// 2. Получите текущее UTC-время и сравните дни
	//    выведите сообщение, если день в пакете t.Day() не совпадает
	//    с текущим днём
	// Получаем текущее UTC-время
	now := time.Now().UTC()

	if now.Day() != t.Day() {
		showMessage("неверный день")
		return storage
	}

	// выводим ошибку, если время в пакете больше текущего времени
	if t.After(now) {
		showMessage(`некорректное значение времени`)
		return storage
	}
	// проверки для непустого storage
	if len(storage) > 0 {
		// 3. Достаточно сравнить первые len(Format) символов пакета с
		//    len(Format) символами последней записи storage
		//    если меньше или равно, то ошибка — некорректное значение времени
		if data[:len(Format)] <= storage[len(storage)-1][:len(Format)] {
			showMessage("некорректное значение времени")
			return storage
		}

		// смотрим, наступили ли новые сутки: YYYYMMDD — 8 символов
		if data[:8] != storage[len(storage)-1][:8] {
			// если наступили,
			// то обнуляем слайс с накопленными данными
			storage = storage[:0]
		}
	}
	// остаётся совсем немного
	// 5. Добавить пакет в storage
	// 6. Получить общее количество шагов
	// 7. Вычислить общее расстояние (в метрах)
	// 8. Получить потраченные килокалории
	// 9. Получить мотивирующий текст
	// 10. Сформировать и вывести полный текст сообщения
	// 11. Вернуть storage
	storage = append(storage, data)
	allSteps := stepsDay(storage)
	allDistance := float64(allSteps) * StepLength
	energy := calories(allDistance)
	achiev := achievement(allDistance / 1000)

	msg := fmt.Sprintf(`Время: %s.
Количество шагов за сегодня: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
%s`, t.Format("15:04:05"), allSteps, allDistance/1000, energy, achiev)
	showMessage(msg)
	return storage
}

func main() {
	// Вы можете сразу проверить работу функции AcceptPackage
	// на небольшом тесте.
	// Если запустить программу после 05:00 UTC, то последнее
	// сообщение должно быть таким:
	// Время: 04:45:21.
	// Количество шагов за сегодня: 16956.
	// Дистанция составила 11.02 км.
	// Вы сожгли 664.23 ккал.
	// Отличный результат! Цель достигнута.

	now := time.Now().UTC()
	today := now.Format("20060102")

	// данные для самопроверки
	input := []string{
		"01:41:03,-100",
		",3456",
		"12:40:00, 3456 ",
		"something is wrong",
		"02:11:34,678",
		"02:11:34,792",
		"17:01:30,1078",
		"03:25:59,7830",
		"04:00:46,5325",
		"04:45:21,3123",
	}

	var storage []string
	storage = AcceptPackage("20230720 00:11:33,100", storage)
	for _, v := range input {
		storage = AcceptPackage(today+" "+v, storage)
	}
}
