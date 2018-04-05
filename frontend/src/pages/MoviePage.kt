package pages

import react.*
import react.dom.*

interface MoviePageProps: RProps {
    var imdbID: String
}

class MoviePage(props: MoviePageProps) : RComponent<MoviePageProps, RState >(props) {


    override fun RBuilder.render() {
        div {
            +"${props.imdbID}"
        }
    }
}

fun RBuilder.moviePage(imdbID: String) = child(MoviePage::class) {
    attrs.imdbID = imdbID
}
