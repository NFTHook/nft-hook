package util

import (
	"math/rand"
	"time"
)

func GetRandomRarity() string {
	// 使用时间戳作为随机种子
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	// 稀有度级别列表
	rarityLevels := []string{"Common", "Rare", "Epic", "Legendary"}

	// 稀有度级别比例
	// Legendary: 1%
	// Epic: 9%
	// Rare: 30%
	// Common: 剩下的
	rarityWeights := []int{1, 9, 30, 60}

	// 随机选择稀有度级别
	raritySum := 0
	for _, weight := range rarityWeights {
		raritySum += weight
	}

	randomNumber := randGen.Intn(raritySum)

	for i, weight := range rarityWeights {
		if randomNumber < weight {
			return rarityLevels[i]
		}
		randomNumber -= weight
	}

	return "Common" // 如果没有匹配到其他级别，返回Common
}
