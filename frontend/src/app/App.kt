package app

import react.*
import react.dom.*
import pages.homePage

class App : RComponent<RProps, RState>() {
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
