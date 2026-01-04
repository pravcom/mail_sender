package app

import (
	"context"
	"fmt"
	"log"
	"mail_sender/cmd/internal/app/config"
	"time"

	"github.com/go-co-op/gocron"
)

func Run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	timeLocation, err := time.LoadLocation(cfg.App.TimeLocation)
	if err != nil {
		log.Printf("time location error: %v", err)
	}

	s := gocron.NewScheduler(timeLocation)

	// Cron выражение: каждый будний день (пн-пт) в 10:00
	// Формат: секунды минуты часы день_месяца месяц день_недели
	// 0-6: воскресенье-суббота (0 и 7 = воскресенье)
	// 1-5: понедельник-пятница
	//"*/1 * * * *"

	// cronExpr := "00 10 * * 1-5"

	cronExpr := cfg.App.CronExpr
	j, err := s.Cron(cronExpr).Do(func() {
		if err := sendMailGoMail(cfg); err != nil {
			log.Printf("Ошибка отправки письма: %v", err)
		}
	})

	fmt.Printf("Задание настроено: %s\n", cronExpr)
	timeNext := j.NextRun()
	fmt.Printf("Следующий запуск: %s\n",
		timeNext)
	go func() {
		s.StartBlocking()
	}()

	<-ctx.Done()

	s.Stop()

	log.Println("shitting down server gracefully")

	log.Println("Send mail success")

	return nil
}
