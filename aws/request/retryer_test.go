package request

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stowelly/aws-sdk-go/aws/awserr"
)

func TestRequestThrottling(t *testing.T) {
	req := Request{}

	req.Error = awserr.New("Throttling", "", nil)
	assert.True(t, req.IsErrorThrottle())
}
