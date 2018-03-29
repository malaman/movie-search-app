package app

import kotlinext.js.jsObject
import react.*
import react.dom.*
import pages.homePage
import utils.UniversalRouter
import utils.Route

class App : RComponent<RProps, RState>() {

    override fun componentDidMount() {
        val routes: Array<Route> = arrayOf(
                jsObject {
                    path = "/"
                    action = { println("hi")}
                }
        )

          val router = UniversalRouter(routes, jsObject {})
            println(router.resolve)
    }

    override fun RBuilder.render() {
        div("container") {
            div {
                +"Enter movie title"
            }
            homePage()
        }
    }
}

fun RBuilder.app() = child(App::class) {}
