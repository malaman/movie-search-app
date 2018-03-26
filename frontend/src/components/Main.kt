package components

import kotlinext.js.jsObject
import react.*
import react.dom.*
import utils.axios
import utils.AxiosConfigSettings
import models.SearchResponse

class Main : RComponent<RProps, RState>() {

    override fun componentDidMount() {
        val config: AxiosConfigSettings = jsObject {
            url = "http://localhost:9000/search?s=lord"
            timeout = 3000
        }
        axios<SearchResponse>(config).then { response ->
            println(response.data.results[0].Title)
        }
    }

    override fun RBuilder.render() {
        div {
            "This is custom component with xhr request"
        }
    }
}

fun RBuilder.main() = child(Main::class) {}
