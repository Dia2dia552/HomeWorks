package game

import (
	"bufio"
	"fmt"
	"os"
)

func StartScene() {

	scene10 := Scene{
		Description: "Ви просто згаяли час, без їжі вам було не вижити. Гра завершилася. Ви програли.",
		Choices:     nil,
	}
	scene9 := Scene{
		Description: "Ви намагаєтеся відкрити сейф, але коли він відчиняється, вас кусає велика комаха і ви непритомнієте. Гра завершилася. Ви програли.",
		Choices:     nil,
	}
	scene8 := Scene{
		Description: "Ви відпочиваєте в наметі і чуєте дзвінок вашого телефону. Ви розбираєтеся і отримуєте допомогу. Гра завершилася. Ви виграли!",
		Choices:     nil,
	}
	scene7 := Scene{
		Description: "Ви намагаєтеся знайти їжу в таборі, але нічого не знаходите." +
			" Ви втомились і вирішили відпочити. Що ви будете робити в наметі?",
		Choices: []Choice{
			{Text: "Спробувати відкрити сейф.", Next: &scene9},
			{Text: "Просто відпочивати.", Next: &scene10},
		},
	}
	scene6 := Scene{
		Description: "Ви досліджуєте тіло тварини і виявляєте щось корисне. Гра завершилася. Ви виграли!",
		Choices:     nil,
	}
	scene5 := Scene{
		Description: "Ви продовжуєте йти через ліс. Через деякий час ви доходите до безлюдного табору. Ви втомились і вирішили відпочити. Що ви будете робити в таборі?",
		Choices: []Choice{
			{Text: "Спробувати знайти їжу.", Next: &scene7},
			{Text: "Відпочити в наметі.", Next: &scene8},
		},
	}

	scene4 := Scene{
		Description: "Продовжуючи йти вглиб печери, ви нарешті доходите до виходу на світло. Гра завершилася. Ви виграли!",
		Choices:     nil,
	}

	scene3 := Scene{
		Description: "Ви вийшли з печери і опинилися в лісі. Через деякий час ви натрапляєте на мертве тіло дивної тварини. Що ви робите?",
		Choices: []Choice{
			{Text: "Нічого з цим не робити і продовжувати йти.", Next: &scene5},
			{Text: "Дослідити тіло тварини.", Next: &scene6},
		},
	}
	scene2 := Scene{
		Description: "Ви рухаєтеся вглиб печери, але там зовсім темно. Навіть ліхтарик не допомагає. Що ви робите?",
		Choices: []Choice{
			{Text: "Продовжувати йти вперед.", Next: &scene4},
		},
	}
	scene1 := Scene{
		Description: "Ви прокинулися біля входу в печеру. Ви лише пам'ятаєте своє ім'я - Стівен. " +
			"Поруч з вами рюкзак, в якому ви знаходите сірники, ліхтарик і ніж. У печері темно. Що ви будете робити?",
		Choices: []Choice{
			{Text: "Піти вглиб печери.", Next: &scene2},
			{Text: "Піти в ліс.", Next: &scene3},
		},
	}

	scene2.Choices = append(scene2.Choices, Choice{Text: "Повернутися назад.", Next: &scene1})

	currentScene := &scene1
	for {
		fmt.Println(currentScene.Description)
		if len(currentScene.Choices) == 0 {
			break
		}

		fmt.Println("Ваші варіанти:")
		for i, choice := range currentScene.Choices {
			fmt.Printf("%d. %s\n", i+1, choice.Text)
		}

		fmt.Print("Виберіть номер варіанту: ")
		reader := bufio.NewReader(os.Stdin)
		choiceStr, _ := reader.ReadString('\n')

		var choiceIndex int
		_, err := fmt.Sscan(choiceStr, &choiceIndex)
		if err != nil || choiceIndex < 1 || choiceIndex > len(currentScene.Choices) {
			fmt.Println("Будь ласка, виберіть правильний номер варіанту.")
		} else {
			currentScene = currentScene.Choices[choiceIndex-1].Next
		}
	}
}
