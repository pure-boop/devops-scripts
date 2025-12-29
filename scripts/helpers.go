package devops_scripts

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// GeneratePrivateKey generates a new RSA private key
func GeneratePrivateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

// GenerateCertificate generates a new SSL/TLS certificate
func GenerateCertificate(privateKey *rsa.PrivateKey, commonName string) (*x509.Certificate, error) {
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"DevOps Scripts"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(1, 1, 0),
		DNSNames:  []string{commonName},
	}

	return template.AppendSubjectAltName(commonName).Sign(rand.Reader, privateKey, x509.SHA256WithRSAEncryption)
}

// GetEC2InstanceMetadata retrieves the instance metadata
func GetEC2InstanceMetadata(ec2Session *ec2.EC2) (*ec2.Instance, error) {
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(os.Getenv("EC2_INSTANCE_ID"))},
	}
	output, err := ec2Session.DescribeInstances(input)
	if err != nil {
		return nil, err
	}
	return output.Reservations[0].Instances[0], nil
}

// GetSSMParameter retrieves a parameter from SSM
func GetSSMParameter(ssmSession *ssm.SSM, parameterName string) (*ssm.Parameter, error) {
	input := &ssm.GetParameterInput{
		Name: aws.String(parameterName),
	}
	output, err := ssmSession.GetParameter(input)
	if err != nil {
		return nil, err
	}
	return output.Parameter, nil
}

// GetExecutablePath returns the path to the executable
func GetExecutablePath() string {
	exePath, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}
	return exePath
}

// GetOS returns the current operating system
func GetOS() string {
	return runtime.GOOS
}

// GetSSMInstanceInfo retrieves information about the current instance
func GetSSMInstanceInfo(ssmSession *ssm.SSM) (*ssm.InstanceInformation, error) {
	input := &ssm.GetInstanceInformationInput{
		InstanceIds: []*string{aws.String(os.Getenv("EC2_INSTANCE_ID"))},
	}
	output, err := ssmSession.GetInstanceInformation(input)
	if err != nil {
		return nil, err
	}
	return output.InstanceInformationList[0], nil
}

// GetSSMParameterValue retrieves the value of a parameter from SSM
func GetSSMParameterValue(ssmSession *ssm.SSM, parameterName string) (string, error) {
	parameter, err := GetSSMParameter(ssmSession, parameterName)
	if err != nil {
		return "", err
	}
	return *parameter.Value, nil
}

// GetEC2InstanceType retrieves the instance type of the current EC2 instance
func GetEC2InstanceType(ec2Session *ec2.EC2) (string, error) {
	instance, err := GetEC2InstanceMetadata(ec2Session)
	if err != nil {
		return "", err
	}
	return *instance.InstanceType, nil
}

// GetEC2AvailabilityZone retrieves the availability zone of the current EC2 instance
func GetEC2AvailabilityZone(ec2Session *ec2.EC2) (string, error) {
	instance, err := GetEC2InstanceMetadata(ec2Session)
	if err != nil {
		return "", err
	}
	return *instance.Placement.AvailabilityZone, nil
}

// GetSSMParameterNames retrieves a list of parameter names from SSM
func GetSSMParameterNames(ssmSession *ssm.SSM) ([]string, error) {
	input := &ssm.GetParametersByPathInput{
		Path: aws.String("/"),
	}
	output, err := ssmSession.GetParametersByPath(input)
	if err != nil {
		return nil, err
	}
	var parameterNames []string
	for _, parameter := range output.Parameters {
		parameterNames = append(parameterNames, *parameter.Name)
	}
	return parameterNames, nil
}

// GetJSONValue retrieves a JSON value from a string
func GetJSONValue(jsonString string, key string) (string, error) {
	var jsonData map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonData)
	if err != nil {
		return "", err
	}
	value := jsonData[key]
	if value == nil {
		return "", fmt.Errorf("key %s not found in JSON", key)
	}
	return fmt.Sprintf("%v", value), nil
}

// GetRandomUUID generates a random UUID
func GetRandomUUID() (string, error) {
	uuid := v4()
	return uuid, nil
}

func v4() string {
 bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])
}

// GetRandomFloat generates a random float between a range
func GetRandomFloat(min, max float64) (float64, error) {
	return math.Floor((max-min+1)*math.Random()*10)/10 + min, nil
}

// GetOSArchitecture returns the architecture of the current operating system
func GetOSArchitecture() string {
	return runtime.GOARCH
}

// GetExecutableName returns the name of the executable
func GetExecutableName() string {
	return filepath.Base(os.Args[0])
}

// GetExecutableDir returns the directory of the executable
func GetExecutableDir() string {
	return filepath.Dir(os.Args[0])
}

// GetEnvironmentVariable retrieves an environment variable
func GetEnvironmentVariable(key string) string {
	return os.Getenv(key)
}

// GetHostname retrieves the hostname of the current instance
func GetHostname() string {
	return os.Getenv("HOSTNAME")
}

// GetSSMParameterNamesRecursive retrieves a list of parameter names from SSM, recursively
func GetSSMParameterNamesRecursive(ssmSession *ssm.SSM, path string) ([]string, error) {
	input := &ssm.GetParametersByPathInput{
		Path: aws.String(path),
	}
	output, err := ssmSession.GetParametersByPath(input)
	if err != nil {
		return nil, err
	}
	var parameterNames []string
	for _, parameter := range output.Parameters {
		parameterNames = append(parameterNames, *parameter.Name)
	}
	for _, parameter := range output.Parameters {
		if strings.HasSuffix(*parameter.Name, "/") {
			subParameters, err := GetSSMParameterNamesRecursive(ssmSession, *parameter.Name)
			if err != nil {
				return nil, err
			}
			parameterNames = append(parameterNames, subParameters...)
		}
	}
	return parameterNames, nil
}

// GetAWSRegion retrieves the AWS region of the current instance
func GetAWSRegion() string {
	return os.Getenv("AWS_REGION")
}

// GetAWSAccountID retrieves the AWS account ID of the current instance
func GetAWSAccountID() string {
	return os.Getenv("AWS_ACCOUNT_ID")
}

// GetEC2InstanceID retrieves the EC2 instance ID of the current instance
func GetEC2InstanceID() string {
	return os.Getenv("EC2_INSTANCE_ID")
}

// GetSSMInstanceProfile retrieves the SSM instance profile of the current instance
func GetSSMInstanceProfile() string {
	return os.Getenv("SSM_INSTANCE_PROFILE")
}

// GetSSMInstanceProfileARN retrieves the SSM instance profile ARN of the current instance
func GetSSMInstanceProfileARN() string {
	return os.Getenv("SSM_INSTANCE_PROFILE_ARN")
}

// GetAWSSession retrieves an AWS session
func GetAWSSession() (*session.Session, error) {
	return session.NewSession(&aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))}, nil)
}