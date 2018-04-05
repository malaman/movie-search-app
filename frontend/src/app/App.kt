package app

import react.*
import react.dom.*
import react.router.dom.*
import pages.HomePage
import pages.moviePage

interface MovieIDProps : RProps {
    var imdbID: String
}


class RootComponent : RComponent<RProps, RState>() {
    override fun RBuilder.render() {
        div("container"){
            hashRouter {
                switch {
                    route("/",  HomePage::class, exact = true)
                    route<MovieIDProps>("/movie/:imdbID") { props ->
                        moviePage(props.match.params.imdbID)
                    }
                }
            }
        }
    }
}

fun RBuilder.app() = child(RootComponent::class) {}
