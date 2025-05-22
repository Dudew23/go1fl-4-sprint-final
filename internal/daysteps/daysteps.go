package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

var (
	ErrConvToInt = errors.New("ошибка преобразования в целое число")
	ErrWrongInfo = errors.New("некорректные данные")
)

func parsePackage(data string) (int, time.Duration, error) {

	dataSlice := strings.Split(data, ",")

	if len(dataSlice) != 2 {
		return 0, 0, ErrWrongInfo
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, 0, ErrConvToInt
	}
	if steps <= 0 {
		return 0, 0, ErrConvToInt
	}

	duration, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return 0, 0, err
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {

	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm

	calories, err := WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}
	
	infoString := fmt.Sprintf(
		"Количество шагов: %d.
		Дистанция составила %.2f км.
		Вы сожгли %.2f ккал.", steps, didistance, cacalories
	)

	return infoString
}
