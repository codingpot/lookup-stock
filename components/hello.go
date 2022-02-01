package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Hello struct {
	app.Compo

	//name             string
	isAppInstallable bool
}

func (h *Hello) OnMount(ctx app.Context) {
	h.isAppInstallable = ctx.IsAppInstallable()
}

func (h *Hello) onInstallButtonClicked(ctx app.Context, e app.Event) {
	ctx.ShowAppInstallPrompt()
}

func (h *Hello) Render() app.UI {
	return app.Div().
		Body(
			app.H1().Text("Hello World!!"),

			app.If(h.isAppInstallable,
				app.Button().
					Text("Install App").
					OnClick(h.onInstallButtonClicked),
			),
		)
}
