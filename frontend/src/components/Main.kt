package components

import kotlinext.js.jsObject
import react.*
import react.dom.*
import utils.axios
import utils.AxiosConfigSettings

class Main : RComponent<RProps, RState>() {

    override fun componentDidMount() {
        val config: AxiosConfigSettings = jsObject {
            url = "http://localhost:9000/search?s=lord"
            timeout = 3000
        }
        axios<Any>(config).then { response ->
            println(response.data)
        }
    }

    override fun RBuilder.render() {
        div {
            "This is custom component with xhr request"
        }
    }
}

fun RBuilder.main() = child(Main::class) {}
