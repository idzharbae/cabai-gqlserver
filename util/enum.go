package util

import "github.com/idzharbae/cabai-gqlserver/gql/transaction/enum"

func OrderCodeToString(code int) string {
	switch code {
	case enum.OrderStatusFulfilledCode:
		return enum.OrderStatusFulfilledString
	case enum.OrderStatusWaitingForSellerCode:
		return enum.OrderStatusWaitingForSellerString
	case enum.OrderStatusOnShipmentCode:
		return enum.OrderStatusOnShipmentString
	case enum.OrderStatusRejectedByShopCode:
		return enum.OrderStatusRejectedByShopString
	}
	return ""
}

func OrderStringToCode(codeString string) int {
	switch codeString {
	case enum.OrderStatusFulfilledString:
		return enum.OrderStatusFulfilledCode
	case enum.OrderStatusWaitingForSellerString:
		return enum.OrderStatusWaitingForSellerCode
	case enum.OrderStatusOnShipmentString:
		return enum.OrderStatusOnShipmentCode
	case enum.OrderStatusRejectedByShopString:
		return enum.OrderStatusRejectedByShopCode
	}
	return 0
}

func PaymentCodeToString(code int) string {
	switch code {
	case enum.PaymentStatusPaidCode:
		return enum.PaymentStatusPaidString
	case enum.PaymentStatusPendingCode:
		return enum.PaymentStatusPendingString
	case enum.PaymentStatusRefundedCode:
		return enum.PaymentStatusRefundedString
	}
	return ""
}

func PaymentStringToCode(codeString string) int {
	switch codeString {
	case enum.PaymentStatusPaidString:
		return enum.PaymentStatusPaidCode
	case enum.PaymentStatusPendingString:
		return enum.PaymentStatusPendingCode
	case enum.PaymentStatusRefundedString:
		return enum.PaymentStatusRefundedCode
	}
	return 0
}
