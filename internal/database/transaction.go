package database

import (
	"context"
	"sync"

	"git.devminer.xyz/devminer/unitel"
	"github.com/versia-pub/versia-go/ent"
)

func BeginTx(ctx context.Context, db *ent.Client, telemetry *unitel.Telemetry) (*Tx, error) {
	span := telemetry.StartSpan(ctx, "db.sql.transaction", "BeginTx")
	ctx = span.Context()

	tx, err := db.Tx(ctx)
	if err != nil {
		return nil, err
	}

	return newTx(tx, ctx, span), nil
}

type TxAction uint8

const (
	TxActionRollback TxAction = iota
	TxActionCommit
)

type Tx struct {
	*ent.Tx
	ctx  context.Context
	span *unitel.Span

	m      sync.Mutex
	action TxAction

	finishOnce func() error
}

func newTx(tx *ent.Tx, ctx context.Context, span *unitel.Span) *Tx {
	t := &Tx{
		Tx:   tx,
		ctx:  ctx,
		span: span,
	}

	t.finishOnce = sync.OnceValue(t.finish)

	return t
}

func (t *Tx) MarkForCommit() {
	t.m.Lock()
	defer t.m.Unlock()

	t.action = TxActionCommit
}

func (t *Tx) finish() error {
	t.m.Lock()
	defer t.m.Unlock()
	defer t.span.End()

	var err error
	switch t.action {
	case TxActionCommit:
		err = t.Tx.Commit()
	case TxActionRollback:
		err = t.Tx.Rollback()
	}
	if err != nil {
		t.span.CaptureError(err)
	}

	return err
}

func (t *Tx) Context() context.Context {
	return t.ctx
}

func (t *Tx) Finish() error {
	return t.finishOnce()
}
