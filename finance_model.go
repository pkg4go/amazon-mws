package mws

import "encoding/xml"

// ListFinancialEventsResponse ...
type ListFinancialEventsResponse struct {
	XMLName                   xml.Name                  `xml:"ListFinancialEventsResponse"`
	Text                      string                    `xml:",chardata"`
	Xmlns                     string                    `xml:"xmlns,attr"`
	ListFinancialEventsResult ListFinancialEventsResult `xml:"ListFinancialEventsResult"`
	ResponseMetadata          ResponseMetadata          `xml:"ResponseMetadata"`
}

// ListFinancialEventsByNextTokenResponse ...
type ListFinancialEventsByNextTokenResponse struct {
	XMLName                   xml.Name                  `xml:"ListFinancialEventsByNextTokenResponse"`
	Text                      string                    `xml:",chardata"`
	Xmlns                     string                    `xml:"xmlns,attr"`
	ListFinancialEventsResult ListFinancialEventsResult `xml:"ListFinancialEventsByNextTokenResult"`
	ResponseMetadata          ResponseMetadata          `xml:"ResponseMetadata"`
}

// ListFinancialEventsResult ...
type ListFinancialEventsResult struct {
	Text            string          `xml:",chardata"`
	NextToken       string          `xml:"NextToken"`
	FinancialEvents FinancialEvents `xml:"FinancialEvents"`
}

// FinancialEvents ...
type FinancialEvents struct {
	Text                                   string                     `xml:",chardata"`
	ProductAdsPaymentEventList             ProductAdsPaymentEventList `xml:"ProductAdsPaymentEventList"`
	RentalTransactionEventList             string                     `xml:"RentalTransactionEventList"`
	ServiceFeeEventList                    ServiceFeeEventList        `xml:"ServiceFeeEventList"`
	ShipmentSettleEventList                string                     `xml:"ShipmentSettleEventList"`
	ServiceProviderCreditEventList         string                     `xml:"ServiceProviderCreditEventList"`
	ImagingServicesFeeEventList            string                     `xml:"ImagingServicesFeeEventList"`
	SellerDealPaymentEventList             string                     `xml:"SellerDealPaymentEventList"`
	SellerReviewEnrollmentPaymentEventList string                     `xml:"SellerReviewEnrollmentPaymentEventList"`
	DebtRecoveryEventList                  string                     `xml:"DebtRecoveryEventList"`
	ShipmentEventList                      ShipmentEventList          `xml:"ShipmentEventList"`
	TaxWithholdingEventList                string                     `xml:"TaxWithholdingEventList"`
	GuaranteeClaimEventList                string                     `xml:"GuaranteeClaimEventList"`
	TDSReimbursementEventList              string                     `xml:"TDSReimbursementEventList"`
	ChargebackEventList                    string                     `xml:"ChargebackEventList"`
	NetworkComminglingTransactionEventList string                     `xml:"NetworkComminglingTransactionEventList"`
	LoanServicingEventList                 string                     `xml:"LoanServicingEventList"`
	RefundEventList                        RefundEventList            `xml:"RefundEventList"`
	RemovalShipmentEventList               string                     `xml:"RemovalShipmentEventList"`
	PerformanceBondRefundEventList         string                     `xml:"PerformanceBondRefundEventList"`
	AffordabilityExpenseReversalEventList  string                     `xml:"AffordabilityExpenseReversalEventList"`
	PayWithAmazonEventList                 string                     `xml:"PayWithAmazonEventList"`
	AdhocDisbursementEventList             string                     `xml:"AdhocDisbursementEventList"`
	CouponPaymentEventList                 CouponPaymentEventList     `xml:"CouponPaymentEventList"`
	ChargeRefundEventList                  string                     `xml:"ChargeRefundEventList"`
	RetrochargeEventList                   string                     `xml:"RetrochargeEventList"`
	TrialShipmentEventList                 string                     `xml:"TrialShipmentEventList"`
	SAFETReimbursementEventList            string                     `xml:"SAFETReimbursementEventList"`
	RemovalShipmentAdjustmentEventList     string                     `xml:"RemovalShipmentAdjustmentEventList"`
	FBALiquidationEventList                string                     `xml:"FBALiquidationEventList"`
	AffordabilityExpenseEventList          string                     `xml:"AffordabilityExpenseEventList"`
	AdjustmentEventList                    AdjustmentEventList        `xml:"AdjustmentEventList"`
}

// ProductAdsPaymentEventList ...
type ProductAdsPaymentEventList struct {
	Text                   string `xml:",chardata"`
	ProductAdsPaymentEvent struct {
		Text            string `xml:",chardata"`
		TransactionType string `xml:"transactionType"`
		BaseValue       struct {
			Text           string `xml:",chardata"`
			CurrencyAmount string `xml:"CurrencyAmount"`
			CurrencyCode   string `xml:"CurrencyCode"`
		} `xml:"baseValue"`
		TaxValue struct {
			Text           string `xml:",chardata"`
			CurrencyAmount string `xml:"CurrencyAmount"`
			CurrencyCode   string `xml:"CurrencyCode"`
		} `xml:"taxValue"`
		InvoiceID        string `xml:"invoiceId"`
		TransactionValue struct {
			Text           string `xml:",chardata"`
			CurrencyAmount string `xml:"CurrencyAmount"`
			CurrencyCode   string `xml:"CurrencyCode"`
		} `xml:"transactionValue"`
		PostedDate string `xml:"postedDate"`
	} `xml:"ProductAdsPaymentEvent"`
}

// ServiceFeeEventList ...
type ServiceFeeEventList struct {
	Text            string `xml:",chardata"`
	ServiceFeeEvent struct {
		Text          string `xml:",chardata"`
		FeeReason     string `xml:"FeeReason"`
		AmazonOrderID string `xml:"AmazonOrderId"`
		FeeList       struct {
			Text         string         `xml:",chardata"`
			FeeComponent []FeeComponent `xml:"FeeComponent"`
		} `xml:"FeeList"`
	} `xml:"ServiceFeeEvent"`
}

// ShipmentEventList ...
type ShipmentEventList struct {
	Text          string `xml:",chardata"`
	ShipmentEvent []struct {
		Text             string `xml:",chardata"`
		ShipmentItemList struct {
			Text         string `xml:",chardata"`
			ShipmentItem []struct {
				Text                string `xml:",chardata"`
				ItemTaxWithheldList struct {
					Text                 string `xml:",chardata"`
					TaxWithheldComponent struct {
						Text               string `xml:",chardata"`
						TaxCollectionModel string `xml:"TaxCollectionModel"`
						TaxesWithheld      struct {
							Text            string            `xml:",chardata"`
							ChargeComponent []ChargeComponent `xml:"ChargeComponent"`
						} `xml:"TaxesWithheld"`
					} `xml:"TaxWithheldComponent"`
				} `xml:"ItemTaxWithheldList"`
				ItemChargeList struct {
					Text            string            `xml:",chardata"`
					ChargeComponent []ChargeComponent `xml:"ChargeComponent"`
				} `xml:"ItemChargeList"`
				ItemFeeList struct {
					Text         string         `xml:",chardata"`
					FeeComponent []FeeComponent `xml:"FeeComponent"`
				} `xml:"ItemFeeList"`
				OrderItemID     string `xml:"OrderItemId"`
				QuantityShipped string `xml:"QuantityShipped"`
				SellerSKU       string `xml:"SellerSKU"`
				PromotionList   struct {
					Text      string      `xml:",chardata"`
					Promotion []Promotion `xml:"Promotion"`
				} `xml:"PromotionList"`
			} `xml:"ShipmentItem"`
		} `xml:"ShipmentItemList"`
		AmazonOrderID   string `xml:"AmazonOrderId"`
		PostedDate      string `xml:"PostedDate"`
		MarketplaceName string `xml:"MarketplaceName"`
		SellerOrderID   string `xml:"SellerOrderId"`
		ShipmentFeeList struct {
			Text         string       `xml:",chardata"`
			FeeComponent FeeComponent `xml:"FeeComponent"`
		} `xml:"ShipmentFeeList"`
	} `xml:"ShipmentEvent"`
}

// RefundEventList ...
type RefundEventList struct {
	Text          string `xml:",chardata"`
	ShipmentEvent struct {
		Text                       string `xml:",chardata"`
		AmazonOrderID              string `xml:"AmazonOrderId"`
		PostedDate                 string `xml:"PostedDate"`
		ShipmentItemAdjustmentList struct {
			Text         string `xml:",chardata"`
			ShipmentItem struct {
				Text                string `xml:",chardata"`
				ItemTaxWithheldList struct {
					Text                 string `xml:",chardata"`
					TaxWithheldComponent struct {
						Text               string `xml:",chardata"`
						TaxCollectionModel string `xml:"TaxCollectionModel"`
						TaxesWithheld      struct {
							Text            string          `xml:",chardata"`
							ChargeComponent ChargeComponent `xml:"ChargeComponent"`
						} `xml:"TaxesWithheld"`
					} `xml:"TaxWithheldComponent"`
				} `xml:"ItemTaxWithheldList"`
				ItemFeeAdjustmentList struct {
					Text         string         `xml:",chardata"`
					FeeComponent []FeeComponent `xml:"FeeComponent"`
				} `xml:"ItemFeeAdjustmentList"`
				OrderAdjustmentItemID    string `xml:"OrderAdjustmentItemId"`
				QuantityShipped          string `xml:"QuantityShipped"`
				ItemChargeAdjustmentList struct {
					Text            string            `xml:",chardata"`
					ChargeComponent []ChargeComponent `xml:"ChargeComponent"`
				} `xml:"ItemChargeAdjustmentList"`
				SellerSKU               string `xml:"SellerSKU"`
				PromotionAdjustmentList struct {
					Text      string      `xml:",chardata"`
					Promotion []Promotion `xml:"Promotion"`
				} `xml:"PromotionAdjustmentList"`
			} `xml:"ShipmentItem"`
		} `xml:"ShipmentItemAdjustmentList"`

		MarketplaceName string `xml:"MarketplaceName"`
		SellerOrderID   string `xml:"SellerOrderId"`
	} `xml:"ShipmentEvent"`
}

// CouponPaymentEventList ...
type CouponPaymentEventList struct {
	Text               string `xml:",chardata"`
	CouponPaymentEvent []struct {
		Text        string `xml:",chardata"`
		TotalAmount struct {
			Text           string `xml:",chardata"`
			CurrencyAmount string `xml:"CurrencyAmount"`
			CurrencyCode   string `xml:"CurrencyCode"`
		} `xml:"TotalAmount"`
		PaymentEventID          string          `xml:"PaymentEventId"`
		SellerCouponDescription string          `xml:"SellerCouponDescription"`
		FeeComponent            FeeComponent    `xml:"FeeComponent"`
		ChargeComponent         ChargeComponent `xml:"ChargeComponent"`
		CouponID                string          `xml:"CouponId"`
		ClipOrRedemptionCount   string          `xml:"ClipOrRedemptionCount"`
		PostedDate              string          `xml:"PostedDate"`
	} `xml:"CouponPaymentEvent"`
}

// AdjustmentEventList ...
type AdjustmentEventList struct {
	Text            string `xml:",chardata"`
	AdjustmentEvent []struct {
		Text               string `xml:",chardata"`
		AdjustmentType     string `xml:"AdjustmentType"`
		AdjustmentItemList struct {
			Text           string `xml:",chardata"`
			AdjustmentItem struct {
				Text          string `xml:",chardata"`
				PerUnitAmount struct {
					Text           string `xml:",chardata"`
					CurrencyAmount string `xml:"CurrencyAmount"`
					CurrencyCode   string `xml:"CurrencyCode"`
				} `xml:"PerUnitAmount"`
				TotalAmount struct {
					Text           string `xml:",chardata"`
					CurrencyAmount string `xml:"CurrencyAmount"`
					CurrencyCode   string `xml:"CurrencyCode"`
				} `xml:"TotalAmount"`
				Quantity           string `xml:"Quantity"`
				SellerSKU          string `xml:"SellerSKU"`
				ProductDescription string `xml:"ProductDescription"`
			} `xml:"AdjustmentItem"`
		} `xml:"AdjustmentItemList"`
		AdjustmentAmount struct {
			Text           string `xml:",chardata"`
			CurrencyAmount string `xml:"CurrencyAmount"`
			CurrencyCode   string `xml:"CurrencyCode"`
		} `xml:"AdjustmentAmount"`
		PostedDate string `xml:"PostedDate"`
	} `xml:"AdjustmentEvent"`
}

// ChargeComponent ...
type ChargeComponent struct {
	Text         string `xml:",chardata"`
	ChargeType   string `xml:"ChargeType"`
	ChargeAmount struct {
		Text           string `xml:",chardata"`
		CurrencyAmount string `xml:"CurrencyAmount"`
		CurrencyCode   string `xml:"CurrencyCode"`
	} `xml:"ChargeAmount"`
}

// FeeComponent ...
type FeeComponent struct {
	Text      string `xml:",chardata"`
	FeeType   string `xml:"FeeType"`
	FeeAmount struct {
		Text           string `xml:",chardata"`
		CurrencyAmount string `xml:"CurrencyAmount"`
		CurrencyCode   string `xml:"CurrencyCode"`
	} `xml:"FeeAmount"`
}

// Promotion ...
type Promotion struct {
	Text            string `xml:",chardata"`
	PromotionType   string `xml:"PromotionType"`
	PromotionAmount struct {
		Text           string `xml:",chardata"`
		CurrencyAmount string `xml:"CurrencyAmount"`
		CurrencyCode   string `xml:"CurrencyCode"`
	} `xml:"PromotionAmount"`
	PromotionID string `xml:"PromotionId"`
}
