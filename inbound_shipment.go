package mws

import (
	"encoding/xml"
	"errors"
	"net/url"
	"time"

	"github.com/pkg4go/tools"
	"github.com/spf13/cast"
)

type shipmentStatus string

const (
	shipmentStatusWorking   shipmentStatus = "WORKING"
	shipmentStatusShipped                  = "SHIPPED"
	shipmentStatusInTransit                = "IN_TRANSIT"
	shipmentStatusDelivered                = "DELIVERED"
	shipmentStatusCheckedIn                = "CHECKED_IN"
	shipmentStatusReceiving                = "RECEIVING"
	shipmentStatusClosed                   = "CLOSED"
	shipmentStatusCancelled                = "CANCELLED"
	shipmentStatusDeleted                  = "DELETED"
	shipmentStatusError                    = "ERROR"
)

// ListInboundShipmentsParams ListInboundShipments API params
type ListInboundShipmentsParams struct {
	ShipmentStatusList []shipmentStatus
	LastUpdatedAfter   time.Time
	LastUpdatedBefore  time.Time
}

// ListInboundShipments ...
func (seller *Seller) ListInboundShipments(params ListInboundShipmentsParams) []ShipmentMember {
	opts := seller.genListInboundShipmentsParams(params, "")

	var members []ShipmentMember

	result, err := seller.requestInboundShipment(opts, false)
	tools.AssertError(err)

	members = append(members, result.ShipmentData.Members...)

	for {
		if result.NextToken == "" {
			break
		}

		result = seller.listInboundShipmentsByNextToken(result.NextToken)
		members = append(members, result.ShipmentData.Members...)
	}

	return members

}

func (seller *Seller) listInboundShipmentsByNextToken(nextToken string) ListInboundShipmentsResult {
	opts := seller.genListInboundShipmentsParams(ListInboundShipmentsParams{}, nextToken)

	result, err := seller.requestInboundShipment(opts, true)
	tools.AssertError(err)

	return result
}

func (seller *Seller) genListInboundShipmentsParams(params ListInboundShipmentsParams, nextToken string) string {
	v := url.Values{}

	v.Add("AWSAccessKeyId", seller.AccessKey)
	v.Add("SellerId", seller.SellerID)
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("Version", "2010-10-01")
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("SignatureVersion", "2")

	if nextToken != "" {
		v.Add("Action", "ListInboundShipmentsByNextToken")
		v.Add("NextToken", nextToken)
	} else {
		v.Add("Action", "ListInboundShipments")
		v.Add("LastUpdatedAfter", params.LastUpdatedAfter.Format(time.RFC3339))
		v.Add("LastUpdatedBefore", params.LastUpdatedBefore.Format(time.RFC3339))
		for key, val := range params.ShipmentStatusList {
			v.Add("ShipmentStatusList.member."+cast.ToString(key+1), string(val))
		}
	}

	s := v.Encode()

	return s
}

func (seller *Seller) requestInboundShipment(qs string, byNextToken bool) (ListInboundShipmentsResult, error) {
	body, err := seller.get(InboundShipmentPath, qs)
	tools.AssertError(err)

	if byNextToken {
		data := ListInboundShipmentsByNextTokenResponse{}
		err = xml.Unmarshal(body, &data)

		if err != nil {
			return data.ListInboundShipmentsResult, errors.New(string(body))
		}

		return data.ListInboundShipmentsResult, nil
	}

	data := ListInboundShipmentsResponse{}
	err = xml.Unmarshal(body, &data)

	if err != nil {
		return data.ListInboundShipmentsResult, errors.New(string(body))
	}

	return data.ListInboundShipmentsResult, nil

}
