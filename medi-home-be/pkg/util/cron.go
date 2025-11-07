package util

import (
	"log"
	"medi-home-be/internal/app/repository"
	"time"

	"github.com/robfig/cron/v3"
)

func StartVoucherCron(voucherRepo repository.VoucherRepository) {
	// dùng timezone Việt Nam
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	c := cron.New(cron.WithLocation(loc))
	c.AddFunc("0 0 * * *", func() {

		if err := voucherRepo.UpdateAllStatus(); err != nil {
			log.Println("Cron lỗi:", err)
		} else {
			log.Println("Cron: cập nhật voucher thành công.")
		}
	})
	c.Start()
}
