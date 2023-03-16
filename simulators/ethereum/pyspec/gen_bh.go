// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package main

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ = (*blockHeaderUnmarshaling)(nil)

// MarshalJSON marshals as JSON.
func (b blockHeader) MarshalJSON() ([]byte, error) {
	type blockHeader struct {
		ParentHash      common.Hash           `json:"parentHash"`
		UncleHash       common.Hash           `json:"sha3Uncles"`
		UncleHashAlt    common.Hash           `json:"uncleHash"`
		Coinbase        common.Address        `json:"coinbase"`
		CoinbaseAlt     common.Address        `json:"author"`
		CoinbaseAlt2    common.Address        `json:"miner"`
		StateRoot       common.Hash           `json:"stateRoot"`
		ReceiptTrie     common.Hash           `json:"receiptsRoot"`
		ReceiptTrieAlt  common.Hash           `json:"receiptTrie"`
		Bloom           types.Bloom           `json:"bloom"`
		Difficulty      *math.HexOrDecimal256 `json:"difficulty"`
		Number          *math.HexOrDecimal256 `json:"number"`
		GasLimit        math.HexOrDecimal64   `json:"gasLimit"`
		GasUsed         math.HexOrDecimal64   `json:"gasUsed"`
		Timestamp       *math.HexOrDecimal256 `json:"timestamp"`
		ExtraData       hexutil.Bytes         `json:"extraData"`
		MixHash         common.Hash           `json:"mixHash"`
		Nonce           types.BlockNonce      `json:"nonce"`
		BaseFee         *math.HexOrDecimal256 `json:"baseFeePerGas"`
		Hash            common.Hash           `json:"hash"`
		WithdrawalsRoot common.Hash           `json:"withdrawalsRoot"`
	}
	var enc blockHeader
	enc.ParentHash = b.ParentHash
	enc.UncleHash = b.UncleHash
	enc.UncleHashAlt = b.UncleHashAlt
	enc.Coinbase = b.Coinbase
	enc.CoinbaseAlt = b.CoinbaseAlt
	enc.CoinbaseAlt2 = b.CoinbaseAlt2
	enc.StateRoot = b.StateRoot
	enc.ReceiptTrie = b.ReceiptTrie
	enc.ReceiptTrieAlt = b.ReceiptTrieAlt
	enc.Bloom = b.Bloom
	enc.Difficulty = (*math.HexOrDecimal256)(b.Difficulty)
	enc.Number = (*math.HexOrDecimal256)(b.Number)
	enc.GasLimit = math.HexOrDecimal64(b.GasLimit)
	enc.GasUsed = math.HexOrDecimal64(b.GasUsed)
	enc.Timestamp = (*math.HexOrDecimal256)(b.Timestamp)
	enc.ExtraData = b.ExtraData
	enc.MixHash = b.MixHash
	enc.Nonce = b.Nonce
	enc.BaseFee = (*math.HexOrDecimal256)(b.BaseFee)
	enc.Hash = b.Hash
	enc.WithdrawalsRoot = b.WithdrawalsRoot
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (b *blockHeader) UnmarshalJSON(input []byte) error {
	type blockHeader struct {
		ParentHash      *common.Hash          `json:"parentHash"`
		UncleHash       *common.Hash          `json:"sha3Uncles"`
		UncleHashAlt    *common.Hash          `json:"uncleHash"`
		Coinbase        *common.Address       `json:"coinbase"`
		CoinbaseAlt     *common.Address       `json:"author"`
		CoinbaseAlt2    *common.Address       `json:"miner"`
		StateRoot       *common.Hash          `json:"stateRoot"`
		ReceiptTrie     *common.Hash          `json:"receiptsRoot"`
		ReceiptTrieAlt  *common.Hash          `json:"receiptTrie"`
		Bloom           *types.Bloom          `json:"bloom"`
		Difficulty      *math.HexOrDecimal256 `json:"difficulty"`
		Number          *math.HexOrDecimal256 `json:"number"`
		GasLimit        *math.HexOrDecimal64  `json:"gasLimit"`
		GasUsed         *math.HexOrDecimal64  `json:"gasUsed"`
		Timestamp       *math.HexOrDecimal256 `json:"timestamp"`
		ExtraData       *hexutil.Bytes        `json:"extraData"`
		MixHash         *common.Hash          `json:"mixHash"`
		Nonce           *types.BlockNonce     `json:"nonce"`
		BaseFee         *math.HexOrDecimal256 `json:"baseFeePerGas"`
		Hash            *common.Hash          `json:"hash"`
		WithdrawalsRoot *common.Hash          `json:"withdrawalsRoot"`
	}
	var dec blockHeader
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash != nil {
		b.ParentHash = *dec.ParentHash
	}
	if dec.UncleHash != nil {
		b.UncleHash = *dec.UncleHash
	}
	if dec.UncleHashAlt != nil {
		b.UncleHashAlt = *dec.UncleHashAlt
	}
	if dec.Coinbase != nil {
		b.Coinbase = *dec.Coinbase
	}
	if dec.CoinbaseAlt != nil {
		b.CoinbaseAlt = *dec.CoinbaseAlt
	}
	if dec.CoinbaseAlt2 != nil {
		b.CoinbaseAlt2 = *dec.CoinbaseAlt2
	}
	if dec.StateRoot != nil {
		b.StateRoot = *dec.StateRoot
	}
	if dec.ReceiptTrie != nil {
		b.ReceiptTrie = *dec.ReceiptTrie
	}
	if dec.ReceiptTrieAlt != nil {
		b.ReceiptTrieAlt = *dec.ReceiptTrieAlt
	}
	if dec.Bloom != nil {
		b.Bloom = *dec.Bloom
	}
	if dec.Difficulty != nil {
		b.Difficulty = (*big.Int)(dec.Difficulty)
	}
	if dec.Number != nil {
		b.Number = (*big.Int)(dec.Number)
	}
	if dec.GasLimit != nil {
		b.GasLimit = uint64(*dec.GasLimit)
	}
	if dec.GasUsed != nil {
		b.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.Timestamp != nil {
		b.Timestamp = (*big.Int)(dec.Timestamp)
	}
	if dec.ExtraData != nil {
		b.ExtraData = *dec.ExtraData
	}
	if dec.MixHash != nil {
		b.MixHash = *dec.MixHash
	}
	if dec.Nonce != nil {
		b.Nonce = *dec.Nonce
	}
	if dec.BaseFee != nil {
		b.BaseFee = (*big.Int)(dec.BaseFee)
	}
	if dec.Hash != nil {
		b.Hash = *dec.Hash
	}
	if dec.WithdrawalsRoot != nil {
		b.WithdrawalsRoot = *dec.WithdrawalsRoot
	}
	return nil
}