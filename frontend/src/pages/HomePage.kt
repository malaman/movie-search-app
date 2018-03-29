package pages

import kotlinext.js.jsObject
import react.*
import react.dom.*
import utils.axios
import utils.AxiosConfigSettings
import models.SearchResponse
import models.SearchResponseItem
import components.*

interface HomePageState : RState {
    var options: Array<Option>
    var selectedOption: Option?
}

class HomePage : RComponent<RProps, HomePageState >() {

    companion object {
        const val HOST = "http://localhost:9000"
        /**
         * Prepares array of options for react select
         *
         */
        fun getSelectOptions(results: Array<SearchResponseItem>, searchString: String): Array<Option> {
            val result = results.map { item ->
                val result: Option = jsObject {
                    value = item.imdbID
                    label = item.Title
                }
                result
            }
            if (result.isNotEmpty()) {
                val lastOption: Option = jsObject {
                    value = "show_all"
                    label = "See all results for: $searchString"
                }
                return result.plus(lastOption).toTypedArray()
            }
            return result.toTypedArray()
        }
    }

    override fun HomePageState.init(props: RProps) {
        options = emptyArray()
        selectedOption = null
    }

    private fun loadMovies(searchString: String) {
        val config: AxiosConfigSettings = jsObject {
            url = "$HOST/search?s=$searchString"
            timeout = 3000
        }
        axios<SearchResponse>(config).then { response ->
            val results = response.data.results

            if (results !== null) {
                val nextOptions = getSelectOptions(results, searchString)
                if (nextOptions.isNotEmpty()) {
                }
                setState {
                    options = nextOptions
                }
            }
        }
    }

    private fun onSelectChange(option: Option?) {
        setState {
            selectedOption = option
        }
    }

    private fun onInputValueChange(value: String) {
        loadMovies(value)
    }

    override fun RBuilder.render() {
        div {
            customSelect(
                    name="12313",
                    value=state.selectedOption?.value,
                    options=state.options,
                    onChange={opt: Option? -> onSelectChange(opt)},
                    onInputChange = {value: String -> onInputValueChange(value)}
            )
        }
    }
}

fun RBuilder.homePage() = child(HomePage::class) {}
