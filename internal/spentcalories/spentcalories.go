package spentcalories

import (
	"fmt"
	"errors"
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
	ErrConvToInt = errors.New("ошибка преобразования в целое число")
	ErrWrongInfo = errors.New("некорректные данные")
	ErrWrongTrain = errors.New("неизвестный тип тренировки")
)

func parseTraining(data string) (int, string, time.Duration, error) {

	threeSlice := strings.Split(data, ",")

	if len(threeSlice) != 3 {
		return 0, "", 0, ErrWrongInfo
	}

	steps, err := strconv.Atoi(threeSlice[0])
	if err != nil {
		return 0, "", 0, ErrConvToInt
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
	
	if duration <= 0 {
		return 0, ErrWrongInfo
	}

	speed := meanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	return weight * speed * minutes / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	
	if duration <= 0 {
		return 0, ErrWrongInfo
	}

	speed := meanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	return weight * speed * minutes / minInH * walkingCaloriesCoefficient, nil

}

func TrainingInfo(data string, weight, height float64) (string, error) {
	
	steps, tip, duration, err := parseTraining(data)

	durationString := fmt.Sprintf("%.2f", duration.Hours())

	if err != nil {
		return "", err
	}

	switch(tip){
	case "Бег":
		distance := didistance(steps, height)
		speed := meanmeanSpeed(steps, height, duration)
		calories, err1 := RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		distance := didistance(steps, height)
		speed := meanmeanSpeed(steps, height, duration)
		calories, err1 := WalkingSpentCalories(steps, weight, height, duration)
	default:
		ruturn "", ErrWrongTrain
	}
	
	if err1 != nil {
		return "", err1
	}

	workout := fmt.Sprintf(
		"Тип тренировки: %s
		Длительность: %s ч.
		Дистанция: %.2f км.
		Скорость: %.2f км/ч
		Сожгли калорий: %.2f",
		tip,
		durationString,
		distance,
		speed,
		calories
	)

	return workout

}