// Code generated by ent, DO NOT EDIT.

package utxoswap

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/lightsparkdev/spark/so/ent/predicate"
	"github.com/lightsparkdev/spark/so/ent/schema"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldUpdateTime, v))
}

// CreditAmountSats applies equality check predicate on the "credit_amount_sats" field. It's identical to CreditAmountSatsEQ.
func CreditAmountSats(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldCreditAmountSats, v))
}

// MaxFeeSats applies equality check predicate on the "max_fee_sats" field. It's identical to MaxFeeSatsEQ.
func MaxFeeSats(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldMaxFeeSats, v))
}

// SspSignature applies equality check predicate on the "ssp_signature" field. It's identical to SspSignatureEQ.
func SspSignature(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldSspSignature, v))
}

// SspIdentityPublicKey applies equality check predicate on the "ssp_identity_public_key" field. It's identical to SspIdentityPublicKeyEQ.
func SspIdentityPublicKey(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldSspIdentityPublicKey, v))
}

// UserSignature applies equality check predicate on the "user_signature" field. It's identical to UserSignatureEQ.
func UserSignature(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldUserSignature, v))
}

// UserIdentityPublicKey applies equality check predicate on the "user_identity_public_key" field. It's identical to UserIdentityPublicKeyEQ.
func UserIdentityPublicKey(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldUserIdentityPublicKey, v))
}

// CoordinatorIdentityPublicKey applies equality check predicate on the "coordinator_identity_public_key" field. It's identical to CoordinatorIdentityPublicKeyEQ.
func CoordinatorIdentityPublicKey(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldCoordinatorIdentityPublicKey, v))
}

// RequestedTransferID applies equality check predicate on the "requested_transfer_id" field. It's identical to RequestedTransferIDEQ.
func RequestedTransferID(v uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldRequestedTransferID, v))
}

// SpendTxSigningResult applies equality check predicate on the "spend_tx_signing_result" field. It's identical to SpendTxSigningResultEQ.
func SpendTxSigningResult(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldSpendTxSigningResult, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLTE(FieldUpdateTime, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v schema.UtxoSwapStatus) predicate.UtxoSwap {
	vc := v
	return predicate.UtxoSwap(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v schema.UtxoSwapStatus) predicate.UtxoSwap {
	vc := v
	return predicate.UtxoSwap(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...schema.UtxoSwapStatus) predicate.UtxoSwap {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UtxoSwap(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...schema.UtxoSwapStatus) predicate.UtxoSwap {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UtxoSwap(sql.FieldNotIn(FieldStatus, v...))
}

// RequestTypeEQ applies the EQ predicate on the "request_type" field.
func RequestTypeEQ(v schema.UtxoSwapRequestType) predicate.UtxoSwap {
	vc := v
	return predicate.UtxoSwap(sql.FieldEQ(FieldRequestType, vc))
}

// RequestTypeNEQ applies the NEQ predicate on the "request_type" field.
func RequestTypeNEQ(v schema.UtxoSwapRequestType) predicate.UtxoSwap {
	vc := v
	return predicate.UtxoSwap(sql.FieldNEQ(FieldRequestType, vc))
}

// RequestTypeIn applies the In predicate on the "request_type" field.
func RequestTypeIn(vs ...schema.UtxoSwapRequestType) predicate.UtxoSwap {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UtxoSwap(sql.FieldIn(FieldRequestType, v...))
}

// RequestTypeNotIn applies the NotIn predicate on the "request_type" field.
func RequestTypeNotIn(vs ...schema.UtxoSwapRequestType) predicate.UtxoSwap {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UtxoSwap(sql.FieldNotIn(FieldRequestType, v...))
}

// CreditAmountSatsEQ applies the EQ predicate on the "credit_amount_sats" field.
func CreditAmountSatsEQ(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldCreditAmountSats, v))
}

// CreditAmountSatsNEQ applies the NEQ predicate on the "credit_amount_sats" field.
func CreditAmountSatsNEQ(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNEQ(FieldCreditAmountSats, v))
}

// CreditAmountSatsIn applies the In predicate on the "credit_amount_sats" field.
func CreditAmountSatsIn(vs ...uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIn(FieldCreditAmountSats, vs...))
}

// CreditAmountSatsNotIn applies the NotIn predicate on the "credit_amount_sats" field.
func CreditAmountSatsNotIn(vs ...uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotIn(FieldCreditAmountSats, vs...))
}

// CreditAmountSatsGT applies the GT predicate on the "credit_amount_sats" field.
func CreditAmountSatsGT(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGT(FieldCreditAmountSats, v))
}

// CreditAmountSatsGTE applies the GTE predicate on the "credit_amount_sats" field.
func CreditAmountSatsGTE(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGTE(FieldCreditAmountSats, v))
}

// CreditAmountSatsLT applies the LT predicate on the "credit_amount_sats" field.
func CreditAmountSatsLT(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLT(FieldCreditAmountSats, v))
}

// CreditAmountSatsLTE applies the LTE predicate on the "credit_amount_sats" field.
func CreditAmountSatsLTE(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLTE(FieldCreditAmountSats, v))
}

// CreditAmountSatsIsNil applies the IsNil predicate on the "credit_amount_sats" field.
func CreditAmountSatsIsNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIsNull(FieldCreditAmountSats))
}

// CreditAmountSatsNotNil applies the NotNil predicate on the "credit_amount_sats" field.
func CreditAmountSatsNotNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotNull(FieldCreditAmountSats))
}

// MaxFeeSatsEQ applies the EQ predicate on the "max_fee_sats" field.
func MaxFeeSatsEQ(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldMaxFeeSats, v))
}

// MaxFeeSatsNEQ applies the NEQ predicate on the "max_fee_sats" field.
func MaxFeeSatsNEQ(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNEQ(FieldMaxFeeSats, v))
}

// MaxFeeSatsIn applies the In predicate on the "max_fee_sats" field.
func MaxFeeSatsIn(vs ...uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIn(FieldMaxFeeSats, vs...))
}

// MaxFeeSatsNotIn applies the NotIn predicate on the "max_fee_sats" field.
func MaxFeeSatsNotIn(vs ...uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotIn(FieldMaxFeeSats, vs...))
}

// MaxFeeSatsGT applies the GT predicate on the "max_fee_sats" field.
func MaxFeeSatsGT(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGT(FieldMaxFeeSats, v))
}

// MaxFeeSatsGTE applies the GTE predicate on the "max_fee_sats" field.
func MaxFeeSatsGTE(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGTE(FieldMaxFeeSats, v))
}

// MaxFeeSatsLT applies the LT predicate on the "max_fee_sats" field.
func MaxFeeSatsLT(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLT(FieldMaxFeeSats, v))
}

// MaxFeeSatsLTE applies the LTE predicate on the "max_fee_sats" field.
func MaxFeeSatsLTE(v uint64) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLTE(FieldMaxFeeSats, v))
}

// MaxFeeSatsIsNil applies the IsNil predicate on the "max_fee_sats" field.
func MaxFeeSatsIsNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIsNull(FieldMaxFeeSats))
}

// MaxFeeSatsNotNil applies the NotNil predicate on the "max_fee_sats" field.
func MaxFeeSatsNotNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotNull(FieldMaxFeeSats))
}

// SspSignatureEQ applies the EQ predicate on the "ssp_signature" field.
func SspSignatureEQ(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldSspSignature, v))
}

// SspSignatureNEQ applies the NEQ predicate on the "ssp_signature" field.
func SspSignatureNEQ(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNEQ(FieldSspSignature, v))
}

// SspSignatureIn applies the In predicate on the "ssp_signature" field.
func SspSignatureIn(vs ...[]byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIn(FieldSspSignature, vs...))
}

// SspSignatureNotIn applies the NotIn predicate on the "ssp_signature" field.
func SspSignatureNotIn(vs ...[]byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotIn(FieldSspSignature, vs...))
}

// SspSignatureGT applies the GT predicate on the "ssp_signature" field.
func SspSignatureGT(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGT(FieldSspSignature, v))
}

// SspSignatureGTE applies the GTE predicate on the "ssp_signature" field.
func SspSignatureGTE(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGTE(FieldSspSignature, v))
}

// SspSignatureLT applies the LT predicate on the "ssp_signature" field.
func SspSignatureLT(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLT(FieldSspSignature, v))
}

// SspSignatureLTE applies the LTE predicate on the "ssp_signature" field.
func SspSignatureLTE(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLTE(FieldSspSignature, v))
}

// SspSignatureIsNil applies the IsNil predicate on the "ssp_signature" field.
func SspSignatureIsNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIsNull(FieldSspSignature))
}

// SspSignatureNotNil applies the NotNil predicate on the "ssp_signature" field.
func SspSignatureNotNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotNull(FieldSspSignature))
}

// SspIdentityPublicKeyEQ applies the EQ predicate on the "ssp_identity_public_key" field.
func SspIdentityPublicKeyEQ(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldSspIdentityPublicKey, v))
}

// SspIdentityPublicKeyNEQ applies the NEQ predicate on the "ssp_identity_public_key" field.
func SspIdentityPublicKeyNEQ(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNEQ(FieldSspIdentityPublicKey, v))
}

// SspIdentityPublicKeyIn applies the In predicate on the "ssp_identity_public_key" field.
func SspIdentityPublicKeyIn(vs ...[]byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIn(FieldSspIdentityPublicKey, vs...))
}

// SspIdentityPublicKeyNotIn applies the NotIn predicate on the "ssp_identity_public_key" field.
func SspIdentityPublicKeyNotIn(vs ...[]byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotIn(FieldSspIdentityPublicKey, vs...))
}

// SspIdentityPublicKeyGT applies the GT predicate on the "ssp_identity_public_key" field.
func SspIdentityPublicKeyGT(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGT(FieldSspIdentityPublicKey, v))
}

// SspIdentityPublicKeyGTE applies the GTE predicate on the "ssp_identity_public_key" field.
func SspIdentityPublicKeyGTE(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGTE(FieldSspIdentityPublicKey, v))
}

// SspIdentityPublicKeyLT applies the LT predicate on the "ssp_identity_public_key" field.
func SspIdentityPublicKeyLT(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLT(FieldSspIdentityPublicKey, v))
}

// SspIdentityPublicKeyLTE applies the LTE predicate on the "ssp_identity_public_key" field.
func SspIdentityPublicKeyLTE(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLTE(FieldSspIdentityPublicKey, v))
}

// SspIdentityPublicKeyIsNil applies the IsNil predicate on the "ssp_identity_public_key" field.
func SspIdentityPublicKeyIsNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIsNull(FieldSspIdentityPublicKey))
}

// SspIdentityPublicKeyNotNil applies the NotNil predicate on the "ssp_identity_public_key" field.
func SspIdentityPublicKeyNotNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotNull(FieldSspIdentityPublicKey))
}

// UserSignatureEQ applies the EQ predicate on the "user_signature" field.
func UserSignatureEQ(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldUserSignature, v))
}

// UserSignatureNEQ applies the NEQ predicate on the "user_signature" field.
func UserSignatureNEQ(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNEQ(FieldUserSignature, v))
}

// UserSignatureIn applies the In predicate on the "user_signature" field.
func UserSignatureIn(vs ...[]byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIn(FieldUserSignature, vs...))
}

// UserSignatureNotIn applies the NotIn predicate on the "user_signature" field.
func UserSignatureNotIn(vs ...[]byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotIn(FieldUserSignature, vs...))
}

// UserSignatureGT applies the GT predicate on the "user_signature" field.
func UserSignatureGT(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGT(FieldUserSignature, v))
}

// UserSignatureGTE applies the GTE predicate on the "user_signature" field.
func UserSignatureGTE(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGTE(FieldUserSignature, v))
}

// UserSignatureLT applies the LT predicate on the "user_signature" field.
func UserSignatureLT(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLT(FieldUserSignature, v))
}

// UserSignatureLTE applies the LTE predicate on the "user_signature" field.
func UserSignatureLTE(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLTE(FieldUserSignature, v))
}

// UserSignatureIsNil applies the IsNil predicate on the "user_signature" field.
func UserSignatureIsNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIsNull(FieldUserSignature))
}

// UserSignatureNotNil applies the NotNil predicate on the "user_signature" field.
func UserSignatureNotNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotNull(FieldUserSignature))
}

// UserIdentityPublicKeyEQ applies the EQ predicate on the "user_identity_public_key" field.
func UserIdentityPublicKeyEQ(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldUserIdentityPublicKey, v))
}

// UserIdentityPublicKeyNEQ applies the NEQ predicate on the "user_identity_public_key" field.
func UserIdentityPublicKeyNEQ(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNEQ(FieldUserIdentityPublicKey, v))
}

// UserIdentityPublicKeyIn applies the In predicate on the "user_identity_public_key" field.
func UserIdentityPublicKeyIn(vs ...[]byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIn(FieldUserIdentityPublicKey, vs...))
}

// UserIdentityPublicKeyNotIn applies the NotIn predicate on the "user_identity_public_key" field.
func UserIdentityPublicKeyNotIn(vs ...[]byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotIn(FieldUserIdentityPublicKey, vs...))
}

// UserIdentityPublicKeyGT applies the GT predicate on the "user_identity_public_key" field.
func UserIdentityPublicKeyGT(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGT(FieldUserIdentityPublicKey, v))
}

// UserIdentityPublicKeyGTE applies the GTE predicate on the "user_identity_public_key" field.
func UserIdentityPublicKeyGTE(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGTE(FieldUserIdentityPublicKey, v))
}

// UserIdentityPublicKeyLT applies the LT predicate on the "user_identity_public_key" field.
func UserIdentityPublicKeyLT(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLT(FieldUserIdentityPublicKey, v))
}

// UserIdentityPublicKeyLTE applies the LTE predicate on the "user_identity_public_key" field.
func UserIdentityPublicKeyLTE(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLTE(FieldUserIdentityPublicKey, v))
}

// UserIdentityPublicKeyIsNil applies the IsNil predicate on the "user_identity_public_key" field.
func UserIdentityPublicKeyIsNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIsNull(FieldUserIdentityPublicKey))
}

// UserIdentityPublicKeyNotNil applies the NotNil predicate on the "user_identity_public_key" field.
func UserIdentityPublicKeyNotNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotNull(FieldUserIdentityPublicKey))
}

// CoordinatorIdentityPublicKeyEQ applies the EQ predicate on the "coordinator_identity_public_key" field.
func CoordinatorIdentityPublicKeyEQ(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldCoordinatorIdentityPublicKey, v))
}

// CoordinatorIdentityPublicKeyNEQ applies the NEQ predicate on the "coordinator_identity_public_key" field.
func CoordinatorIdentityPublicKeyNEQ(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNEQ(FieldCoordinatorIdentityPublicKey, v))
}

// CoordinatorIdentityPublicKeyIn applies the In predicate on the "coordinator_identity_public_key" field.
func CoordinatorIdentityPublicKeyIn(vs ...[]byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIn(FieldCoordinatorIdentityPublicKey, vs...))
}

// CoordinatorIdentityPublicKeyNotIn applies the NotIn predicate on the "coordinator_identity_public_key" field.
func CoordinatorIdentityPublicKeyNotIn(vs ...[]byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotIn(FieldCoordinatorIdentityPublicKey, vs...))
}

// CoordinatorIdentityPublicKeyGT applies the GT predicate on the "coordinator_identity_public_key" field.
func CoordinatorIdentityPublicKeyGT(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGT(FieldCoordinatorIdentityPublicKey, v))
}

// CoordinatorIdentityPublicKeyGTE applies the GTE predicate on the "coordinator_identity_public_key" field.
func CoordinatorIdentityPublicKeyGTE(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGTE(FieldCoordinatorIdentityPublicKey, v))
}

// CoordinatorIdentityPublicKeyLT applies the LT predicate on the "coordinator_identity_public_key" field.
func CoordinatorIdentityPublicKeyLT(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLT(FieldCoordinatorIdentityPublicKey, v))
}

// CoordinatorIdentityPublicKeyLTE applies the LTE predicate on the "coordinator_identity_public_key" field.
func CoordinatorIdentityPublicKeyLTE(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLTE(FieldCoordinatorIdentityPublicKey, v))
}

// RequestedTransferIDEQ applies the EQ predicate on the "requested_transfer_id" field.
func RequestedTransferIDEQ(v uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldRequestedTransferID, v))
}

// RequestedTransferIDNEQ applies the NEQ predicate on the "requested_transfer_id" field.
func RequestedTransferIDNEQ(v uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNEQ(FieldRequestedTransferID, v))
}

// RequestedTransferIDIn applies the In predicate on the "requested_transfer_id" field.
func RequestedTransferIDIn(vs ...uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIn(FieldRequestedTransferID, vs...))
}

// RequestedTransferIDNotIn applies the NotIn predicate on the "requested_transfer_id" field.
func RequestedTransferIDNotIn(vs ...uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotIn(FieldRequestedTransferID, vs...))
}

// RequestedTransferIDGT applies the GT predicate on the "requested_transfer_id" field.
func RequestedTransferIDGT(v uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGT(FieldRequestedTransferID, v))
}

// RequestedTransferIDGTE applies the GTE predicate on the "requested_transfer_id" field.
func RequestedTransferIDGTE(v uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGTE(FieldRequestedTransferID, v))
}

// RequestedTransferIDLT applies the LT predicate on the "requested_transfer_id" field.
func RequestedTransferIDLT(v uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLT(FieldRequestedTransferID, v))
}

// RequestedTransferIDLTE applies the LTE predicate on the "requested_transfer_id" field.
func RequestedTransferIDLTE(v uuid.UUID) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLTE(FieldRequestedTransferID, v))
}

// RequestedTransferIDIsNil applies the IsNil predicate on the "requested_transfer_id" field.
func RequestedTransferIDIsNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIsNull(FieldRequestedTransferID))
}

// RequestedTransferIDNotNil applies the NotNil predicate on the "requested_transfer_id" field.
func RequestedTransferIDNotNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotNull(FieldRequestedTransferID))
}

// SpendTxSigningResultEQ applies the EQ predicate on the "spend_tx_signing_result" field.
func SpendTxSigningResultEQ(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldEQ(FieldSpendTxSigningResult, v))
}

// SpendTxSigningResultNEQ applies the NEQ predicate on the "spend_tx_signing_result" field.
func SpendTxSigningResultNEQ(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNEQ(FieldSpendTxSigningResult, v))
}

// SpendTxSigningResultIn applies the In predicate on the "spend_tx_signing_result" field.
func SpendTxSigningResultIn(vs ...[]byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIn(FieldSpendTxSigningResult, vs...))
}

// SpendTxSigningResultNotIn applies the NotIn predicate on the "spend_tx_signing_result" field.
func SpendTxSigningResultNotIn(vs ...[]byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotIn(FieldSpendTxSigningResult, vs...))
}

// SpendTxSigningResultGT applies the GT predicate on the "spend_tx_signing_result" field.
func SpendTxSigningResultGT(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGT(FieldSpendTxSigningResult, v))
}

// SpendTxSigningResultGTE applies the GTE predicate on the "spend_tx_signing_result" field.
func SpendTxSigningResultGTE(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldGTE(FieldSpendTxSigningResult, v))
}

// SpendTxSigningResultLT applies the LT predicate on the "spend_tx_signing_result" field.
func SpendTxSigningResultLT(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLT(FieldSpendTxSigningResult, v))
}

// SpendTxSigningResultLTE applies the LTE predicate on the "spend_tx_signing_result" field.
func SpendTxSigningResultLTE(v []byte) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldLTE(FieldSpendTxSigningResult, v))
}

// SpendTxSigningResultIsNil applies the IsNil predicate on the "spend_tx_signing_result" field.
func SpendTxSigningResultIsNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldIsNull(FieldSpendTxSigningResult))
}

// SpendTxSigningResultNotNil applies the NotNil predicate on the "spend_tx_signing_result" field.
func SpendTxSigningResultNotNil() predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.FieldNotNull(FieldSpendTxSigningResult))
}

// HasUtxo applies the HasEdge predicate on the "utxo" edge.
func HasUtxo() predicate.UtxoSwap {
	return predicate.UtxoSwap(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UtxoTable, UtxoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUtxoWith applies the HasEdge predicate on the "utxo" edge with a given conditions (other predicates).
func HasUtxoWith(preds ...predicate.Utxo) predicate.UtxoSwap {
	return predicate.UtxoSwap(func(s *sql.Selector) {
		step := newUtxoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTransfer applies the HasEdge predicate on the "transfer" edge.
func HasTransfer() predicate.UtxoSwap {
	return predicate.UtxoSwap(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TransferTable, TransferColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransferWith applies the HasEdge predicate on the "transfer" edge with a given conditions (other predicates).
func HasTransferWith(preds ...predicate.Transfer) predicate.UtxoSwap {
	return predicate.UtxoSwap(func(s *sql.Selector) {
		step := newTransferStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UtxoSwap) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UtxoSwap) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UtxoSwap) predicate.UtxoSwap {
	return predicate.UtxoSwap(sql.NotPredicates(p))
}
