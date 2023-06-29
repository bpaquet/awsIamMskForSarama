package awsIamMskForSarama

type GeneratorFunc struct {
	AwsRegion string
	addr      string
	userAgent string
	done      bool
}

func (x *GeneratorFunc) Begin(addr string) error {
	x.addr = addr
	x.userAgent = "awsIamMskForSarama"
	x.done = false
	return nil
}

func (x *GeneratorFunc) Step(challenge string) (string, error) {
	if challenge == "" {
		response, err := buildAwsIAMPayload(x.addr, x.userAgent, AWSMSKIAMConfig{Region: x.AwsRegion})
		if err != nil {
			return "", err
		}
		return string(response), nil
	}
	return "", nil
}

func (x *GeneratorFunc) Done() bool {
	result := x.done
	x.done = true
	return result
}

func (x *GeneratorFunc) MechanismName() string {
	return "AWS_MSK_IAM"
}
