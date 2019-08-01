// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42615|42625

package nvm

import "errors"

// Error list of package nvm.
var (
	ErrNaN      = errors.New("nvm: Not a Number")
	ErrNaV      = errors.New("nvm: Not a Vector")
	ErrNaM      = errors.New("nvm: Not a Matrix")
	ErrDim      = errors.New("nvm: Bad dimension")
	ErrIndex    = errors.New("nvm: Index out of range")
	ErrRows     = errors.New("nvm: Bad rows")
	ErrCols     = errors.New("nvm: Bad cols")
	ErrRowIndex = errors.New("nvm: Row index out of range")
	ErrColIndex = errors.New("nvm: Col index out of range")
	ErrShape    = errors.New("nvm: Shape mismatch")
	ErrInverse  = errors.New("nvm: Matrix is NOT Invertible")
	ErrImpType  = errors.New("nvm: No available implementation type")
)
