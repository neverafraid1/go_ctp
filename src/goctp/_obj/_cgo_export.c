/* Created by cgo - DO NOT EDIT. */
#include <stdlib.h>
#include "_cgo_export.h"

extern void crosscall2(void (*fn)(void *, int, __SIZE_TYPE__), void *, int, __SIZE_TYPE__);
extern __SIZE_TYPE__ _cgo_wait_runtime_init_done();
extern void _cgo_release_context(__SIZE_TYPE__);

extern char* _cgo_topofstack(void);
#define CGO_NO_SANITIZE_THREAD
#define _cgo_tsan_acquire()
#define _cgo_tsan_release()

extern void _cgoexp_9df2bd5a2456_GoExOnFrontConnected(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void GoExOnFrontConnected(void* p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		void* p0;
	} __attribute__((__packed__, __gcc_struct__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_9df2bd5a2456_GoExOnFrontConnected, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
