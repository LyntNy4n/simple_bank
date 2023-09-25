package mail

import (
	"simple_bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, privateConfig, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewEmailConfig(config.EmailSenderName, privateConfig.EmailSenderAddress, privateConfig.EmailSenderPassword)

	subject := "Test send email with gmail"
	content := `
	<h1>This is a test email</h1>
	`
	to := []string{"lyntny4n@gmail.com"}
	attachFiles := []string{"../app.env"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
