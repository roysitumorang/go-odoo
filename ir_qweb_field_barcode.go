package odoo

import (
	"fmt"
)

// IrQwebFieldBarcode represents ir.qweb.field.barcode model.
type IrQwebFieldBarcode struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// IrQwebFieldBarcodes represents array of ir.qweb.field.barcode model.
type IrQwebFieldBarcodes []IrQwebFieldBarcode

// IrQwebFieldBarcodeModel is the odoo model name.
const IrQwebFieldBarcodeModel = "ir.qweb.field.barcode"

// Many2One convert IrQwebFieldBarcode to *Many2One.
func (iqfb *IrQwebFieldBarcode) Many2One() *Many2One {
	return NewMany2One(iqfb.Id.Get(), "")
}

// CreateIrQwebFieldBarcode creates a new ir.qweb.field.barcode model and returns its id.
func (c *Client) CreateIrQwebFieldBarcode(iqfb *IrQwebFieldBarcode) (int64, error) {
	return c.Create(IrQwebFieldBarcodeModel, iqfb)
}

// UpdateIrQwebFieldBarcode updates an existing ir.qweb.field.barcode record.
func (c *Client) UpdateIrQwebFieldBarcode(iqfb *IrQwebFieldBarcode) error {
	return c.UpdateIrQwebFieldBarcodes([]int64{iqfb.Id.Get()}, iqfb)
}

// UpdateIrQwebFieldBarcodes updates existing ir.qweb.field.barcode records.
// All records (represented by ids) will be updated by iqfb values.
func (c *Client) UpdateIrQwebFieldBarcodes(ids []int64, iqfb *IrQwebFieldBarcode) error {
	return c.Update(IrQwebFieldBarcodeModel, ids, iqfb)
}

// DeleteIrQwebFieldBarcode deletes an existing ir.qweb.field.barcode record.
func (c *Client) DeleteIrQwebFieldBarcode(id int64) error {
	return c.DeleteIrQwebFieldBarcodes([]int64{id})
}

// DeleteIrQwebFieldBarcodes deletes existing ir.qweb.field.barcode records.
func (c *Client) DeleteIrQwebFieldBarcodes(ids []int64) error {
	return c.Delete(IrQwebFieldBarcodeModel, ids)
}

// GetIrQwebFieldBarcode gets ir.qweb.field.barcode existing record.
func (c *Client) GetIrQwebFieldBarcode(id int64) (*IrQwebFieldBarcode, error) {
	iqfbs, err := c.GetIrQwebFieldBarcodes([]int64{id})
	if err != nil {
		return nil, err
	}
	if iqfbs != nil && len(*iqfbs) > 0 {
		return &((*iqfbs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.qweb.field.barcode not found", id)
}

// GetIrQwebFieldBarcodes gets ir.qweb.field.barcode existing records.
func (c *Client) GetIrQwebFieldBarcodes(ids []int64) (*IrQwebFieldBarcodes, error) {
	iqfbs := &IrQwebFieldBarcodes{}
	if err := c.Read(IrQwebFieldBarcodeModel, ids, nil, iqfbs); err != nil {
		return nil, err
	}
	return iqfbs, nil
}

// FindIrQwebFieldBarcode finds ir.qweb.field.barcode record by querying it with criteria.
func (c *Client) FindIrQwebFieldBarcode(criteria *Criteria) (*IrQwebFieldBarcode, error) {
	iqfbs := &IrQwebFieldBarcodes{}
	if err := c.SearchRead(IrQwebFieldBarcodeModel, criteria, NewOptions().Limit(1), iqfbs); err != nil {
		return nil, err
	}
	if iqfbs != nil && len(*iqfbs) > 0 {
		return &((*iqfbs)[0]), nil
	}
	return nil, fmt.Errorf("no ir.qweb.field.barcode was found with criteria %v", criteria)
}

// FindIrQwebFieldBarcodes finds ir.qweb.field.barcode records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldBarcodes(criteria *Criteria, options *Options) (*IrQwebFieldBarcodes, error) {
	iqfbs := &IrQwebFieldBarcodes{}
	if err := c.SearchRead(IrQwebFieldBarcodeModel, criteria, options, iqfbs); err != nil {
		return nil, err
	}
	return iqfbs, nil
}

// FindIrQwebFieldBarcodeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldBarcodeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrQwebFieldBarcodeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrQwebFieldBarcodeId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldBarcodeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldBarcodeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no ir.qweb.field.barcode was found with criteria %v and options %v", criteria, options)
}
