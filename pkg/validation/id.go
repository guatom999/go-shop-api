package validation

import (
	"github.com/labstack/echo/v4"

	_adminException "github.com/guatom999/go-shop-api/pkg/admin/exception"
	_playerException "github.com/guatom999/go-shop-api/pkg/player/exception"
)

func AdminIDGetting(pctx echo.Context) (string, error) {
	if adminID, ok := pctx.Get("adminID").(string); !ok || adminID == "" {
		return "", &_adminException.AdminNotFound{AdminID: "unknow"}
	} else {
		return adminID, nil
	}

}

func PlayerIDGetting(pctx echo.Context) (string, error) {
	if playerID, ok := pctx.Get("playerID").(string); !ok || playerID == "" {
		return "", &_playerException.PlayerNotFound{PlayerID: "unknow"}
	} else {
		return playerID, nil
	}

}
