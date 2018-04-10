package components.movie.title


import react.*
import react.dom.*


external interface Props : RProps {
    var title: String
    var year: String
}

class Title(props: Props) : RComponent<Props, RState>(props) {

    override fun RBuilder.render() {
        div("movie-title") {
            +"${props.title} (${props.year})"
            }
    }
}

fun RBuilder.title(
        title: String,
        year: String
) = child(Title::class) {
    attrs.title=title
    attrs.year=year
}
