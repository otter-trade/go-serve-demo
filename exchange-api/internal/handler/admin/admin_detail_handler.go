package admin

import (
	"github.com/otter-trade/go-serve-demo/common/i18n"
	"github.com/otter-trade/go-serve-demo/common/xresp"
	"github.com/otter-trade/go-serve-demo/exchange-api/internal/logic/admin"
	"github.com/otter-trade/go-serve-demo/exchange-api/internal/svc"
	"github.com/otter-trade/go-serve-demo/exchange-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func AdminDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminDetailReq

		if err := httpx.Parse(r, &req); err != nil {
			xresp.Fail(r, w, i18n.NewCodeError(i18n.ParseParamsError))
			return
		}

		if err := xresp.Validate.StructCtx(r.Context(), req); err != nil {
			xresp.Fail(r, w, err)
			return
		}

		l := admin.NewAdminDetailLogic(r, svcCtx)
		resp, err := l.AdminDetail(&req)
		xresp.Success(r, w, resp, err)
	}
}
