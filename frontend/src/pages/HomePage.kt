package pages

import kotlinext.js.jsObject
import react.*
import react.dom.*
import utils.axios
import utils.AxiosConfigSettings
import models.SearchResponse
import models.SearchResponseItem
import components.*

class HomePage : RComponent<RProps, RState>() {

    companion object {
        const val HOST = "http://localhost:9000"

        fun getSelectOptions(results: Array<SearchResponseItem>): List<Option> {
            return results.map { item ->
                val result: Option = jsObject {
                    value = item.imdbID
                    label = item.Title
                }
                result
            }
        }
    }

    private fun loadMovies(searchString: String) {
        val config: AxiosConfigSettings = jsObject {
            url = "$HOST/search?s=$searchString"
            timeout = 3000
        }
        axios<SearchResponse>(config).then { response ->
            val results = response.data.results

            if (results !== null) {
                val options = getSelectOptions(results)
                options.forEach { item -> println(JSON.stringify(item)) }
            }
        }
    }

    private fun onSelectChange(option: Option) {
        println(option)
    }

    private fun onInputValueChange(value: String) {
        loadMovies(value)
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

fun RBuilder.homePage() = child(HomePage::class) {}
