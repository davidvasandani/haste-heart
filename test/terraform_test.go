package test

import (
	"crypto/tls"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/shell"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerraformQuestion1(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../question_1",
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	output := terraform.Output(t, terraformOptions, "myip")

	cmd := shell.Command{
		Command: "curl",
		Args:    []string{"-s", "http://ipv4.icanhazip.com"},
	}

	out := shell.RunCommandAndGetOutput(t, cmd)
	assert.Equal(t, output, strings.TrimSpace(out))
}

func TestTerraformQuestion2(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../question_2",
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	message_one := terraform.Output(t, terraformOptions, "message_one_input")
	message_two := terraform.Output(t, terraformOptions, "message_two_input")

	assert.Equal(t, "I am message one", message_one)
	assert.Equal(t, "I am message two", message_two)
}

func TestTerraformQuestion3(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../question_3",
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	image_tag := terraform.Output(t, terraformOptions, "image_tag")
	require.Equal(t, 64, len(image_tag))

	// Setup a TLS configuration to submit with the helper, a blank struct is acceptable
	tlsConfig := tls.Config{}

	// It can take a second or so for the container to boot up, so retry a few times
	retries := 15
	sleep := 5 * time.Second

	endpoint := "http://0.0.0.0:8080/haste-heart"

	// Verify that we get back a 200 OK with the expected value
	http_helper.HttpGetWithRetryWithCustomValidation(
		t,
		endpoint,
		&tlsConfig,
		retries,
		sleep,
		func(statusCode int, body string) bool {
			isOk := statusCode == 200
			isRespondingToInput := strings.Contains(body, "haste-heart")
			return isOk && isRespondingToInput
		},
	)

	terraformDirectory := "../question_3"
	testFile := filepath.Join(terraformDirectory, "helloWorld")
	// Check if file exists
	assert.FileExists(t, testFile)

	// remove testFile, if it exists
	if _, err := os.Stat(testFile); err == nil {
		os.Remove(testFile)
	}

	// Verify database tables were created

	shell.RunCommand(t, shell.Command{
		WorkingDir: "../question_3/",
		Command:    "./check.sh",
		Args:       []string{"hasteheart"},
	})

	shell.RunCommand(t, shell.Command{
		WorkingDir: "../question_3/",
		Command:    "./check.sh",
		Args:       []string{"test"},
	})
}
