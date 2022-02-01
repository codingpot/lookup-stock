package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Hello struct {
	app.Compo

	name             string
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
			app.H1().Body(
				app.Text("Hello, "),
				app.If(h.name != "",
					app.Text(h.name),
				).Else(
					app.Text("World!"),
				),
			),
			app.P().Body(
				app.Input().
					Type("text").
					Value(h.name).
					Placeholder("What is your name?").
					AutoFocus(true).
					OnChange(h.ValueTo(&h.name)),
			),
			app.If(h.isAppInstallable,
				app.Button().
					Text("Install App").
					OnClick(h.onInstallButtonClicked),
			),
		)
}
