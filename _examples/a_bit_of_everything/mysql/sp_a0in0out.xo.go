package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// A0In0Out calls the stored procedure 'a_bit_of_everything.a_0_in_0_out()' on db.
func A0In0Out(ctx context.Context, db DB) error {
	// call a_bit_of_everything.a_0_in_0_out
	const sqlstr = `CALL a_bit_of_everything.a_0_in_0_out()`
	// run
	logf(sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr); err != nil {
		return logerror(err)
	}
	return nil
}
