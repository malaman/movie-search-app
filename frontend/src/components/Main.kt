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

    fun onSelectChange(option: Option) {
        println(option)
    }

    fun onInputValueChange(value: String) {
        println(value)
    }

    val options : Array<Option> = arrayOf(
        jsObject {
            value = "1"
            label = "1"
        },
        jsObject {
            value = "2"
            label = "2"
        }        
    )  


    override fun RBuilder.render() {
        div {
            "This is custom component with xhr request"
            customSelect(
                name="12313",
                value="1",
                options=options,
                onChange={opt: Option -> onSelectChange(opt)},
                onInputChange = {value: String -> onInputValueChange(value)}                
            )
        }
    }
}

fun RBuilder.main() = child(Main::class) {}
