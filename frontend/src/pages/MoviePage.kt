package pages

import kotlinext.js.jsObject
import react.*
import react.dom.*
import utils.axios
import utils.Config
import utils.AxiosConfigSettings
import models.MovieDetails

interface MoviePageProps: RProps {
    var imdbID: String
}

interface MoviePageState : RState {
    var movieDetails: MovieDetails?
}

class MoviePage(props: MoviePageProps) : RComponent<MoviePageProps, MoviePageState>(props) {

    override fun componentDidMount() {
        loadMovies(props.imdbID)

    }
    private fun loadMovies(imdbID: String) {
        val config: AxiosConfigSettings = jsObject {
            url = "${Config.HOST}/movie/$imdbID"
            timeout = 3000
        }
        axios<MovieDetails>(config).then { response ->
            val newMovieDetails = response.data
            setState {
                movieDetails = newMovieDetails
            }
        }
    }

    override fun RBuilder.render() {
        if (state.movieDetails != null) {
            div {
                +"${JSON.stringify(state.movieDetails)}"
            }
        } else {
            div {}
        }
    }
}

fun RBuilder.moviePage(imdbID: String) = child(MoviePage::class) {
    attrs.imdbID = imdbID
}
