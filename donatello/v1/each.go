package v1

import (
	"context"
	"fmt"
	"time"
)

// Count of items per page
var DonatesPerPage = 20

// Pause between requests, Because we can catch:
//
// {"success": false, "message": "Забагато запитів. Спробуйте пізніше."}
var ThrottlingInMilliseconds = 1000

// EachDonate is iterate over all donates and use pagination and throttling.
// Always check returned error and use transaction for preventing to save bad data.
// Count of items per page and throttling value can be changed once in whole package, it's not a constants
//
// Please throw error inside fn function to break whole process, or return just nil. Example:
//
//	isFound := false
//	client.EachDonate(ctx, func(donate *v1.ResponseDonatesContent) error {
//		if donate.PubID == "12345" {
//			fmt.Printf("Donate: %#v\n", donate)
//			isFound = true
//		}
//		if isFound {
//			return fmt.Errorf("break")
//		}
//		return nil
//	})
func (c *Client) EachDonate(ctx context.Context, fn func(donate *ResponseDonatesContent) error) error {
	var resp *ResponseDonates
	var err error
	p := int64(0)

	for {
		resp, err = c.Donates(ctx, p, int64(DonatesPerPage))
		if err != nil {
			return err
		}

		if !resp.Success {
			return fmt.Errorf("%s", resp.Message)
		}

		if len(resp.Content) > 0 {
			for _, donate := range resp.Content {
				err = fn(&donate)
				if err != nil {
					return err
				}

				select {
				case <-ctx.Done():
					return fmt.Errorf("context canceled")
				default:
				}
			}
		}

		p++

		if err != nil {
			break
		}

		if p >= resp.Pages {
			break
		}

		if resp.Page >= resp.Pages {
			break
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("context canceled")
		default:
		}

		// Throttling
		time.Sleep(time.Duration(ThrottlingInMilliseconds) * time.Millisecond)
	}

	return err
}
