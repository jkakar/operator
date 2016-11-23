// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package route53

import (
	"github.com/aws/aws-sdk-go/private/waiter"
)

func (c *Route53) WaitUntilResourceRecordSetsChanged(input *GetChangeInput) error {
	waiterCfg := waiter.Config{
		Operation:   "GetChange",
		Delay:       30,
		MaxAttempts: 60,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "ChangeInfo.Status",
				Expected: "INSYNC",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
