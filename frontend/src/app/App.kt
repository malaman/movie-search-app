package app

import react.*
import react.dom.*
import components.*
import pages.HomePage

class RootComponent : RComponent<RProps, RState>() {
    override fun RBuilder.render() {
        div("container"){
            hashRouter {
                switch {
                    route("/",  HomePage::class, exact = true)
                }
            }
        }
    }
}

fun RBuilder.app() = child(RootComponent::class) {}
