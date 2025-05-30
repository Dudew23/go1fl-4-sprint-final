package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

var (
	ErrConvToInt     = errors.New("integer conversion error")
	ErrWrongInfo     = errors.New("incorrect data")
	ErrWrongTrain    = errors.New("unknown training type")
	ErrStepsLessZero = errors.New("steps cannot be less than zero")
)

func parseTraining(data string) (int, string, time.Duration, error) {
	threeSlice := strings.Split(data, ",")

	if len(threeSlice) != 3 {
		return 0, "", 0, ErrWrongInfo
	}

	steps, err := strconv.Atoi(threeSlice[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps <= 0 {
		return 0, "", 0, ErrStepsLessZero
	}

	activity := strings.TrimSpace(threeSlice[1])
	if activity != "Бег" && activity != "Ходьба" {
		return 0, "", 0, ErrWrongTrain
	}

	duration, err := time.ParseDuration(threeSlice[2])
	if err != nil {
		return 0, "", 0, err
	}

	return steps, threeSlice[1], duration, nil
}

func distance(steps int, height float64) float64 {
	return height * stepLengthCoefficient * float64(steps) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)

	return dist / duration.Hours()
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, ErrWrongInfo
	}

	speed := meanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	return weight * speed * minutes / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, ErrWrongInfo
	}

	speed := meanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	return weight * speed * minutes / minInH * walkingCaloriesCoefficient, nil

}

func TrainingInfo(data string, weight, height float64) (string, error) {

	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	var info string
	switch activity {
	case "Бег":
		distance := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		calories, _ := RunningSpentCalories(steps, weight, height, duration)
		info = fmt.Sprintf(
			"Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			duration.Hours(), distance, speed, calories,
		)
	case "Ходьба":
		distance := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		calories, _ := WalkingSpentCalories(steps, weight, height, duration)
		info = fmt.Sprintf(
			"Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			duration.Hours(), distance, speed, calories,
		)
	default:
		return "", ErrWrongTrain
	}

	return info, nil
}
