func Deposit(c *fiber.Ctx) error {
	type Request struct {
		NoRekening string `json:"no_rekening"`
		Nominal    int    `json:"nominal"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		config.Log.WithError(err).Error("Failed to parse deposit request")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "Invalid request"})
	}

	var user models.User
	if err := config.DB.Where("no_rekening = ?", req.NoRekening).First(&user).Error; err != nil {
		config.Log.WithFields(logrus.Fields{"no_rekening": req.NoRekening}).Warning("Deposit failed: Account not found")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "Nomor rekening tidak ditemukan"})
	}

	user.Saldo += req.Nominal
	if err := config.DB.Save(&user).Error; err != nil {
		config.Log.WithError(err).Error("Failed to update balance after deposit")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"remark": "Gagal memproses transaksi"})
	}

	config.Log.WithFields(logrus.Fields{
		"no_rekening": req.NoRekening,
		"nominal":     req.Nominal,
		"saldo":       user.Saldo,
	}).Info("Deposit transaction successful")

	return c.JSON(fiber.Map{"saldo": user.Saldo})
}


func Withdraw(c *fiber.Ctx) error {
	type Request struct {
		NoRekening string `json:"no_rekening"`
		Nominal    int    `json:"nominal"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		config.Log.WithError(err).Error("Failed to parse withdrawal request")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "Invalid request"})
	}

	var user models.User
	if err := config.DB.Where("no_rekening = ?", req.NoRekening).First(&user).Error; err != nil {
		config.Log.WithFields(logrus.Fields{"no_rekening": req.NoRekening}).Warning("Withdrawal failed: Account not found")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "Nomor rekening tidak ditemukan"})
	}

	if user.Saldo < req.Nominal {
		config.Log.WithFields(logrus.Fields{
			"no_rekening": req.NoRekening,
			"nominal":     req.Nominal,
			"saldo":       user.Saldo,
		}).Warning("Withdrawal failed: Insufficient balance")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "Saldo tidak cukup"})
	}

	user.Saldo -= req.Nominal
	if err := config.DB.Save(&user).Error; err != nil {
		config.Log.WithError(err).Error("Failed to update balance after withdrawal")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"remark": "Gagal memproses transaksi"})
	}

	config.Log.WithFields(logrus.Fields{
		"no_rekening": req.NoRekening,
		"nominal":     req.Nominal,
		"saldo":       user.Saldo,
	}).Info("Withdrawal transaction successful")

	return c.JSON(fiber.Map{"saldo": user.Saldo})
}

